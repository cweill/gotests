package ai

import (
	"fmt"
	"strings"

	"github.com/cweill/gotests/internal/models"
)

// buildTestScaffold generates a minimal test scaffold to show the LLM.
func buildTestScaffold(fn *models.Function) string {
	var sb strings.Builder

	// Build the test struct definition
	sb.WriteString("tests := []struct {\n")
	sb.WriteString("\tname string\n")

	// Receiver if present
	if fn.Receiver != nil {
		// Use helpers from render package to get consistent naming
		receiverName := "c" // simplified for example
		if fn.Receiver.Type != nil {
			sb.WriteString(fmt.Sprintf("\t%s %s\n", receiverName, fn.Receiver.Type.String()))
		}
	}

	// Args struct if has parameters
	if len(fn.TestParameters()) > 0 {
		sb.WriteString("\targs args\n")
	}

	// Want fields for results
	for _, result := range fn.TestResults() {
		wantName := "want"
		if result.Index > 0 {
			wantName = fmt.Sprintf("want%d", result.Index)
		}
		sb.WriteString(fmt.Sprintf("\t%s %s\n", wantName, result.Type.String()))
	}

	// WantErr if returns error
	if fn.ReturnsError {
		sb.WriteString("\twantErr bool\n")
	}

	sb.WriteString("} {\n")
	sb.WriteString("\t// TODO: Add test cases.\n")
	sb.WriteString("}\n")

	return sb.String()
}

// buildGoPrompt creates a prompt asking the LLM to generate a complete test function.
func buildGoPrompt(fn *models.Function, scaffold string, numCases int, previousError string) string {
	var sb strings.Builder

	sb.WriteString("You are a Go testing expert. Generate test cases for the following function.\n\n")

	// Function with body (limit size to prevent memory issues)
	sb.WriteString("Function to test:\n```go\n")
	sb.WriteString(buildFunctionSignature(fn))
	if fn.Body != "" {
		body := fn.Body
		// Truncate very large function bodies
		if len(body) > MaxFunctionBodySize {
			body = body[:MaxFunctionBodySize] + "\n// ... (truncated)"
		}
		sb.WriteString(" ")
		sb.WriteString(body)
	}
	sb.WriteString("\n```\n\n")

	// Show the test scaffold
	sb.WriteString("Here is the test scaffold:\n```go\n")
	sb.WriteString(scaffold)
	sb.WriteString("```\n\n")

	// Instructions
	sb.WriteString(fmt.Sprintf("Generate %d meaningful test cases.\n", numCases))
	sb.WriteString("Fill in the test cases using the EXACT struct format from the scaffold.\n")
	sb.WriteString("Use named fields (field: value) not positional values.\n\n")

	// Build a concrete example using the actual scaffold
	sb.WriteString("Fill in test cases like this (using the exact field names from your scaffold):\n")
	sb.WriteString("```go\n")
	sb.WriteString("{\n")
	sb.WriteString("    name: \"descriptive_test_name\",\n")

	// Show receiver if present
	if fn.Receiver != nil {
		receiverName := "c"
		if fn.Receiver.Type != nil {
			sb.WriteString(fmt.Sprintf("    %s: &%s{},\n", receiverName, strings.TrimPrefix(fn.Receiver.Type.String(), "*")))
		}
	}

	// Show args if present
	if len(fn.TestParameters()) > 0 {
		sb.WriteString("    args: args{")
		for i, param := range fn.TestParameters() {
			if i > 0 {
				sb.WriteString(", ")
			}
			sb.WriteString(fmt.Sprintf("%s: <value>", param.Name))
		}
		sb.WriteString("},\n")
	}

	// Show want fields
	for _, result := range fn.TestResults() {
		wantName := "want"
		if result.Index > 0 {
			wantName = fmt.Sprintf("want%d", result.Index)
		}
		sb.WriteString(fmt.Sprintf("    %s: <expected_value>,\n", wantName))
	}

	// Show wantErr if needed
	if fn.ReturnsError {
		sb.WriteString("    wantErr: false,\n")
	}

	sb.WriteString("},\n")
	sb.WriteString("```\n\n")

	sb.WriteString("IMPORTANT:\n")
	sb.WriteString("- Use NAMED FIELDS (field: value) not positional struct literals\n")
	sb.WriteString("- Use the EXACT field names shown in the scaffold above\n")
	sb.WriteString("- Generate realistic test values based on the function body\n")
	sb.WriteString("- Use valid Go literal syntax\n")
	if fn.ReturnsError {
		sb.WriteString("- For error returns, set wantErr: true or false appropriately\n")
	}
	sb.WriteString("\n")

	// Add error feedback if retrying
	if previousError != "" {
		sb.WriteString("PREVIOUS ATTEMPT FAILED:\n")
		sb.WriteString(previousError)
		sb.WriteString("\n\nPlease fix the above issue.\n\n")
	}

	sb.WriteString("Now generate the test cases:\n")

	return sb.String()
}
