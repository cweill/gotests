package ai

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/cweill/gotests/internal/models"
)

func TestNewOllamaProvider(t *testing.T) {
	tests := []struct {
		name         string
		cfg          *Config
		wantEndpoint string
		wantModel    string
		wantNumCases int
	}{
		{
			name:         "with_nil_config_uses_defaults",
			cfg:          nil,
			wantEndpoint: "http://localhost:11434",
			wantModel:    "qwen2.5-coder:0.5b",
			wantNumCases: 3,
		},
		{
			name: "with_custom_config",
			cfg: &Config{
				Endpoint:       "http://custom:8080",
				Model:          "llama3.2:latest",
				NumCases:       5,
				MaxRetries:     3,
				RequestTimeout: 60,
				HealthTimeout:  2,
			},
			wantEndpoint: "http://custom:8080",
			wantModel:    "llama3.2:latest",
			wantNumCases: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			provider, err := NewOllamaProvider(tt.cfg)
			if err != nil {
				t.Fatalf("NewOllamaProvider() error = %v", err)
			}

			if provider.endpoint != tt.wantEndpoint {
				t.Errorf("endpoint = %q, want %q", provider.endpoint, tt.wantEndpoint)
			}
			if provider.model != tt.wantModel {
				t.Errorf("model = %q, want %q", provider.model, tt.wantModel)
			}
			if provider.numCases != tt.wantNumCases {
				t.Errorf("numCases = %d, want %d", provider.numCases, tt.wantNumCases)
			}
			if provider.client == nil {
				t.Error("client is nil")
			}
			if provider.client.Timeout != 60*time.Second {
				t.Errorf("client timeout = %v, want 60s", provider.client.Timeout)
			}
		})
	}
}

func TestNewOllamaProvider_InvalidEndpoints(t *testing.T) {
	tests := []struct {
		name     string
		endpoint string
		wantErr  bool
		errMsg   string
	}{
		{
			name:     "empty_endpoint",
			endpoint: "",
			wantErr:  true,
			errMsg:   "endpoint cannot be empty",
		},
		{
			name:     "invalid_scheme",
			endpoint: "ftp://localhost:8080",
			wantErr:  true,
			errMsg:   "only http and https are allowed",
		},
		{
			name:     "file_scheme",
			endpoint: "file:///etc/passwd",
			wantErr:  true,
			errMsg:   "only http and https are allowed",
		},
		{
			name:     "no_host",
			endpoint: "http://",
			wantErr:  true,
			errMsg:   "must include a host",
		},
		{
			name:     "invalid_url_format",
			endpoint: "not a url at all",
			wantErr:  true,
			errMsg:   "only http and https are allowed",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := &Config{
				Endpoint:       tt.endpoint,
				Model:          "test-model",
				NumCases:       3,
				MaxRetries:     3,
				RequestTimeout: 60,
				HealthTimeout:  2,
			}

			_, err := NewOllamaProvider(cfg)

			if (err != nil) != tt.wantErr {
				t.Errorf("NewOllamaProvider() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr && tt.errMsg != "" {
				if err == nil || !contains(err.Error(), tt.errMsg) {
					t.Errorf("NewOllamaProvider() error = %v, want error containing %q", err, tt.errMsg)
				}
			}
		})
	}
}

func TestOllamaProvider_Name(t *testing.T) {
	provider, err := NewOllamaProvider(nil)
	if err != nil {
		t.Fatalf("NewOllamaProvider() error = %v", err)
	}
	if got := provider.Name(); got != "ollama" {
		t.Errorf("Name() = %q, want %q", got, "ollama")
	}
}

