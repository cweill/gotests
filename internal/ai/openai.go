package ai

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/cweill/gotests/internal/models"
)

// OpenAIProvider implements the Provider interface for OpenAI-compatible APIs.
// This provider works with any service that implements the OpenAI Chat Completions API format:
//   - OpenAI (api.openai.com)
//   - Azure OpenAI
//   - Together AI (api.together.xyz)
//   - Groq (api.groq.com)
//   - Anyscale (api.endpoints.anyscale.com)
//   - Fireworks AI (api.fireworks.ai)
//   - DeepSeek (api.deepseek.com)
//   - Mistral AI (api.mistral.ai)
//   - Perplexity (api.perplexity.ai)
//   - OpenRouter (openrouter.ai)
//   - LM Studio (localhost:1234)
//   - And many more...
type OpenAIProvider struct {
	endpoint       string
	model          string
	apiKey         string
	minCases       int
	maxCases       int
	maxRetries     int
	requestTimeout time.Duration
	client         *http.Client
}

// NewOpenAIProvider creates a new OpenAI-compatible provider with the given config.
// This works with any service implementing the OpenAI Chat Completions API.
// Returns an error if the API key is missing or endpoint URL is invalid.
func NewOpenAIProvider(cfg *Config) (*OpenAIProvider, error) {
	if cfg == nil {
		cfg = &Config{
			Provider: ProviderOpenAI,
			Model:    DefaultOpenAIModel,
			Endpoint: DefaultOpenAIEndpoint,
		}
	}

	// API key is required for cloud providers
	// For local providers (localhost), it might be optional
	if cfg.APIKey == "" && !isLocalEndpoint(cfg.Endpoint) {
		return nil, fmt.Errorf("API key is required for cloud providers (set -ai-api-key or GOTESTS_AI_API_KEY)")
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

	return &OpenAIProvider{
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
func (o *OpenAIProvider) Name() string {
	// Return a more descriptive name based on the endpoint
	if strings.Contains(o.endpoint, "openai.com") {
		return "openai"
	}
	if strings.Contains(o.endpoint, "together.xyz") {
		return "together"
	}
	if strings.Contains(o.endpoint, "groq.com") {
		return "groq"
	}
	if strings.Contains(o.endpoint, "anthropic.com") {
		return "anthropic"
	}
	if strings.Contains(o.endpoint, "mistral.ai") {
		return "mistral"
	}
	if strings.Contains(o.endpoint, "deepseek.com") {
		return "deepseek"
	}
	if strings.Contains(o.endpoint, "fireworks.ai") {
		return "fireworks"
	}
	if strings.Contains(o.endpoint, "openrouter.ai") {
		return "openrouter"
	}
	if isLocalEndpoint(o.endpoint) {
		return "openai-compatible (local)"
	}
	return "openai-compatible"
}

// isLocalEndpoint checks if the endpoint is a local address.
func isLocalEndpoint(endpoint string) bool {
	return strings.Contains(endpoint, "localhost") ||
		strings.Contains(endpoint, "127.0.0.1") ||
		strings.Contains(endpoint, "0.0.0.0") ||
		strings.HasPrefix(endpoint, "http://192.168.") ||
		strings.HasPrefix(endpoint, "http://10.") ||
		strings.HasPrefix(endpoint, "http://172.")
}

// IsAvailable checks if the API is accessible.
// For cloud providers, it verifies the API key by making a lightweight request.
// For local providers, it just checks connectivity.
func (o *OpenAIProvider) IsAvailable() bool {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Check by listing models - a lightweight API call
	req, err := http.NewRequestWithContext(ctx, "GET", o.endpoint+"/models", nil)
	if err != nil {
		return false
	}
	if o.apiKey != "" {
		req.Header.Set("Authorization", "Bearer "+o.apiKey)
	}

	resp, err := o.client.Do(req)
	if err != nil {
		// For local endpoints, try a simple health check
		if isLocalEndpoint(o.endpoint) {
			return o.checkLocalHealth(ctx)
		}
		return false
	}
	defer resp.Body.Close()

	// Accept 200 (OK) or 401 (unauthorized but endpoint exists)
	// 401 means the endpoint is there but API key might be wrong (still "available")
	return resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusUnauthorized
}

// checkLocalHealth performs a simple connectivity check for local endpoints.
func (o *OpenAIProvider) checkLocalHealth(ctx context.Context) bool {
	req, err := http.NewRequestWithContext(ctx, "GET", o.endpoint, nil)
	if err != nil {
		return false
	}
	resp, err := o.client.Do(req)
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	return true
}

// GenerateTestCases generates test cases using OpenAI API.
func (o *OpenAIProvider) GenerateTestCases(ctx context.Context, fn *models.Function) ([]TestCase, error) {
	scaffold := buildTestScaffold(fn)
	prompt := buildGoPrompt(fn, scaffold, o.minCases, o.maxCases, "")

	var lastErr error
	for attempt := 0; attempt < o.maxRetries; attempt++ {
		if err := ctx.Err(); err != nil {
			return nil, fmt.Errorf("context cancelled: %w", err)
		}

		if attempt > 0 && lastErr != nil {
			prompt = buildGoPrompt(fn, scaffold, o.minCases, o.maxCases, lastErr.Error())
		}

		cases, err := o.callAPI(ctx, prompt)
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

	return nil, fmt.Errorf("failed after %d attempts: %w", o.maxRetries, lastErr)
}

// openAIRequest represents the request body for OpenAI chat completions API.
type openAIRequest struct {
	Model       string          `json:"model"`
	Messages    []openAIMessage `json:"messages"`
	Temperature float64         `json:"temperature"`
	MaxTokens   int             `json:"max_tokens,omitempty"`
}

// openAIMessage represents a message in the OpenAI chat format.
type openAIMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// openAIResponse represents the response from OpenAI chat completions API.
type openAIResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
	Error *struct {
		Message string `json:"message"`
		Type    string `json:"type"`
	} `json:"error,omitempty"`
}

// callAPI makes a request to the OpenAI chat completions API.
func (o *OpenAIProvider) callAPI(ctx context.Context, prompt string) ([]TestCase, error) {
	reqBody := openAIRequest{
		Model: o.model,
		Messages: []openAIMessage{
			{
				Role:    "system",
				Content: "You are a Go testing expert. Generate test cases in the exact format requested. Output only the Go code for test cases, no explanations.",
			},
			{
				Role:    "user",
				Content: prompt,
			},
		},
		Temperature: 0.0, // Deterministic output
		MaxTokens:   4096,
	}

	body, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("marshal request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", o.endpoint+"/chat/completions", bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	if o.apiKey != "" {
		req.Header.Set("Authorization", "Bearer "+o.apiKey)
	}

	resp, err := o.client.Do(req)
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
		return nil, fmt.Errorf("OpenAI API returned %d: %s", resp.StatusCode, string(respBody))
	}

	var result openAIResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("decode response: %w", err)
	}

	if result.Error != nil {
		return nil, fmt.Errorf("OpenAI API error: %s", result.Error.Message)
	}

	if len(result.Choices) == 0 {
		return nil, fmt.Errorf("OpenAI returned no choices")
	}

	content := result.Choices[0].Message.Content
	cases, err := parseGoTestCases(content, o.maxCases)
	if err != nil {
		return nil, fmt.Errorf("parse test cases: %w\n\nLLM Response:\n%s", err, content)
	}

	return cases, nil
}

