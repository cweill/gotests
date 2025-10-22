// Package ai provides AI-powered test case generation for gotests.
package ai

import (
	"context"

	"github.com/cweill/gotests/internal/models"
)

// Provider is the interface for AI test case generation backends.
type Provider interface {
	// GenerateTestCases generates test cases for the given function.
	GenerateTestCases(ctx context.Context, fn *models.Function) ([]TestCase, error)

	// IsAvailable checks if the provider is available and configured.
	IsAvailable() bool

	// Name returns the provider name for logging/debugging.
	Name() string
}

// TestCase represents a single generated test case.
type TestCase struct {
	Name        string            // Test case name (e.g., "positive_numbers")
	Description string            // Optional description
	Args        map[string]string // Parameter name -> Go code value
	Want        map[string]string // Return value name -> Go code value
	WantErr     bool              // Whether an error is expected
}

// Config holds configuration for AI providers.
type Config struct {
	Provider       string // Provider name: "ollama", "openai", "claude"
	Model          string // Model name (e.g., "qwen2.5-coder:0.5b")
	Endpoint       string // API endpoint URL
	APIKey         string // API key (for cloud providers)
	NumCases       int    // Number of test cases to generate (default: 3)
	MaxRetries     int    // Maximum number of retry attempts (default: 3)
	RequestTimeout int    // HTTP request timeout in seconds (default: 60)
	HealthTimeout  int    // Health check timeout in seconds (default: 2)
}

// DefaultConfig returns the default AI configuration.
func DefaultConfig() *Config {
	return &Config{
		Provider:       "ollama",
		Model:          "qwen2.5-coder:0.5b",
		Endpoint:       "http://localhost:11434",
		NumCases:       3,
		MaxRetries:     3,
		RequestTimeout: 60,
		HealthTimeout:  2,
	}
}