func TestOllamaProvider_IsAvailable(t *testing.T) {
	tests := []struct {
		name       string
		statusCode int
		delay      time.Duration
		want       bool
	}{
		{
			name:       "service_available",
			statusCode: http.StatusOK,
			delay:      0,
			want:       true,
		},
		{
			name:       "service_unavailable",
			statusCode: http.StatusServiceUnavailable,
			delay:      0,
			want:       false,
		},
		{
			name:       "service_timeout",
			statusCode: http.StatusOK,
			delay:      3 * time.Second, // Longer than 2s timeout
			want:       false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if tt.delay > 0 {
					time.Sleep(tt.delay)
				}
				w.WriteHeader(tt.statusCode)
			}))
			defer server.Close()

			provider, err := NewOllamaProvider(&Config{
				Endpoint:       server.URL,
				Model:          "test-model",
				NumCases:       3,
				MaxRetries:     3,
				RequestTimeout: 60,
				HealthTimeout:  2,
			})
			if err != nil {
				t.Fatalf("NewOllamaProvider() error = %v", err)
			}

			got := provider.IsAvailable()
			if got != tt.want {
				t.Errorf("IsAvailable() = %v, want %v", got, tt.want)
			}
		})
	}

	// Test with invalid endpoint
	t.Run("invalid_endpoint", func(t *testing.T) {
		provider, err := NewOllamaProvider(&Config{
			Endpoint:       "http://invalid-host-that-does-not-exist:9999",
			Model:          "test-model",
			NumCases:       3,
			MaxRetries:     3,
			RequestTimeout: 60,
			HealthTimeout:  2,
		})
		if err != nil {
			t.Fatalf("NewOllamaProvider() error = %v", err)
		}

		if provider.IsAvailable() {
			t.Error("IsAvailable() = true for invalid endpoint, want false")
		}
	})
}

func TestOllamaProvider_generateGo(t *testing.T) {
	tests := []struct {
		name           string
		responseBody   map[string]interface{}
		statusCode     int
		wantErr        bool
		wantNumCases   int
		validateResult func(*testing.T, []TestCase)
	}{
		{
			name: "successful_generation",
			responseBody: map[string]interface{}{
				"response": `package testdata

import "testing"

type args struct {
	a int
	b int
}

func TestAdd(t *testing.T) {
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "positive_numbers",
			args: args{a: 2, b: 3},
			want: 5,
		},
		{
			name: "negative_numbers",
			args: args{a: -2, b: -3},
			want: -5,
		},
	}
}`,
			},
			statusCode:   http.StatusOK,
			wantErr:      false,
			wantNumCases: 2,
			validateResult: func(t *testing.T, cases []TestCase) {
				if len(cases) != 2 {
					t.Errorf("got %d cases, want 2", len(cases))
					return
				}
				if cases[0].Name != "positive_numbers" {
					t.Errorf("case[0].Name = %q, want %q", cases[0].Name, "positive_numbers")
				}
				if cases[0].Args["a"] != "2" {
					t.Errorf("case[0].Args[a] = %q, want %q", cases[0].Args["a"], "2")
				}
			},
		},
		{
			name: "http_error",
			responseBody: map[string]interface{}{
				"error": "model not found",
			},
			statusCode: http.StatusNotFound,
			wantErr:    true,
		},
		{
			name: "invalid_json_response",
			responseBody: map[string]interface{}{
				"response": "not valid go code {{{",
			},
			statusCode: http.StatusOK,
			wantErr:    true,
		},
		{
			name: "response_with_markdown",
			responseBody: map[string]interface{}{
				"response": "```go\npackage testdata\n\nimport \"testing\"\n\ntype args struct { x int }\n\nfunc Test(t *testing.T) {\n\ttests := []struct {\n\t\tname string\n\t\targs args\n\t\twant int\n\t}{\n\t\t{name: \"test1\", args: args{x: 10}, want: 20},\n\t}\n}\n```",
			},
			statusCode:   http.StatusOK,
			wantErr:      false,
			wantNumCases: 1,
			validateResult: func(t *testing.T, cases []TestCase) {
				if len(cases) != 1 {
					t.Errorf("got %d cases, want 1", len(cases))
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				// Verify request
				if r.Method != "POST" {
					t.Errorf("method = %q, want POST", r.Method)
				}
				if r.URL.Path != "/api/generate" {
					t.Errorf("path = %q, want /api/generate", r.URL.Path)
				}

				w.WriteHeader(tt.statusCode)
				json.NewEncoder(w).Encode(tt.responseBody)
			}))
			defer server.Close()

			provider, err := NewOllamaProvider(&Config{
				Endpoint:       server.URL,
				Model:          "test-model",
				NumCases:       3,
				MaxRetries:     3,
				RequestTimeout: 60,
				HealthTimeout:  2,
			})
			if err != nil {
				t.Fatalf("NewOllamaProvider() error = %v", err)
			}

			ctx := context.Background()
			cases, err := provider.generateGo(ctx, "test prompt")

			if (err != nil) != tt.wantErr {
				t.Errorf("generateGo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				if len(cases) != tt.wantNumCases {
					t.Errorf("got %d cases, want %d", len(cases), tt.wantNumCases)
				}
				if tt.validateResult != nil {
					tt.validateResult(t, cases)
				}
			}
		})
	}
}

