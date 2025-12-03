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

// ClaudeProvider implements the Provider interface for Anthropic Claude API.
type ClaudeProvider struct {
	endpoint       string
	model          string
	apiKey         string
	minCases       int
	maxCases       int
	maxRetries     int
	requestTimeout time.Duration
	client         *http.Client
}

// NewClaudeProvider creates a new Claude provider with the given config.
// Returns an error if the API key is missing or endpoint URL is invalid.
func NewClaudeProvider(cfg *Config) (*ClaudeProvider, error) {
	if cfg == nil {
		cfg = &Config{
			Provider: ProviderClaude,
			Model:    DefaultClaudeModel,
			Endpoint: DefaultClaudeEndpoint,
		}
	}

	if cfg.APIKey == "" {
		return nil, fmt.Errorf("Claude API key is required")
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

	return &ClaudeProvider{
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
func (c *ClaudeProvider) Name() string {
	return ProviderClaude
}

// IsAvailable checks if the Claude API is accessible with the configured API key.
// Claude doesn't have a simple health check endpoint, so we verify the API key format.
func (c *ClaudeProvider) IsAvailable() bool {
	// Claude API keys start with "sk-ant-"
	// We can't do a lightweight health check without making a real API call
	// So we just validate the API key format
	return len(c.apiKey) > 0 && c.endpoint != ""
}

// GenerateTestCases generates test cases using Claude API.
func (c *ClaudeProvider) GenerateTestCases(ctx context.Context, fn *models.Function) ([]TestCase, error) {
	scaffold := buildTestScaffold(fn)
	prompt := buildGoPrompt(fn, scaffold, c.minCases, c.maxCases, "")

	var lastErr error
	for attempt := 0; attempt < c.maxRetries; attempt++ {
		if err := ctx.Err(); err != nil {
			return nil, fmt.Errorf("context cancelled: %w", err)
		}

		if attempt > 0 && lastErr != nil {
			prompt = buildGoPrompt(fn, scaffold, c.minCases, c.maxCases, lastErr.Error())
		}

		cases, err := c.callAPI(ctx, prompt)
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

	return nil, fmt.Errorf("failed after %d attempts: %w", c.maxRetries, lastErr)
}

// claudeRequest represents the request body for Claude messages API.
type claudeRequest struct {
	Model     string          `json:"model"`
	MaxTokens int             `json:"max_tokens"`
	System    string          `json:"system,omitempty"`
	Messages  []claudeMessage `json:"messages"`
}

// claudeMessage represents a message in the Claude format.
type claudeMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// claudeResponse represents the response from Claude messages API.
type claudeResponse struct {
	Content []struct {
		Type string `json:"type"`
		Text string `json:"text"`
	} `json:"content"`
	Error *struct {
		Type    string `json:"type"`
		Message string `json:"message"`
	} `json:"error,omitempty"`
}

// callAPI makes a request to the Claude messages API.
func (c *ClaudeProvider) callAPI(ctx context.Context, prompt string) ([]TestCase, error) {
	reqBody := claudeRequest{
		Model:     c.model,
		MaxTokens: 4096,
		System:    "You are a Go testing expert. Generate test cases in the exact format requested. Output only the Go code for test cases, no explanations.",
		Messages: []claudeMessage{
			{
				Role:    "user",
				Content: prompt,
			},
		},
	}

	body, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("marshal request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", c.endpoint+"/messages", bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", c.apiKey)
	req.Header.Set("anthropic-version", "2023-06-01")

	resp, err := c.client.Do(req)
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
		return nil, fmt.Errorf("Claude API returned %d: %s", resp.StatusCode, string(respBody))
	}

	var result claudeResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("decode response: %w", err)
	}

	if result.Error != nil {
		return nil, fmt.Errorf("Claude API error: %s", result.Error.Message)
	}

	if len(result.Content) == 0 {
		return nil, fmt.Errorf("Claude returned no content")
	}

	// Find text content
	var content string
	for _, c := range result.Content {
		if c.Type == "text" {
			content = c.Text
			break
		}
	}

	if content == "" {
		return nil, fmt.Errorf("Claude returned no text content")
	}

	cases, err := parseGoTestCases(content, c.maxCases)
	if err != nil {
		return nil, fmt.Errorf("parse test cases: %w\n\nLLM Response:\n%s", err, content)
	}

	return cases, nil
}
