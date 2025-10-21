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

// buildGoPrompt creates a prompt asking the LLM to generate Go test cases directly.
func buildGoPrompt(fn *models.Function, scaffold string, numCases int, previousError string) string {
	var sb strings.Builder

	sb.WriteString("You are a Go testing expert. Generate test cases for the following function.\n\n")

	// Function with body
	sb.WriteString("Function to test:\n```go\n")
	sb.WriteString(buildFunctionSignature(fn))
	if fn.Body != "" {
		sb.WriteString(" ")
		sb.WriteString(fn.Body)
	}
	sb.WriteString("\n```\n\n")

	// Show the test scaffold
	sb.WriteString("Here is the test scaffold that was generated:\n```go\n")
	sb.WriteString(scaffold)
	sb.WriteString("```\n\n")

	// Instructions
	sb.WriteString(fmt.Sprintf("Generate %d meaningful test cases.\n", numCases))
	sb.WriteString("Return ONLY the test cases array in Go syntax, replacing the `// TODO: Add test cases.` line.\n\n")

	// Example
	sb.WriteString("EXAMPLE:\n")
	sb.WriteString("If the scaffold contains:\n")
	sb.WriteString("```go\n")
	sb.WriteString("tests := []struct {\n")
	sb.WriteString("    name string\n")
	sb.WriteString("    args args\n")
	sb.WriteString("    want int\n")
	sb.WriteString("} {\n")
	sb.WriteString("    // TODO: Add test cases.\n")
	sb.WriteString("}\n")
	sb.WriteString("```\n\n")

	sb.WriteString("You should return:\n")
	sb.WriteString("```go\n")
	sb.WriteString("{\n")
	sb.WriteString("    name: \"positive_numbers\",\n")
	sb.WriteString("    args: args{a: 5, b: 3},\n")
	sb.WriteString("    want: 8,\n")
	sb.WriteString("},\n")
	sb.WriteString("{\n")
	sb.WriteString("    name: \"zero_values\",\n")
	sb.WriteString("    args: args{a: 0, b: 0},\n")
	sb.WriteString("    want: 0,\n")
	sb.WriteString("},\n")
	sb.WriteString("```\n\n")

	// Error handling example if needed
	if fn.ReturnsError {
		sb.WriteString("For functions returning errors, include wantErr:\n")
		sb.WriteString("```go\n")
		sb.WriteString("{\n")
		sb.WriteString("    name: \"division_by_zero\",\n")
		sb.WriteString("    args: args{a: 10, b: 0},\n")
		sb.WriteString("    want: 0,\n")
		sb.WriteString("    wantErr: true,\n")
		sb.WriteString("},\n")
		sb.WriteString("```\n\n")
	}

	sb.WriteString("IMPORTANT:\n")
	sb.WriteString("- Return ONLY the test case structs, not the full test function\n")
	sb.WriteString("- Do NOT include 'tests := []struct{...}{' or the closing '}'\n")
	sb.WriteString("- Each test case should end with a comma\n")
	sb.WriteString("- Use the exact field names from the scaffold\n")
	sb.WriteString("- Use valid Go literal syntax\n\n")

	// Add error feedback if retrying
	if previousError != "" {
		sb.WriteString("PREVIOUS ATTEMPT FAILED:\n")
		sb.WriteString(previousError)
		sb.WriteString("\n\nPlease fix the above issue.\n\n")
	}

	sb.WriteString("Now generate the test cases:\n")

	return sb.String()
}
