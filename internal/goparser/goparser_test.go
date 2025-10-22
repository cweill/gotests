package goparser

import (
	"go/importer"
	"os"
	"path/filepath"
	"testing"

	"github.com/cweill/gotests/internal/models"
)

func TestParser_Parse(t *testing.T) {
	tests := []struct {
		name       string
		setupFn    func(t *testing.T) (string, []models.Path)
		wantErr    bool
		wantFuncs  int
		wantPkg    string
	}{
		{
			name: "simple function",
			setupFn: func(t *testing.T) (string, []models.Path) {
				tmpDir := t.TempDir()
				srcPath := filepath.Join(tmpDir, "test.go")
				content := `package mypackage

func Add(a, b int) int {
	return a + b
}
`
				if err := os.WriteFile(srcPath, []byte(content), 0644); err != nil {
					t.Fatal(err)
				}
				return srcPath, nil
			},
			wantErr:   false,
			wantFuncs: 1,
			wantPkg:   "mypackage",
		},
		{
			name: "multiple functions",
			setupFn: func(t *testing.T) (string, []models.Path) {
				tmpDir := t.TempDir()
				srcPath := filepath.Join(tmpDir, "test.go")
				content := `package calc

func Add(a, b int) int {
	return a + b
}

func Subtract(a, b int) int {
	return a - b
}

func Multiply(a, b int) int {
	return a * b
}
`
				if err := os.WriteFile(srcPath, []byte(content), 0644); err != nil {
					t.Fatal(err)
				}
				return srcPath, nil
			},
			wantErr:   false,
			wantFuncs: 3,
			wantPkg:   "calc",
		},
		{
			name: "function with method receiver",
			setupFn: func(t *testing.T) (string, []models.Path) {
				tmpDir := t.TempDir()
				srcPath := filepath.Join(tmpDir, "test.go")
				content := `package service

type Handler struct {
	name string
}

func (h *Handler) Process() error {
	return nil
}

func NewHandler() *Handler {
	return &Handler{}
}
`
				if err := os.WriteFile(srcPath, []byte(content), 0644); err != nil {
					t.Fatal(err)
				}
				return srcPath, nil
			},
			wantErr:   false,
			wantFuncs: 2,
			wantPkg:   "service",
		},
		{
			name: "empty file returns error",
			setupFn: func(t *testing.T) (string, []models.Path) {
				tmpDir := t.TempDir()
				srcPath := filepath.Join(tmpDir, "empty.go")
				if err := os.WriteFile(srcPath, []byte(""), 0644); err != nil {
					t.Fatal(err)
				}
				return srcPath, nil
			},
			wantErr: true,
		},
		{
			name: "nonexistent file returns error",
			setupFn: func(t *testing.T) (string, []models.Path) {
				tmpDir := t.TempDir()
				return filepath.Join(tmpDir, "nonexistent.go"), nil
			},
			wantErr: true,
		},
		{
			name: "generic function",
			setupFn: func(t *testing.T) (string, []models.Path) {
				tmpDir := t.TempDir()
				srcPath := filepath.Join(tmpDir, "test.go")
				content := `package generic

func Transform[T, U any](input []T, fn func(T) U) []U {
	result := make([]U, len(input))
	for i, v := range input {
		result[i] = fn(v)
	}
	return result
}
`
				if err := os.WriteFile(srcPath, []byte(content), 0644); err != nil {
					t.Fatal(err)
				}
				return srcPath, nil
			},
			wantErr:   false,
			wantFuncs: 1,
			wantPkg:   "generic",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srcPath, files := tt.setupFn(t)
			p := &Parser{
				Importer: importer.Default(),
			}

			result, err := p.Parse(srcPath, files)

			if (err != nil) != tt.wantErr {
				t.Errorf("Parser.Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr {
				return
			}

			if result == nil {
				t.Fatal("Parser.Parse() returned nil result")
			}

			if result.Header == nil {
				t.Fatal("Parser.Parse() returned nil header")
			}

			if result.Header.Package != tt.wantPkg {
				t.Errorf("Parser.Parse() package = %v, want %v", result.Header.Package, tt.wantPkg)
			}

			if len(result.Funcs) != tt.wantFuncs {
				t.Errorf("Parser.Parse() returned %d functions, want %d", len(result.Funcs), tt.wantFuncs)
			}
		})
	}
}

func TestParser_Parse_WithImports(t *testing.T) {
	tmpDir := t.TempDir()
	srcPath := filepath.Join(tmpDir, "test.go")
	content := `package mypackage

import (
	"fmt"
	"strings"
)

func Greet(name string) string {
	return fmt.Sprintf("Hello, %s", strings.ToUpper(name))
}
`
	if err := os.WriteFile(srcPath, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	p := &Parser{
		Importer: importer.Default(),
	}

	result, err := p.Parse(srcPath, nil)
	if err != nil {
		t.Fatalf("Parser.Parse() error = %v", err)
	}

	if len(result.Header.Imports) < 2 {
		t.Errorf("Parser.Parse() returned %d imports, want at least 2", len(result.Header.Imports))
	}

	// Check that imports were parsed
	hasImport := func(path string) bool {
		for _, imp := range result.Header.Imports {
			if imp.Path == `"`+path+`"` {
				return true
			}
		}
		return false
	}

	if !hasImport("fmt") {
		t.Error("Parser.Parse() did not find fmt import")
	}
	if !hasImport("strings") {
		t.Error("Parser.Parse() did not find strings import")
	}
}

func TestErrEmptyFile(t *testing.T) {
	if ErrEmptyFile == nil {
		t.Error("ErrEmptyFile should not be nil")
	}

	if ErrEmptyFile.Error() != "file is empty" {
		t.Errorf("ErrEmptyFile.Error() = %q, want %q", ErrEmptyFile.Error(), "file is empty")
	}
}
