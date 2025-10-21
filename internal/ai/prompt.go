package ai

import (
	"fmt"
	"strings"

	"github.com/cweill/gotests/internal/models"
)

// buildPrompt creates a prompt for the LLM to generate test cases.
func buildPrompt(fn *models.Function, numCases int, previousError string) string {
	var sb strings.Builder

	// System instruction
	sb.WriteString("You are a Go testing expert. Generate test cases for the following function.\n\n")

	// Function signature
	sb.WriteString("Function to test:\n```go\n")
	sb.WriteString(buildFunctionSignature(fn))
	sb.WriteString("\n```\n\n")

	// Instructions
	sb.WriteString(fmt.Sprintf("Generate %d test cases that cover:\n", numCases))
	sb.WriteString("1. Normal/typical inputs\n")
	sb.WriteString("2. Edge cases (zero values, empty strings, nil pointers)\n")
	sb.WriteString("3. Boundary conditions\n")
	if fn.ReturnsError {
		sb.WriteString("4. Error conditions\n")
	}
	sb.WriteString("\n")

	// Output format
	sb.WriteString("Return ONLY a JSON array with this exact structure:\n")
	sb.WriteString("[\n")
	sb.WriteString("  {\n")
	sb.WriteString("    \"name\": \"test_case_name\",\n")
	sb.WriteString("    \"description\": \"what this tests\",\n")
	sb.WriteString("    \"args\": {\n")

	// Show expected arg structure
	for i, param := range fn.TestParameters() {
		if i > 0 {
			sb.WriteString(",\n")
		}
		sb.WriteString(fmt.Sprintf("      \"%s\": <value>", param.Name))
	}
	sb.WriteString("\n    },\n")

	// Show expected return structure
	sb.WriteString("    \"want\": {\n")
	for i, result := range fn.TestResults() {
		if i > 0 {
			sb.WriteString(",\n")
		}
		resultName := result.Name
		if resultName == "" {
			resultName = "result"
		}
		sb.WriteString(fmt.Sprintf("      \"%s\": <expected_value>", resultName))
	}
	sb.WriteString("\n    }")

	if fn.ReturnsError {
		sb.WriteString(",\n    \"wantErr\": true or false")
	}

	sb.WriteString("\n  }\n")
	sb.WriteString("]\n\n")

	// Important notes
	sb.WriteString("IMPORTANT:\n")
	sb.WriteString("- Use valid Go literal syntax for values (e.g., \"hello\", 42, true, nil)\n")
	sb.WriteString("- For strings, use double quotes\n")
	sb.WriteString("- For zero values: 0 for int, \"\" for string, false for bool, nil for pointers\n")
	sb.WriteString("- Return ONLY the JSON array, no other text\n")

	// Add error feedback if retrying
	if previousError != "" {
		sb.WriteString("\n")
		sb.WriteString("PREVIOUS ATTEMPT FAILED:\n")
		sb.WriteString(previousError)
		sb.WriteString("\n\nPlease fix the above issue.\n")
	}

	return sb.String()
}

// buildFunctionSignature creates a Go function signature string.
func buildFunctionSignature(fn *models.Function) string {
	var sb strings.Builder

	// Receiver
	if fn.Receiver != nil {
		sb.WriteString(fmt.Sprintf("func (%s %s) ", fn.Receiver.Name, fn.Receiver.Type.String()))
	} else {
		sb.WriteString("func ")
	}

	// Function name
	sb.WriteString(fn.Name)

	// Parameters
	sb.WriteString("(")
	for i, param := range fn.Parameters {
		if i > 0 {
			sb.WriteString(", ")
		}
		if param.Name != "" {
			sb.WriteString(param.Name)
			sb.WriteString(" ")
		}
		sb.WriteString(param.Type.String())
	}
	sb.WriteString(")")

	// Return types
	if len(fn.Results) > 0 || fn.ReturnsError {
		sb.WriteString(" ")
		hasMultiple := len(fn.Results) > 1 || (len(fn.Results) == 1 && fn.ReturnsError)

		if hasMultiple {
			sb.WriteString("(")
		}

		for i, result := range fn.Results {
			if i > 0 {
				sb.WriteString(", ")
			}
			sb.WriteString(result.Type.String())
		}

		if fn.ReturnsError {
			if len(fn.Results) > 0 {
				sb.WriteString(", ")
			}
			sb.WriteString("error")
		}

		if hasMultiple {
			sb.WriteString(")")
		}
	}

	return sb.String()
}
