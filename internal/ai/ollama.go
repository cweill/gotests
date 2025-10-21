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

// OllamaProvider implements the Provider interface for Ollama.
type OllamaProvider struct {
	endpoint string
	model    string
	numCases int
	client   *http.Client
}

// NewOllamaProvider creates a new Ollama provider with the given config.
func NewOllamaProvider(cfg *Config) *OllamaProvider {
	if cfg == nil {
		cfg = DefaultConfig()
	}
	return &OllamaProvider{
		endpoint: cfg.Endpoint,
		model:    cfg.Model,
		numCases: cfg.NumCases,
		client: &http.Client{
			Timeout: 60 * time.Second,
		},
	}
}

// Name returns the provider name.
func (o *OllamaProvider) Name() string {
	return "ollama"
}

// IsAvailable checks if Ollama is running and accessible.
func (o *OllamaProvider) IsAvailable() bool {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
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
func (o *OllamaProvider) GenerateTestCases(ctx context.Context, fn *models.Function) ([]TestCase, error) {
	prompt := buildPrompt(fn, o.numCases, "")

	// Try up to 3 times with validation feedback
	var lastErr error
	for attempt := 0; attempt < 3; attempt++ {
		if attempt > 0 && lastErr != nil {
			// Retry with error feedback
			prompt = buildPrompt(fn, o.numCases, lastErr.Error())
		}

		cases, err := o.generate(ctx, prompt)
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

	return nil, fmt.Errorf("failed after 3 attempts: %w", lastErr)
}

// generate makes the actual API call to Ollama.
func (o *OllamaProvider) generate(ctx context.Context, prompt string) ([]TestCase, error) {
	reqBody := map[string]interface{}{
		"model":  o.model,
		"prompt": prompt,
		"stream": false,
		"options": map[string]interface{}{
			"temperature": 0.0, // Deterministic generation for test cases
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
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("ollama returned %d: %s", resp.StatusCode, string(body))
	}

	var result struct {
		Response string `json:"response"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("decode response: %w", err)
	}

	// Parse the LLM response
	cases, err := parseTestCases(result.Response, o.numCases)
	if err != nil {
		return nil, fmt.Errorf("parse test cases: %w", err)
	}

	return cases, nil
}

// validateTestCases performs basic validation on generated test cases.
func validateTestCases(cases []TestCase, fn *models.Function) error {
	if len(cases) == 0 {
		return fmt.Errorf("no test cases generated")
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

// parseTestCases extracts test cases from the LLM response.
// Expected format: JSON array or multiple JSON objects
func parseTestCases(response string, maxCases int) ([]TestCase, error) {
	// Try to find JSON in the response
	start := strings.Index(response, "[")
	end := strings.LastIndex(response, "]")

	if start == -1 || end == -1 || start > end {
		return nil, fmt.Errorf("no JSON array found in response")
	}

	jsonStr := response[start : end+1]

	var rawCases []map[string]interface{}
	if err := json.Unmarshal([]byte(jsonStr), &rawCases); err != nil {
		return nil, fmt.Errorf("unmarshal JSON: %w", err)
	}

	var cases []TestCase
	for i, raw := range rawCases {
		if i >= maxCases {
			break
		}

		tc := TestCase{
			Args: make(map[string]string),
			Want: make(map[string]string),
		}

		// Extract name
		if name, ok := raw["name"].(string); ok {
			tc.Name = name
		} else {
			tc.Name = fmt.Sprintf("test_case_%d", i+1)
		}

		// Extract description
		if desc, ok := raw["description"].(string); ok {
			tc.Description = desc
		}

		// Extract args
		if args, ok := raw["args"].(map[string]interface{}); ok {
			for k, v := range args {
				tc.Args[k] = fmt.Sprintf("%v", v)
			}
		}

		// Extract want/expected
		if want, ok := raw["want"].(map[string]interface{}); ok {
			for k, v := range want {
				tc.Want[k] = fmt.Sprintf("%v", v)
			}
		} else if expected, ok := raw["expected"].(map[string]interface{}); ok {
			for k, v := range expected {
				tc.Want[k] = fmt.Sprintf("%v", v)
			}
		}

		// Extract wantErr
		if wantErr, ok := raw["wantErr"].(bool); ok {
			tc.WantErr = wantErr
		} else if wantErr, ok := raw["want_error"].(bool); ok {
			tc.WantErr = wantErr
		}

		cases = append(cases, tc)
	}

	if len(cases) == 0 {
		return nil, fmt.Errorf("no valid test cases parsed")
	}

	return cases, nil
}
