package ai

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/cweill/gotests/internal/models"
)

// GeminiProvider implements the Provider interface for Google Gemini API.
type GeminiProvider struct {
	endpoint       string
	model          string
	apiKey         string
	minCases       int
	maxCases       int
	maxRetries     int
	requestTimeout time.Duration
	client         *http.Client
}

// NewGeminiProvider creates a new Gemini provider with the given config.
// Returns an error if the API key is missing or endpoint URL is invalid.
func NewGeminiProvider(cfg *Config) (*GeminiProvider, error) {
	if cfg == nil {
		cfg = &Config{
			Provider: ProviderGemini,
			Model:    DefaultGeminiModel,
			Endpoint: DefaultGeminiEndpoint,
		}
	}

	if cfg.APIKey == "" {
		return nil, fmt.Errorf("Gemini API key is required")
	}

	// Validate endpoint URL
	if err := validateEndpointURL(cfg.Endpoint); err != nil {
		return nil, fmt.Errorf("invalid endpoint: %w", err)
	}

	timeout := time.Duration(cfg.RequestTimeout) * time.Second
	if timeout == 0 {
		timeout = 60 * time.Second
	}

	minCases := cfg.MinCases
	if minCases == 0 {
		minCases = 3
	}

	maxCases := cfg.MaxCases
	if maxCases == 0 {
		maxCases = 10
	}

	maxRetries := cfg.MaxRetries
	if maxRetries == 0 {
		maxRetries = 3
	}

	return &GeminiProvider{
		endpoint:       cfg.Endpoint,
		model:          cfg.Model,
		apiKey:         cfg.APIKey,
		minCases:       minCases,
		maxCases:       maxCases,
		maxRetries:     maxRetries,
		requestTimeout: timeout,
		client: &http.Client{
			Timeout: timeout,
		},
	}, nil
}

// Name returns the provider name.
func (g *GeminiProvider) Name() string {
	return ProviderGemini
}

// IsAvailable checks if the Gemini API is accessible with the configured API key.
func (g *GeminiProvider) IsAvailable() bool {
	// Gemini doesn't have a simple health check endpoint
	// We validate the API key format and endpoint
	return len(g.apiKey) > 0 && g.endpoint != ""
}

// GenerateTestCases generates test cases using Gemini API.
func (g *GeminiProvider) GenerateTestCases(ctx context.Context, fn *models.Function) ([]TestCase, error) {
	scaffold := buildTestScaffold(fn)
	prompt := buildGoPrompt(fn, scaffold, g.minCases, g.maxCases, "")

	var lastErr error
	for attempt := 0; attempt < g.maxRetries; attempt++ {
		if err := ctx.Err(); err != nil {
			return nil, fmt.Errorf("context cancelled: %w", err)
		}

		if attempt > 0 && lastErr != nil {
			prompt = buildGoPrompt(fn, scaffold, g.minCases, g.maxCases, lastErr.Error())
		}

		cases, err := g.callAPI(ctx, prompt)
		if err != nil {
			lastErr = err
			continue
		}

		if err := validateTestCases(cases, fn); err != nil {
			lastErr = err
			continue
		}

		return cases, nil
	}

	return nil, fmt.Errorf("failed after %d attempts: %w", g.maxRetries, lastErr)
}

// geminiRequest represents the request body for Gemini generateContent API.
type geminiRequest struct {
	Contents          []geminiContent        `json:"contents"`
	GenerationConfig  geminiGenerationConfig `json:"generationConfig,omitempty"`
	SystemInstruction *geminiContent         `json:"systemInstruction,omitempty"`
}

// geminiContent represents content in the Gemini format.
type geminiContent struct {
	Parts []geminiPart `json:"parts"`
	Role  string       `json:"role,omitempty"`
}

// geminiPart represents a part of content.
type geminiPart struct {
	Text string `json:"text"`
}

// geminiGenerationConfig represents generation configuration.
type geminiGenerationConfig struct {
	Temperature     float64 `json:"temperature"`
	MaxOutputTokens int     `json:"maxOutputTokens,omitempty"`
}

// geminiResponse represents the response from Gemini generateContent API.
type geminiResponse struct {
	Candidates []struct {
		Content struct {
			Parts []struct {
				Text string `json:"text"`
			} `json:"parts"`
		} `json:"content"`
	} `json:"candidates"`
	Error *struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Status  string `json:"status"`
	} `json:"error,omitempty"`
}

// callAPI makes a request to the Gemini generateContent API.
func (g *GeminiProvider) callAPI(ctx context.Context, prompt string) ([]TestCase, error) {
	reqBody := geminiRequest{
		Contents: []geminiContent{
			{
				Parts: []geminiPart{
					{Text: prompt},
				},
				Role: "user",
			},
		},
		GenerationConfig: geminiGenerationConfig{
			Temperature:     0.0, // Deterministic output
			MaxOutputTokens: 4096,
		},
		SystemInstruction: &geminiContent{
			Parts: []geminiPart{
				{Text: "You are a Go testing expert. Generate test cases in the exact format requested. Output only the Go code for test cases, no explanations."},
			},
		},
	}

	body, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("marshal request: %w", err)
	}

	// Gemini uses query parameter for API key
	url := fmt.Sprintf("%s/models/%s:generateContent?key=%s", g.endpoint, g.model, g.apiKey)

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := g.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	limitedReader := io.LimitReader(resp.Body, MaxResponseSize)
	respBody, err := io.ReadAll(limitedReader)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Gemini API returned %d: %s", resp.StatusCode, string(respBody))
	}

	var result geminiResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("decode response: %w", err)
	}

	if result.Error != nil {
		return nil, fmt.Errorf("Gemini API error: %s", result.Error.Message)
	}

	if len(result.Candidates) == 0 {
		return nil, fmt.Errorf("Gemini returned no candidates")
	}

	if len(result.Candidates[0].Content.Parts) == 0 {
		return nil, fmt.Errorf("Gemini returned no content parts")
	}

	content := result.Candidates[0].Content.Parts[0].Text
	cases, err := parseGoTestCases(content, g.maxCases)
	if err != nil {
		return nil, fmt.Errorf("parse test cases: %w\n\nLLM Response:\n%s", err, content)
	}

	return cases, nil
}