func TestOllamaProvider_GenerateTestCases(t *testing.T) {
	t.Run("successful_with_valid_response", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			response := map[string]interface{}{
				"response": `package testdata

import "testing"

type args struct {
	a int
	b int
}

func TestAdd(t *testing.T) {
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "positive_numbers",
			args: args{a: 2, b: 3},
			want: 5,
		},
		{
			name: "negative_numbers",
			args: args{a: -2, b: -3},
			want: -5,
		},
	}
}`,
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(response)
		}))
		defer server.Close()

		provider, err := NewOllamaProvider(&Config{
			Endpoint:       server.URL,
			Model:          "test-model",
			NumCases:       3,
			MaxRetries:     3,
			RequestTimeout: 60,
			HealthTimeout:  2,
		})
		if err != nil {
			t.Fatalf("NewOllamaProvider() error = %v", err)
		}

		fn := &models.Function{
			Name: "Add",
			Parameters: []*models.Field{
				{Name: "a", Type: &models.Expression{Value: "int"}},
				{Name: "b", Type: &models.Expression{Value: "int"}},
			},
			Results: []*models.Field{
				{Type: &models.Expression{Value: "int"}},
			},
		}

		ctx := context.Background()
		cases, err := provider.GenerateTestCases(ctx, fn)

		if err != nil {
			t.Fatalf("GenerateTestCases() error = %v", err)
		}

		if len(cases) != 2 {
			t.Errorf("got %d cases, want 2", len(cases))
		}
	})

	t.Run("retry_on_validation_failure", func(t *testing.T) {
		attemptCount := 0
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			attemptCount++
			var response string

			// First attempt: missing required field
			if attemptCount == 1 {
				response = `package testdata
import "testing"
type args struct { x int }
func Test(t *testing.T) {
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test1", want: 10},
	}
}`
			} else {
				// Second attempt: correct
				response = `package testdata
import "testing"
type args struct { x int }
func Test(t *testing.T) {
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "valid", args: args{x: 5}, want: 10},
	}
}`
			}

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]interface{}{"response": response})
		}))
		defer server.Close()

		provider, err := NewOllamaProvider(&Config{
			Endpoint:       server.URL,
			Model:          "test-model",
			NumCases:       3,
			MaxRetries:     3,
			RequestTimeout: 60,
			HealthTimeout:  2,
		})
		if err != nil {
			t.Fatalf("NewOllamaProvider() error = %v", err)
		}

		fn := &models.Function{
			Name: "Double",
			Parameters: []*models.Field{
				{Name: "x", Type: &models.Expression{Value: "int"}},
			},
			Results: []*models.Field{
				{Type: &models.Expression{Value: "int"}},
			},
		}

		ctx := context.Background()
		cases, err := provider.GenerateTestCases(ctx, fn)

		if err != nil {
			t.Fatalf("GenerateTestCases() error = %v", err)
		}

		if attemptCount != 2 {
			t.Errorf("attemptCount = %d, want 2 (should retry on validation failure)", attemptCount)
		}

		if len(cases) != 1 {
			t.Errorf("got %d cases, want 1", len(cases))
		}
	})

	t.Run("fails_after_max_retries", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Always return invalid response
			response := map[string]interface{}{
				"response": "invalid go code {{{",
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(response)
		}))
		defer server.Close()

		provider, err := NewOllamaProvider(&Config{
			Endpoint:       server.URL,
			Model:          "test-model",
			NumCases:       3,
			MaxRetries:     3,
			RequestTimeout: 60,
			HealthTimeout:  2,
		})
		if err != nil {
			t.Fatalf("NewOllamaProvider() error = %v", err)
		}

		fn := &models.Function{
			Name:       "Foo",
			Parameters: []*models.Field{},
			Results:    []*models.Field{},
		}

		ctx := context.Background()
		_, err2 := provider.GenerateTestCases(ctx, fn)

		if err2 == nil {
			t.Error("GenerateTestCases() expected error after max retries, got nil")
		}
	})
}

