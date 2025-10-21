package ai

import (
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
)

// ValidateGeneratedTest checks if the generated test code compiles.
// Uses in-memory parsing and type-checking without writing files.
func ValidateGeneratedTest(testCode, pkgName string) error {
	fset := token.NewFileSet()

	// Parse the test code
	file, err := parser.ParseFile(fset, "test.go", testCode, parser.AllErrors)
	if err != nil {
		return fmt.Errorf("parse error: %w", err)
	}

	// Type-check the parsed code
	conf := types.Config{
		Importer: importer.Default(),
	}

	_, err = conf.Check(pkgName, fset, []*ast.File{file}, nil)
	if err != nil {
		return fmt.Errorf("type check error: %w", err)
	}

	return nil
}
