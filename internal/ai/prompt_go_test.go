package ai

import (
	"strings"
	"testing"

	"github.com/cweill/gotests/internal/models"
)

func Test_buildTestScaffold(t *testing.T) {
	tests := []struct {
		name     string
		fn       *models.Function
		wantName bool
		wantArgs bool
		wantWant bool
		wantErr  bool
	}{
		{
			name: "simple_function_no_params_no_results",
			fn: &models.Function{
				Name:       "Foo",
				Parameters: []*models.Field{},
				Results:    []*models.Field{},
			},
			wantName: true,
			wantArgs: false,
			wantWant: false,
			wantErr:  false,
		},
		{
			name: "function_with_parameters",
			fn: &models.Function{
				Name: "Add",
				Parameters: []*models.Field{
					{Name: "a", Type: &models.Expression{Value: "int"}},
					{Name: "b", Type: &models.Expression{Value: "int"}},
				},
				Results: []*models.Field{
					{Type: &models.Expression{Value: "int"}},
				},
			},
			wantName: true,
			wantArgs: true,
			wantWant: true,
			wantErr:  false,
		},
		{
			name: "function_returning_error",
			fn: &models.Function{
				Name: "Validate",
				Parameters: []*models.Field{
					{Name: "email", Type: &models.Expression{Value: "string"}},
				},
				Results: []*models.Field{
					{Type: &models.Expression{Value: "error"}},
				},
				ReturnsError: true,
			},
			wantName: true,
			wantArgs: true,
			wantWant: false,
			wantErr:  true,
		},
		{
			name: "function_with_receiver",
			fn: &models.Function{
				Name: "Method",
				Receiver: &models.Receiver{
					Field: &models.Field{
						Name: "c",
						Type: &models.Expression{Value: "*Calculator"},
					},
				},
				Parameters: []*models.Field{
					{Name: "n", Type: &models.Expression{Value: "int"}},
				},
				Results: []*models.Field{
					{Type: &models.Expression{Value: "int"}},
				},
			},
			wantName: true,
			wantArgs: true,
			wantWant: true,
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := buildTestScaffold(tt.fn)

			// Check for expected components
			if tt.wantName && !strings.Contains(got, "name string") {
				t.Errorf("buildTestScaffold() missing 'name string' field")
			}

			if tt.wantArgs && !strings.Contains(got, "args args") {
				t.Errorf("buildTestScaffold() missing 'args args' field")
			}

			if tt.wantWant && !strings.Contains(got, "want ") {
				t.Errorf("buildTestScaffold() missing 'want' field")
			}

			if tt.wantErr && !strings.Contains(got, "wantErr bool") {
				t.Errorf("buildTestScaffold() missing 'wantErr bool' field")
			}

			// Should always have test struct definition
			if !strings.Contains(got, "tests := []struct {") {
				t.Errorf("buildTestScaffold() missing test struct definition")
			}

			// Should have TODO placeholder
			if !strings.Contains(got, "// TODO: Add test cases.") {
				t.Errorf("buildTestScaffold() missing TODO placeholder")
			}
		})
	}
}

