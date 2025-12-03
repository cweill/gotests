package ai

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"regexp"
	"strings"
)

// extractCodeFromMarkdown removes markdown code blocks and cleans up the code.
func extractCodeFromMarkdown(text string) string {
	// Remove ```go and ``` markers (non-greedy to handle multiple blocks)
	re := regexp.MustCompile("(?s)```go\\s*(.*?)```")
	matches := re.FindStringSubmatch(text)
	if len(matches) > 1 {
		code := strings.TrimSpace(matches[1])
		code = ensureTrailingComma(code)
		code = fixCommonSyntaxIssues(code)
		return code
	}

	// Try without the 'go' language specifier (non-greedy)
	re = regexp.MustCompile("(?s)```\\s*(.*?)```")
	matches = re.FindStringSubmatch(text)
	if len(matches) > 1 {
		code := strings.TrimSpace(matches[1])
		code = ensureTrailingComma(code)
		code = fixCommonSyntaxIssues(code)
		return code
	}

	// No markdown blocks, return as-is
	code := strings.TrimSpace(text)
	code = ensureTrailingComma(code)
	code = fixCommonSyntaxIssues(code)
	return code
}

// fixCommonSyntaxIssues fixes common syntax mistakes LLMs make.
func fixCommonSyntaxIssues(code string) string {
	// Fix "} {" -> "}{" (space between closing brace and opening brace in struct literals)
	// This is a common LLM mistake: []struct { ... } { ... } should be []struct { ... }{ ... }
	re := regexp.MustCompile(`\}\s+\{`)
	code = re.ReplaceAllString(code, "}{")

	return code
}

// ensureTrailingComma adds a trailing comma to the last struct if missing.
// Only applies when parsing standalone test case arrays, not variable declarations or functions.
func ensureTrailingComma(code string) string {
	code = strings.TrimSpace(code)

	// Don't modify complete functions or variable declarations
	if strings.Contains(code, "func Test") || strings.Contains(code, "tests :=") || strings.Contains(code, "tests=") {
		return code
	}

	// If code ends with }, add a comma (for test case arrays only)
	// This is only for the old parseTestCaseArray format
	if strings.HasSuffix(code, "}") && !strings.HasSuffix(code, "},") {
		return code + ","
	}
	return code
}

// parseGoTestCases extracts test cases from various LLM output formats:
// 1. Complete test function: func TestAdd(t *testing.T) { ... }
// 2. Variable declaration: tests := []struct{...}{...}
// 3. Just test cases: { name: "...", ... }, ...
func parseGoTestCases(goCode string, maxCases int) ([]TestCase, error) {
	// Extract code from markdown code blocks if present
	cleaned := extractCodeFromMarkdown(goCode)

	// Check if this is a complete function
	if strings.Contains(cleaned, "func Test") {
		return parseCompleteTestFunction(cleaned, maxCases)
	}

	// Check if this is a variable declaration (tests := []struct{...}{...})
	if strings.Contains(cleaned, "tests :=") || strings.Contains(cleaned, "tests=") {
		return parseTestsVariableDeclaration(cleaned, maxCases)
	}

	// Fallback: parse as just test case array (old approach)
	return parseTestCaseArray(cleaned, maxCases)
}

// parseTestsVariableDeclaration handles LLM output like:
// tests := []struct { name string; args args; want int }{
//     { name: "test1", args: args{a: 5}, want: 8 },
// }
func parseTestsVariableDeclaration(code string, maxCases int) ([]TestCase, error) {
	// Debug: uncomment to see the code being parsed
	// fmt.Printf("DEBUG parseTestsVariableDeclaration input:\n%s\n---\n", code)
	
	// Wrap in a package with type definitions to make it valid Go
	// We define 'args' as an empty struct since we only care about extracting values
	wrapped := fmt.Sprintf(`package main

// Dummy type definitions to satisfy the parser
type args struct{}
type fields struct{}

func init() {
	%s
}`, code)
	
	// Debug: print wrapped code
	// fmt.Printf("DEBUG wrapped code:\n%s\n---\n", wrapped)

	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "test.go", wrapped, parser.AllErrors)
	if err != nil {
		return nil, fmt.Errorf("parse Go code: %w", err)
	}

	var cases []TestCase
	ast.Inspect(file, func(n ast.Node) bool {
		// Look for composite literals (the test case structs)
		if cl, ok := n.(*ast.CompositeLit); ok {
			// Skip if this has a type (it's the outer []struct{...}{} or args{})
			if cl.Type != nil {
				return true
			}

			// This is a test case struct literal (no type, just {})
			tc := parseTestCase(cl)
			if tc != nil && len(cases) < maxCases {
				cases = append(cases, *tc)
			}
		}
		return true
	})

	if len(cases) == 0 {
		return nil, fmt.Errorf("no test cases found in response")
	}

	return cases, nil
}

