//go:build e2e

package ai

import (
	"context"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/cweill/gotests/internal/goparser"
	"github.com/cweill/gotests/internal/models"
)

// e2eTestCase represents a test case for E2E validation against golden files.
type e2eTestCase struct {
	name       string // Test name
	sourceFile string // Path to source file
	funcName   string // Function name to generate tests for
	goldenFile string // Path to expected golden output
}

// E2E test cases that validate AI generation against golden files.
// These use real Ollama + qwen2.5-coder:0.5b to ensure deterministic generation.
var e2eTestCases = []e2eTestCase{
	{
		name:       "business_logic_calculate_discount",
		sourceFile: "../../testdata/business_logic.go",
		funcName:   "CalculateDiscount",
		goldenFile: "../../testdata/goldens/business_logic_calculate_discount_ai.go",
	},
	{
		name:       "math_ops_clamp",
		sourceFile: "../../testdata/math_ops.go",
		funcName:   "Clamp",
		goldenFile: "../../testdata/goldens/math_ops_clamp_ai.go",
	},
	{
		name:       "data_processing_filter_positive",
		sourceFile: "../../testdata/data_processing.go",
		funcName:   "FilterPositive",
		goldenFile: "../../testdata/goldens/data_processing_filter_positive_ai.go",
	},
	{
		name:       "user_service_hash_password",
		sourceFile: "../../testdata/user_service.go",
		funcName:   "HashPassword",
		goldenFile: "../../testdata/goldens/user_service_hash_password_ai.go",
	},
	{
		name:       "function_with_pointer_parameter",
		sourceFile: "../../testdata/test008.go",
		funcName:   "Foo8",
		goldenFile: "../../testdata/goldens/function_with_pointer_parameter_ai.go",
	},
}

// TestE2E_OllamaGeneration_ValidatesStructure validates that real Ollama+qwen generation
// produces valid test cases with correct structure. This ensures:
// 1. AI generation works end-to-end with real Ollama
// 2. Generated test cases have all required fields
// 3. Test cases match function signature (correct args and return values)
// 4. AI produces reasonable test case names and values
//
// This test REQUIRES Ollama to be running with qwen2.5-coder:0.5b model.
// It will FAIL (not skip) if Ollama is not available.
func TestE2E_OllamaGeneration_ValidatesStructure(t *testing.T) {
	// Ensure Ollama is running with qwen model (fails if not)
	provider := requireOllama(t)

	for _, tc := range e2eTestCases {
		tc := tc // capture range variable
		t.Run(tc.name, func(t *testing.T) {
			// Read source file to get full function body
			sourceCode, err := ioutil.ReadFile(tc.sourceFile)
			if err != nil {
				t.Fatalf("Failed to read source file %s: %v", tc.sourceFile, err)
			}

			// Parse source file to get function metadata
			funcs, err := goparser.Parse(tc.sourceFile, nil)
			if err != nil {
				t.Fatalf("Failed to parse %s: %v", tc.sourceFile, err)
			}

			// Find the target function
			var targetFunc *models.Function
			for _, fn := range funcs {
				if fn.Name == tc.funcName {
					targetFunc = fn
					// Set full body for AI context
					targetFunc.FullBody = string(sourceCode)
					break
				}
			}
			if targetFunc == nil {
				t.Fatalf("Function %s not found in %s", tc.funcName, tc.sourceFile)
			}

			// Generate tests with real AI
			ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
			defer cancel()

			cases, err := provider.GenerateTestCases(ctx, targetFunc)
			if err != nil {
				t.Fatalf("GenerateTestCases() failed for %s: %v", tc.funcName, err)
			}

			if len(cases) == 0 {
				t.Fatalf("GenerateTestCases() returned no test cases for %s", tc.funcName)
			}

			t.Logf("✓ Generated %d test cases for %s", len(cases), tc.funcName)

			// Validate each test case has correct structure
			for i, testCase := range cases {
				// Check test case has a name
				if testCase.Name == "" {
					t.Errorf("Test case %d missing name", i)
				}

				// Check all function parameters are present in Args
				for _, param := range targetFunc.TestParameters() {
					if _, exists := testCase.Args[param.Name]; !exists {
						t.Errorf("Test case %q missing argument %q", testCase.Name, param.Name)
					}
				}

				// Check return values are present in Want
				expectedReturns := len(targetFunc.TestResults())
				if len(testCase.Want) != expectedReturns {
					t.Errorf("Test case %q has %d return values, expected %d",
						testCase.Name, len(testCase.Want), expectedReturns)
				}

				// Log test case for debugging
				t.Logf("  Test case %d: %s", i+1, testCase.Name)
				t.Logf("    Args: %v", testCase.Args)
				t.Logf("    Want: %v", testCase.Want)
				t.Logf("    WantErr: %v", testCase.WantErr)
			}

			// Optional: Validate against golden file expectations
			// Read golden file and check if it contains similar test case patterns
			if goldenContent, err := ioutil.ReadFile(tc.goldenFile); err == nil {
				goldenStr := string(goldenContent)

				// Check that at least one generated test case name appears in golden
				foundMatch := false
				for _, testCase := range cases {
					// Normalize test case name to match golden format (snake_case or similar)
					if strings.Contains(goldenStr, testCase.Name) ||
					   strings.Contains(goldenStr, strings.ReplaceAll(testCase.Name, " ", "_")) {
						foundMatch = true
						break
					}
				}

				if !foundMatch {
					t.Logf("Warning: None of the generated test case names found in golden file %s", tc.goldenFile)
					t.Logf("This might indicate model output has changed - review manually")
				}
			}
		})
	}
}

// requireOllama ensures Ollama is running with qwen2.5-coder:0.5b model.
// This function FAILS the test (not skips) if Ollama is not available.
func requireOllama(t *testing.T) Provider {
	t.Helper()

	// Check if we're in CI or running E2E tests explicitly
	if os.Getenv("CI") != "true" && os.Getenv("GOTESTS_E2E") != "true" {
		t.Log("Hint: Set GOTESTS_E2E=true to run E2E tests locally")
	}

	cfg := &Config{
		Provider:       "ollama",
		Model:          "qwen2.5-coder:0.5b",
		Endpoint:       "http://localhost:11434",
		NumCases:       3,
		MaxRetries:     3,
		RequestTimeout: 60,
		HealthTimeout:  2,
	}

	provider, err := NewOllamaProvider(cfg)
	if err != nil {
		t.Fatalf("REQUIRED: Ollama provider creation failed: %v\n\nTo install Ollama:\n  curl -fsSL https://ollama.com/install.sh | sh\n  ollama serve &\n  ollama pull qwen2.5-coder:0.5b", err)
	}

	if !provider.IsAvailable() {
		t.Fatalf("REQUIRED: Ollama is not running or qwen2.5-coder:0.5b model not available\n\nEnsure Ollama is running:\n  ollama serve &\n  ollama pull qwen2.5-coder:0.5b\n\nThen check:\n  curl http://localhost:11434/api/tags")
	}

	t.Log("✓ Ollama is running with qwen2.5-coder:0.5b model")
	return provider
}

// TestE2E_OllamaHealthCheck validates that Ollama is properly set up.
// This is a simple smoke test to verify the E2E environment is ready.
func TestE2E_OllamaHealthCheck(t *testing.T) {
	provider := requireOllama(t)

	if provider.Name() != "ollama" {
		t.Errorf("Expected provider name 'ollama', got %q", provider.Name())
	}

	t.Log("✓ Ollama health check passed")
}