func Test_buildGoPrompt(t *testing.T) {
	tests := []struct {
		name          string
		fn            *models.Function
		scaffold      string
		minCases      int
		maxCases      int
		previousError string
		wantContains  []string
	}{
		{
			name: "simple_function",
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
			scaffold: "tests := []struct {\n    name string\n}",
			minCases: 3,
			maxCases: 3,
			wantContains: []string{
				"You are a Go testing expert",
				"Generate 3 test cases. Each test case must have UNIQUE, DIFFERENT input values.",
				"Function to test:",
				"return a + b",
				"test scaffold:",
				"Requirements:",
				"Use NAMED fields",
				"a: <value>",
				"b: <value>",
			},
		},
		{
			name: "function_with_error_return",
			fn: &models.Function{
				Name: "Validate",
				Parameters: []*models.Field{
					{Name: "email", Type: &models.Expression{Value: "string"}},
				},
				Results: []*models.Field{
					{Type: &models.Expression{Value: "error"}},
				},
				ReturnsError: true,
				Body:         "{\n    if email == \"\" {\n        return errors.New(\"empty\")\n    }\n    return nil\n}",
			},
			scaffold: "tests := []struct {\n    name string\n    wantErr bool\n}",
			minCases: 2,
			maxCases: 2,
			wantContains: []string{
				"Generate 2 test cases. Each test case must have UNIQUE, DIFFERENT input values.",
				"wantErr: false",
				"Include: 1 valid case, 1 edge case, 1 error case.",
			},
		},
		{
			name: "with_previous_error",
			fn: &models.Function{
				Name:       "Foo",
				Parameters: []*models.Field{},
				Results:    []*models.Field{},
			},
			scaffold:      "tests := []struct {}",
			minCases:      3,
			maxCases:      3,
			previousError: "missing field: name",
			wantContains: []string{
				"PREVIOUS ATTEMPT FAILED:",
				"missing field: name",
				"Please fix the above issue",
			},
		},
		{
			name: "function_with_receiver",
			fn: &models.Function{
				Name: "Multiply",
				Receiver: &models.Receiver{
					Field: &models.Field{
						Name: "c",
						Type: &models.Expression{Value: "*Calculator"},
					},
				},
				Parameters: []*models.Field{
					{Name: "n", Type: &models.Expression{Value: "int"}},
					{Name: "d", Type: &models.Expression{Value: "int"}},
				},
				Results: []*models.Field{
					{Type: &models.Expression{Value: "int"}},
				},
				Body: "{\n    return n * d\n}",
			},
			scaffold: "tests := []struct {\n    name string\n    c *Calculator\n}",
			minCases: 3,
			maxCases: 3,
			wantContains: []string{
				"return n * d",
				"c: &Calculator{}",
				"n: <value>",
				"d: <value>",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := buildGoPrompt(tt.fn, tt.scaffold, tt.minCases, tt.maxCases, tt.previousError)

			for _, want := range tt.wantContains {
				if !strings.Contains(got, want) {
					t.Errorf("buildGoPrompt() missing expected string: %q", want)
				}
			}
		})
	}
}

func Test_buildFunctionSignature(t *testing.T) {
	tests := []struct {
		name string
		fn   *models.Function
		want string
	}{
		{
			name: "simple_function",
			fn: &models.Function{
				Name: "Add",
				Parameters: []*models.Field{
					{Name: "a", Type: &models.Expression{Value: "int"}},
					{Name: "b", Type: &models.Expression{Value: "int"}},
				},
				Results: []*models.Field{
					{Type: &models.Expression{Value: "int"}},
				},
			},
			want: "func Add(a int, b int) int",
		},
		{
			name: "method_with_receiver",
			fn: &models.Function{
				Name: "Method",
				Receiver: &models.Receiver{
					Field: &models.Field{
						Name: "c",
						Type: &models.Expression{Value: "*Calculator"},
					},
				},
				Parameters: []*models.Field{
					{Name: "n", Type: &models.Expression{Value: "int"}},
				},
				Results: []*models.Field{
					{Type: &models.Expression{Value: "int"}},
				},
			},
			want: "func (c *Calculator) Method(n int) int",
		},
		{
			name: "function_with_error",
			fn: &models.Function{
				Name: "Validate",
				Parameters: []*models.Field{
					{Name: "s", Type: &models.Expression{Value: "string"}},
				},
				Results: []*models.Field{
					{Type: &models.Expression{Value: "error"}},
				},
			},
			want: "func Validate(s string) error",
		},
		{
			name: "function_no_params_no_results",
			fn: &models.Function{
				Name:       "Foo",
				Parameters: []*models.Field{},
				Results:    []*models.Field{},
			},
			want: "func Foo()",
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
