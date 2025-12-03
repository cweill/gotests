// Package ai provides AI-powered test case generation for gotests.
package ai

import (
	"context"
	"fmt"
	"strings"

	"github.com/cweill/gotests/internal/models"
)

// Provider name constants.
const (
	ProviderOllama     = "ollama"     // Ollama native API
	ProviderOpenAI     = "openai"     // OpenAI-compatible API (works with many providers)
	ProviderClaude     = "claude"     // Anthropic Claude native API
	ProviderGemini     = "gemini"     // Google Gemini native API
	ProviderCompatible = "compatible" // Alias for openai (OpenAI-compatible)
)

// Default endpoints for providers.
const (
	DefaultOllamaEndpoint = "http://localhost:11434"
	DefaultOpenAIEndpoint = "https://api.openai.com/v1"
	DefaultClaudeEndpoint = "https://api.anthropic.com/v1"
	DefaultGeminiEndpoint = "https://generativelanguage.googleapis.com/v1beta"
)

// Default models for each provider.
const (
	DefaultOllamaModel = "qwen2.5-coder:0.5b"
	DefaultOpenAIModel = "gpt-4o-mini"
	DefaultClaudeModel = "claude-3-haiku-20240307"
	DefaultGeminiModel = "gemini-1.5-flash"
)

// Common OpenAI-compatible endpoints for reference.
// Users can use -ai-provider openai with any of these endpoints.
var CommonEndpoints = map[string]string{
	"openai":     "https://api.openai.com/v1",
	"azure":      "https://{resource}.openai.azure.com/openai/deployments/{deployment}",
	"together":   "https://api.together.xyz/v1",
	"groq":       "https://api.groq.com/openai/v1",
	"anyscale":   "https://api.endpoints.anyscale.com/v1",
	"fireworks":  "https://api.fireworks.ai/inference/v1",
	"deepseek":   "https://api.deepseek.com/v1",
	"mistral":    "https://api.mistral.ai/v1",
	"perplexity": "https://api.perplexity.ai",
	"openrouter": "https://openrouter.ai/api/v1",
	"lmstudio":   "http://localhost:1234/v1",
}

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
	Provider       string // Provider name: "ollama", "openai", "claude", "gemini"
	Model          string // Model name (e.g., "qwen2.5-coder:0.5b", "gpt-4o", "claude-3-sonnet")
	Endpoint       string // API endpoint URL
	APIKey         string // API key (for cloud providers)
	MinCases       int    // Minimum number of test cases to generate (default: 3)
	MaxCases       int    // Maximum number of test cases to generate (default: 10)
	MaxRetries     int    // Maximum number of retry attempts (default: 3)
	RequestTimeout int    // HTTP request timeout in seconds (default: 60)
	HealthTimeout  int    // Health check timeout in seconds (default: 2)
}

// DefaultConfig returns the default AI configuration.
func DefaultConfig() *Config {
	return &Config{
		Provider:       ProviderOllama,
		Model:          DefaultOllamaModel,
		Endpoint:       DefaultOllamaEndpoint,
		MinCases:       3,
		MaxCases:       10,
		MaxRetries:     3,
		RequestTimeout: 60,
		HealthTimeout:  2,
	}
}

// NewProvider creates a new AI provider based on the configuration.
// It returns an error if the provider is unknown or configuration is invalid.
//
// Supported providers:
//   - "ollama": Ollama native API (local, free)
//   - "openai" or "compatible": OpenAI-compatible API (works with OpenAI, Azure, Together, Groq, etc.)
//   - "claude": Anthropic Claude native API
//   - "gemini": Google Gemini native API
//
// For OpenAI-compatible providers, you can use any endpoint that follows the OpenAI API format.
func NewProvider(cfg *Config) (Provider, error) {
	if cfg == nil {
		cfg = DefaultConfig()
	}

	// Normalize provider name
	provider := strings.ToLower(strings.TrimSpace(cfg.Provider))
	cfg.Provider = provider

	// Set default endpoint and model based on provider if not specified
	cfg = applyProviderDefaults(cfg)

	switch provider {
	case ProviderOllama:
		return NewOllamaProvider(cfg)
	case ProviderOpenAI, ProviderCompatible, "openai-compatible":
		// "openai", "compatible", and "openai-compatible" all use the OpenAI-compatible provider
		return NewOpenAIProvider(cfg)
	case ProviderClaude, "anthropic":
		// "claude" and "anthropic" both use the Claude provider
		return NewClaudeProvider(cfg)
	case ProviderGemini, "google":
		// "gemini" and "google" both use the Gemini provider
		return NewGeminiProvider(cfg)
	default:
		return nil, fmt.Errorf("unknown AI provider: %q (supported: ollama, openai/compatible, claude/anthropic, gemini/google)", provider)
	}
}

// applyProviderDefaults fills in default endpoint and model based on the provider.
func applyProviderDefaults(cfg *Config) *Config {
	// Create a copy to avoid modifying the original
	result := *cfg

	switch cfg.Provider {
	case ProviderOllama:
		if result.Endpoint == "" {
			result.Endpoint = DefaultOllamaEndpoint
		}
		if result.Model == "" {
			result.Model = DefaultOllamaModel
		}
	case ProviderOpenAI, ProviderCompatible, "openai-compatible":
		if result.Endpoint == "" {
			result.Endpoint = DefaultOpenAIEndpoint
		}
		if result.Model == "" {
			result.Model = DefaultOpenAIModel
		}
	case ProviderClaude, "anthropic":
		if result.Endpoint == "" {
			result.Endpoint = DefaultClaudeEndpoint
		}
		if result.Model == "" {
			result.Model = DefaultClaudeModel
		}
	case ProviderGemini, "google":
		if result.Endpoint == "" {
			result.Endpoint = DefaultGeminiEndpoint
		}
		if result.Model == "" {
			result.Model = DefaultGeminiModel
		}
	}

	return &result
}

// ListSupportedProviders returns a list of supported provider names.
func ListSupportedProviders() []string {
	return []string{
		ProviderOllama,
		ProviderOpenAI,
		ProviderClaude,
		ProviderGemini,
	}
}

// GetCommonEndpoint returns the endpoint URL for a well-known provider name.
// Returns empty string if the provider is not in the common endpoints list.
func GetCommonEndpoint(name string) string {
	return CommonEndpoints[strings.ToLower(name)]
}
