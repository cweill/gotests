package ai

import (
	"strings"
	"testing"

	"github.com/cweill/gotests/internal/models"
)

func Test_buildPrompt(t *testing.T) {
	tests := []struct {
		name          string
		fn            *models.Function
		numCases      int
		previousError string
		wantContains  []string
	}{
		{
			name: "simple_function_no_error",
			fn: &models.Function{
				Name: "Add",
				Parameters: []*models.Field{
					{Name: "a", Type: &models.Expression{Value: "int"}},
					{Name: "b", Type: &models.Expression{Value: "int"}},
				},
				Results: []*models.Field{
					{Type: &models.Expression{Value: "int"}},
				},
				Body: "{\n    return a + b\n}",
			},
			numCases: 3,
			wantContains: []string{
				"You are a Go testing expert",
				"Generate 3 test cases",
				"func Add(a int, b int) int",
				"return a + b",
				"\"a\": <value>",
				"\"b\": <value>",
				"\"want\": <expected_value>",
				"Normal/typical inputs",
				"Edge cases",
			},
		},
		{
			name: "function_with_error_return",
			fn: &models.Function{
				Name: "Divide",
				Parameters: []*models.Field{
					{Name: "a", Type: &models.Expression{Value: "float64"}},
					{Name: "b", Type: &models.Expression{Value: "float64"}},
				},
				Results: []*models.Field{
					{Type: &models.Expression{Value: "float64"}},
				},
				ReturnsError: true,
				Body:         "{\n    if b == 0 {\n        return 0, errors.New(\"division by zero\")\n    }\n    return a / b, nil\n}",
			},
			numCases: 5,
			wantContains: []string{
				"Generate 5 test cases",
				"func Divide(a float64, b float64) (float64, error)",
				"Error conditions",
				"\"wantErr\": true or false",
				"EXAMPLE 2 (function with error)",
				"division_by_zero",
			},
		},
		{
			name: "function_with_receiver",
			fn: &models.Function{
				Name: "Process",
				Receiver: &models.Receiver{
					Field: &models.Field{
						Name: "s",
						Type: &models.Expression{Value: "*Service"},
					},
				},
				Parameters: []*models.Field{
					{Name: "data", Type: &models.Expression{Value: "string"}},
				},
				Results: []*models.Field{
					{Type: &models.Expression{Value: "string"}},
				},
				Body: "{\n    return strings.ToUpper(data)\n}",
			},
			numCases: 3,
			wantContains: []string{
				"func (s *Service) Process(data string) string",
				"\"data\": <value>",
				"strings.ToUpper",
			},
		},
		{
			name: "with_previous_error",
			fn: &models.Function{
				Name: "Foo",
				Parameters: []*models.Field{
					{Name: "x", Type: &models.Expression{Value: "int"}},
				},
				Results: []*models.Field{
					{Type: &models.Expression{Value: "int"}},
				},
			},
			numCases:      3,
			previousError: "missing field: name in test case",
			wantContains: []string{
				"PREVIOUS ATTEMPT FAILED:",
				"missing field: name in test case",
				"Please fix the above issue",
			},
		},
		{
			name: "function_multiple_return_values",
			fn: &models.Function{
				Name: "Split",
				Parameters: []*models.Field{
					{Name: "s", Type: &models.Expression{Value: "string"}},
				},
				Results: []*models.Field{
					{Type: &models.Expression{Value: "string"}},
					{Type: &models.Expression{Value: "string"}},
				},
			},
			numCases: 2,
			wantContains: []string{
				"func Split(s string) (string, string)",
				"\"want\": <expected_value>",
				"\"want1\": <expected_value>",
			},
		},
		{
			name: "function_no_parameters",
			fn: &models.Function{
				Name:       "GetTime",
				Parameters: []*models.Field{},
				Results: []*models.Field{
					{Type: &models.Expression{Value: "time.Time"}},
				},
				Body: "{\n    return time.Now()\n}",
			},
			numCases: 1,
			wantContains: []string{
				"func GetTime() time.Time",
				"return time.Now()",
				"Generate 1 test cases",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := buildPrompt(tt.fn, tt.numCases, tt.previousError)

			for _, want := range tt.wantContains {
				if !strings.Contains(got, want) {
					t.Errorf("buildPrompt() missing expected string: %q", want)
				}
			}

			// Verify it contains JSON format examples
			if !strings.Contains(got, "[") || !strings.Contains(got, "]") {
				t.Error("buildPrompt() should contain JSON array example")
			}

			// Verify it contains important notes
			if !strings.Contains(got, "IMPORTANT:") {
				t.Error("buildPrompt() should contain IMPORTANT notes")
			}
		})
	}
}

func Test_buildPrompt_JSONStructure(t *testing.T) {
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

	prompt := buildPrompt(fn, 3, "")

	// Check that it includes proper JSON structure guidance
	requiredElements := []string{
		"\"name\":",
		"\"description\":",
		"\"args\":",
		"\"want\":",
		"Return ONLY the JSON array",
		"valid Go literal syntax",
	}

	for _, elem := range requiredElements {
		if !strings.Contains(prompt, elem) {
			t.Errorf("buildPrompt() missing required JSON element: %q", elem)
		}
	}
}

func Test_buildFunctionSignature_Coverage(t *testing.T) {
	// This test is primarily to increase coverage for edge cases in buildFunctionSignature
	tests := []struct {
		name string
		fn   *models.Function
		want string
	}{
		{
			name: "no_params_no_results",
			fn: &models.Function{
				Name:       "DoNothing",
				Parameters: []*models.Field{},
				Results:    []*models.Field{},
			},
			want: "func DoNothing()",
		},
		{
			name: "single_unnamed_param",
			fn: &models.Function{
				Name: "Process",
				Parameters: []*models.Field{
					{Name: "", Type: &models.Expression{Value: "int"}},
				},
				Results: []*models.Field{},
			},
			want: "func Process(int)",
		},
		{
			name: "single_result_no_error",
			fn: &models.Function{
				Name:       "GetValue",
				Parameters: []*models.Field{},
				Results: []*models.Field{
					{Type: &models.Expression{Value: "int"}},
				},
				ReturnsError: false,
			},
			want: "func GetValue() int",
		},
		{
			name: "only_error_return",
			fn: &models.Function{
				Name:         "Validate",
				Parameters:   []*models.Field{},
				Results:      []*models.Field{},
				ReturnsError: true,
			},
			want: "func Validate() error",
		},
		{
			name: "multiple_results_with_error",
			fn: &models.Function{
				Name:       "Parse",
				Parameters: []*models.Field{},
				Results: []*models.Field{
					{Type: &models.Expression{Value: "int"}},
					{Type: &models.Expression{Value: "string"}},
				},
				ReturnsError: true,
			},
			want: "func Parse() (int, string, error)",
		},
		{
			name: "receiver_with_generics",
			fn: &models.Function{
				Name: "Get",
				Receiver: &models.Receiver{
					Field: &models.Field{
						Name: "c",
						Type: &models.Expression{Value: "*Cache[K, V]"},
					},
				},
				Parameters: []*models.Field{
					{Name: "key", Type: &models.Expression{Value: "K"}},
				},
				Results: []*models.Field{
					{Type: &models.Expression{Value: "V"}},
				},
			},
			want: "func (c *Cache[K, V]) Get(key K) V",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := buildFunctionSignature(tt.fn)
			if got != tt.want {
				t.Errorf("buildFunctionSignature() = %q, want %q", got, tt.want)
			}
		})
	}
}