// parseCompleteTestFunction extracts test cases from a complete test function.
func parseCompleteTestFunction(goCode string, maxCases int) ([]TestCase, error) {
	// Parse the complete function
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "test.go", goCode, parser.AllErrors)
	if err != nil {
		return nil, fmt.Errorf("parse Go code: %w", err)
	}

	// Find the test function and extract test cases
	var cases []TestCase
	ast.Inspect(file, func(n ast.Node) bool {
		// Look for composite literals (the test case structs)
		if cl, ok := n.(*ast.CompositeLit); ok {
			// Skip the outer array literal (the tests := []struct{...}{...})
			if _, isArray := cl.Type.(*ast.ArrayType); isArray {
				return true
			}

			// Skip the struct type definition
			if _, isStruct := cl.Type.(*ast.StructType); isStruct {
				return true
			}

			// This should be a test case struct literal (no type, just {})
			if cl.Type == nil {
				tc := parseTestCase(cl)
				if tc != nil && len(cases) < maxCases {
					cases = append(cases, *tc)
				}
			}
		}
		return true
	})

	if len(cases) == 0 {
		return nil, fmt.Errorf("no test cases found in response")
	}

	return cases, nil
}

// parseTestCaseArray extracts test cases from just the test case array.
// This is the old approach for backward compatibility.
func parseTestCaseArray(testCases string, maxCases int) ([]TestCase, error) {
	// Wrap the test cases in a valid Go structure so we can parse it
	wrapped := fmt.Sprintf(`package main
func init() {
	_ = []struct {
		name string
		args interface{}
		want interface{}
		wantErr bool
	}{
%s
	}
}`, testCases)

	// Parse the wrapped code
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "test.go", wrapped, parser.AllErrors)
	if err != nil {
		return nil, fmt.Errorf("parse Go code: %w", err)
	}

	// Extract the composite literal from the AST
	var cases []TestCase
	ast.Inspect(file, func(n ast.Node) bool {
		// Look for composite literals (the test case structs)
		if cl, ok := n.(*ast.CompositeLit); ok {
			// Skip the outer array literal
			if _, isArray := cl.Type.(*ast.ArrayType); isArray {
				return true
			}

			// This should be a test case struct
			tc := parseTestCase(cl)
			if tc != nil && len(cases) < maxCases {
				cases = append(cases, *tc)
			}
		}
		return true
	})

	if len(cases) == 0 {
		return nil, fmt.Errorf("no test cases found in response")
	}

	return cases, nil
}

// parseTestCase extracts a TestCase from a composite literal AST node.
func parseTestCase(cl *ast.CompositeLit) *TestCase {
	tc := &TestCase{
		Args: make(map[string]string),
		Want: make(map[string]string),
	}

	for _, elt := range cl.Elts {
		kv, ok := elt.(*ast.KeyValueExpr)
		if !ok {
			// Positional struct literal without keys - skip for now
			continue
		}

		key, ok := kv.Key.(*ast.Ident)
		if !ok {
			continue
		}

		switch key.Name {
		case "name":
			if lit, ok := kv.Value.(*ast.BasicLit); ok {
				tc.Name = strings.Trim(lit.Value, `"`)
			}
		case "args":
			if comp, ok := kv.Value.(*ast.CompositeLit); ok {
				tc.Args = parseArgsStruct(comp)
			}
		case "want":
			tc.Want["want"] = exprToString(kv.Value)
		case "wantErr":
			if id, ok := kv.Value.(*ast.Ident); ok {
				tc.WantErr = id.Name == "true"
			}
		case "c", "receiver": // Handle receiver field
			// Skip receiver field
		default:
			// Handle want1, want2, etc. or other result fields
			if strings.HasPrefix(key.Name, "want") && key.Name != "wantErr" {
				tc.Want[key.Name] = exprToString(kv.Value)
			}
		}
	}

	if tc.Name == "" {
		return nil
	}

	return tc
}

// parseArgsStruct extracts argument values from an args{} composite literal.
func parseArgsStruct(cl *ast.CompositeLit) map[string]string {
	args := make(map[string]string)

	for _, elt := range cl.Elts {
		kv, ok := elt.(*ast.KeyValueExpr)
		if !ok {
			continue
		}

		key, ok := kv.Key.(*ast.Ident)
		if !ok {
			continue
		}

		args[key.Name] = exprToString(kv.Value)
	}

	return args
}

// exprToString converts an AST expression to its string representation.
func exprToString(expr ast.Expr) string {
	switch e := expr.(type) {
	case *ast.BasicLit:
		return e.Value
	case *ast.Ident:
		return e.Name
	case *ast.UnaryExpr:
		return e.Op.String() + exprToString(e.X)
	case *ast.BinaryExpr:
		return exprToString(e.X) + " " + e.Op.String() + " " + exprToString(e.Y)
	case *ast.CompositeLit:
		// For complex types (structs/maps/slices), convert AST back to source code
		var buf bytes.Buffer
		if err := printer.Fprint(&buf, token.NewFileSet(), e); err != nil {
			return "nil" // Fallback on error
		}
		return buf.String()
	default:
		return "nil"
	}
}
