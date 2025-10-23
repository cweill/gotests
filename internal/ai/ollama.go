package ai

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/cweill/gotests/internal/models"
)

const (
	// MaxResponseSize is the maximum size of HTTP responses (1MB)
	MaxResponseSize = 1 * 1024 * 1024
	// MaxFunctionBodySize is the maximum size of function body in prompts (100KB)
	MaxFunctionBodySize = 100 * 1024
)

// OllamaProvider implements the Provider interface for Ollama.
type OllamaProvider struct {
	endpoint       string
	model          string
	numCases       int
	maxRetries     int
	requestTimeout time.Duration
	healthTimeout  time.Duration
	client         *http.Client
}

// NewOllamaProvider creates a new Ollama provider with the given config.
// Returns an error if the endpoint URL is invalid or unsafe.
func NewOllamaProvider(cfg *Config) (*OllamaProvider, error) {
	if cfg == nil {
		cfg = DefaultConfig()
	}

	// Validate endpoint URL
	if err := validateEndpointURL(cfg.Endpoint); err != nil {
		return nil, fmt.Errorf("invalid endpoint: %w", err)
	}

	return &OllamaProvider{
		endpoint:       cfg.Endpoint,
		model:          cfg.Model,
		numCases:       cfg.NumCases,
		maxRetries:     cfg.MaxRetries,
		requestTimeout: time.Duration(cfg.RequestTimeout) * time.Second,
		healthTimeout:  time.Duration(cfg.HealthTimeout) * time.Second,
		client: &http.Client{
			Timeout: time.Duration(cfg.RequestTimeout) * time.Second,
		},
	}, nil
}

// validateEndpointURL checks if the endpoint URL is safe to use.
// Prevents SSRF attacks by validating the URL scheme and format.
func validateEndpointURL(endpoint string) error {
	if endpoint == "" {
		return fmt.Errorf("endpoint cannot be empty")
	}

	parsedURL, err := url.Parse(endpoint)
	if err != nil {
		return fmt.Errorf("invalid URL format: %w", err)
	}

	// Only allow http and https schemes
	if parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
		return fmt.Errorf("invalid URL scheme %q: only http and https are allowed", parsedURL.Scheme)
	}

	// Ensure host is present
	if parsedURL.Host == "" {
		return fmt.Errorf("URL must include a host")
	}

	return nil
}

// Name returns the provider name.
func (o *OllamaProvider) Name() string {
	return "ollama"
}

// IsAvailable checks if Ollama is running and accessible.
func (o *OllamaProvider) IsAvailable() bool {
	ctx, cancel := context.WithTimeout(context.Background(), o.healthTimeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", o.endpoint+"/api/tags", nil)
	if err != nil {
		return false
	}

	resp, err := o.client.Do(req)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	return resp.StatusCode == http.StatusOK
}

// GenerateTestCases generates test cases using Ollama.
// Now uses Go code generation instead of JSON for better small-model compatibility.
func (o *OllamaProvider) GenerateTestCases(ctx context.Context, fn *models.Function) ([]TestCase, error) {
	// Generate a minimal scaffold to show the LLM the struct format
	scaffold := buildTestScaffold(fn)

	prompt := buildGoPrompt(fn, scaffold, o.numCases, "")

	// Try up to maxRetries times with validation feedback
	var lastErr error
	for attempt := 0; attempt < o.maxRetries; attempt++ {
		// Check if context has been cancelled
		if err := ctx.Err(); err != nil {
			return nil, fmt.Errorf("context cancelled: %w", err)
		}

		if attempt > 0 && lastErr != nil {
			// Retry with error feedback
			prompt = buildGoPrompt(fn, scaffold, o.numCases, lastErr.Error())
		}

		cases, err := o.generateGo(ctx, prompt)
		if err != nil {
			lastErr = err
			continue
		}

		// Basic validation of test cases
		if err := validateTestCases(cases, fn); err != nil {
			lastErr = err
			continue
		}

		return cases, nil
	}

	return nil, fmt.Errorf("failed after %d attempts: %w", o.maxRetries, lastErr)
}

// generateGo calls Ollama and parses Go code response instead of JSON.
func (o *OllamaProvider) generateGo(ctx context.Context, prompt string) ([]TestCase, error) {
	reqBody := map[string]interface{}{
		"model":  o.model,
		"prompt": prompt,
		"stream": false,
		"options": map[string]interface{}{
			"temperature": 0.0, // Deterministic generation for test cases
			"seed":        42,  // Fixed seed for deterministic output
		},
	}

	body, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("marshal request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", o.endpoint+"/api/generate", bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := o.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(io.LimitReader(resp.Body, MaxResponseSize))
		return nil, fmt.Errorf("ollama returned %d: %s", resp.StatusCode, string(body))
	}

	// Limit response size to prevent memory exhaustion
	limitedReader := io.LimitReader(resp.Body, MaxResponseSize)

	var result struct {
		Response string `json:"response"`
	}
	if err := json.NewDecoder(limitedReader).Decode(&result); err != nil {
		return nil, fmt.Errorf("decode response: %w", err)
	}

	// Parse the Go code response
	cases, err := parseGoTestCases(result.Response, o.numCases)
	if err != nil {
		return nil, fmt.Errorf("parse Go test cases: %w\n\nLLM Response:\n%s", err, result.Response)
	}

	return cases, nil
}

// validateTestCases performs basic validation on generated test cases.
func validateTestCases(cases []TestCase, fn *models.Function) error {
	if len(cases) == 0 {
		return fmt.Errorf("no test cases generated")
	}

	// Check for duplicate test cases (same args/want values)
	if hasDuplicates(cases) {
		return fmt.Errorf("generated test cases have duplicate values - need more diversity")
	}

	// Check each test case has required fields
	for i, tc := range cases {
		if tc.Name == "" {
			return fmt.Errorf("test case %d missing name", i)
		}

		// Check all function parameters are provided
		for _, param := range fn.TestParameters() {
			if _, exists := tc.Args[param.Name]; !exists {
				return fmt.Errorf("test case %q missing argument: %s", tc.Name, param.Name)
			}
		}

		// Check return values match
		// Note: fn.TestResults() already excludes the error (error is indicated by fn.ReturnsError)
		// so we just need to check that Want has the same number of entries as TestResults
		expectedReturns := len(fn.TestResults())
		if len(tc.Want) != expectedReturns {
			return fmt.Errorf("test case %q has %d return values, expected %d", tc.Name, len(tc.Want), expectedReturns)
		}
	}

	return nil
}

// hasDuplicates checks if test cases have duplicate args+want combinations.
func hasDuplicates(cases []TestCase) bool {
	if len(cases) <= 1 {
		return false
	}

	// Check if all test cases have identical args+want
	// (simplified check - if all args match for first case, likely all duplicates)
	firstCase := cases[0]
	allIdentical := true

	for i := 1; i < len(cases); i++ {
		tc := cases[i]

		// Check if args match
		if len(tc.Args) != len(firstCase.Args) {
			allIdentical = false
			break
		}
		for key, val := range tc.Args {
			if firstVal, exists := firstCase.Args[key]; !exists || firstVal != val {
				allIdentical = false
				break
			}
		}

		// Check if want matches
		if len(tc.Want) != len(firstCase.Want) {
			allIdentical = false
			break
		}
		for key, val := range tc.Want {
			if firstVal, exists := firstCase.Want[key]; !exists || firstVal != val {
				allIdentical = false
				break
			}
		}

		if !allIdentical {
			break
		}
	}

	return allIdentical
}