func Test_validateTestCases(t *testing.T) {
	tests := []struct {
		name    string
		cases   []TestCase
		fn      *models.Function
		wantErr bool
		errMsg  string
	}{
		{
			name: "valid_test_cases",
			cases: []TestCase{
				{
					Name: "test1",
					Args: map[string]string{"a": "1", "b": "2"},
					Want: map[string]string{"want": "3"},
				},
			},
			fn: &models.Function{
				Parameters: []*models.Field{
					{Name: "a", Type: &models.Expression{Value: "int"}},
					{Name: "b", Type: &models.Expression{Value: "int"}},
				},
				Results: []*models.Field{
					{Type: &models.Expression{Value: "int"}},
				},
			},
			wantErr: false,
		},
		{
			name:  "empty_test_cases",
			cases: []TestCase{},
			fn: &models.Function{
				Parameters: []*models.Field{},
				Results:    []*models.Field{},
			},
			wantErr: true,
			errMsg:  "no test cases generated",
		},
		{
			name: "missing_test_name",
			cases: []TestCase{
				{
					Name: "",
					Args: map[string]string{},
					Want: map[string]string{},
				},
			},
			fn: &models.Function{
				Parameters: []*models.Field{},
				Results:    []*models.Field{},
			},
			wantErr: true,
			errMsg:  "missing name",
		},
		{
			name: "missing_required_argument",
			cases: []TestCase{
				{
					Name: "test1",
					Args: map[string]string{"a": "1"},
					Want: map[string]string{"want": "3"},
				},
			},
			fn: &models.Function{
				Parameters: []*models.Field{
					{Name: "a", Type: &models.Expression{Value: "int"}},
					{Name: "b", Type: &models.Expression{Value: "int"}},
				},
				Results: []*models.Field{
					{Type: &models.Expression{Value: "int"}},
				},
			},
			wantErr: true,
			errMsg:  "missing argument: b",
		},
		{
			name: "wrong_number_of_return_values",
			cases: []TestCase{
				{
					Name: "test1",
					Args: map[string]string{"a": "1"},
					Want: map[string]string{"want": "3", "want1": "4"},
				},
			},
			fn: &models.Function{
				Parameters: []*models.Field{
					{Name: "a", Type: &models.Expression{Value: "int"}},
				},
				Results: []*models.Field{
					{Type: &models.Expression{Value: "int"}},
				},
			},
			wantErr: true,
			errMsg:  "has 2 return values, expected 1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateTestCases(tt.cases, tt.fn)

			if (err != nil) != tt.wantErr {
				t.Errorf("validateTestCases() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr && tt.errMsg != "" {
				if err == nil || !contains(err.Error(), tt.errMsg) {
					t.Errorf("validateTestCases() error = %v, want error containing %q", err, tt.errMsg)
				}
			}
		})
	}
}

func Test_parseTestCases(t *testing.T) {
	tests := []struct {
		name     string
		response string
		maxCases int
		wantLen  int
		wantErr  bool
	}{
		{
			name: "valid_json_array",
			response: `[
				{"name": "test1", "args": {"x": 10}, "want": {"result": 20}, "wantErr": false},
				{"name": "test2", "args": {"x": 5}, "want": {"result": 10}, "wantErr": true}
			]`,
			maxCases: 5,
			wantLen:  2,
			wantErr:  false,
		},
		{
			name: "json_with_extra_text",
			response: `Here are the test cases:
[
	{"name": "test1", "args": {"x": 10}, "want": {"result": 20}}
]
That's all!`,
			maxCases: 5,
			wantLen:  1,
			wantErr:  false,
		},
		{
			name:     "no_json_array",
			response: "No valid JSON here",
			maxCases: 5,
			wantLen:  0,
			wantErr:  true,
		},
		{
			name: "max_cases_limit",
			response: `[
				{"name": "test1", "want": {"result": 1}},
				{"name": "test2", "want": {"result": 2}},
				{"name": "test3", "want": {"result": 3}},
				{"name": "test4", "want": {"result": 4}}
			]`,
			maxCases: 2,
			wantLen:  2,
			wantErr:  false,
		},
		{
			name: "uses_expected_instead_of_want",
			response: `[
				{"name": "test1", "args": {"x": 10}, "expected": {"result": 20}}
			]`,
			maxCases: 5,
			wantLen:  1,
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cases, err := parseTestCases(tt.response, tt.maxCases)

			if (err != nil) != tt.wantErr {
				t.Errorf("parseTestCases() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && len(cases) != tt.wantLen {
				t.Errorf("parseTestCases() returned %d cases, want %d", len(cases), tt.wantLen)
			}
		})
	}
}

// Helper function to check if a string contains a substring
func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(substr) == 0 ||
		(len(s) > 0 && len(substr) > 0 && indexOf(s, substr) >= 0))
}

func indexOf(s, substr string) int {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return i
		}
	}
	return -1
}
