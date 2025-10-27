package gomod

import (
	"path/filepath"
	"testing"
)

func TestGetFullImportPath(t *testing.T) {
	tests := []struct {
		name     string
		startAt  string
		expected string
		hasError bool
	}{
		{
			name:     "file at module root",
			startAt:  "testdata/simple_module/main.go",
			expected: "example.com/project",
		},
		{
			name:     "directory at module root",
			startAt:  "testdata/simple_module",
			expected: "example.com/project",
		},
		{
			name:     "file in subdirectory",
			startAt:  "testdata/simple_module/pkg/utils.go",
			expected: "example.com/project/pkg",
		},
		{
			name:     "directory in subdirectory",
			startAt:  "testdata/simple_module/pkg",
			expected: "example.com/project/pkg",
		},
		{
			name:     "nested module file",
			startAt:  "testdata/nested_module/lib.go",
			expected: "example.com/nested",
		},
		{
			name:     "nested module directory",
			startAt:  "testdata/nested_module",
			expected: "example.com/nested",
		},
		{
			name:     "path within module",
			startAt:  "testdata/no_module/orphan.go",
			expected: "github.com/cweill/gotests/internal/gomod/testdata/no_module",
		},
		{
			name:     "nonexistent path",
			startAt:  "/tmp/nonexistent/path/file.go",
			hasError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := GetFullImportPath(tt.startAt)
			if tt.hasError {
				if err == nil {
					t.Errorf("expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if result != tt.expected {
					t.Errorf("expected %q, got %q", tt.expected, result)
				}
			}
		})
	}
}

func TestFindGoMod(t *testing.T) {
	tests := []struct {
		name     string
		startDir string
		expected string
		hasError bool
	}{
		{
			name:     "go.mod in current directory",
			startDir: "testdata/simple_module",
			expected: "testdata/simple_module",
		},
		{
			name:     "go.mod in subdirectory",
			startDir: "testdata/simple_module/pkg",
			expected: "testdata/simple_module",
		},
		{
			name:     "no go.mod found",
			startDir: "/tmp",
			hasError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := findGoMod(tt.startDir)
			if tt.hasError {
				if err == nil {
					t.Errorf("expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				// Convert to absolute path for comparison
				expectedAbs, _ := filepath.Abs(tt.expected)
				resultAbs, _ := filepath.Abs(result)
				if resultAbs != expectedAbs {
					t.Errorf("expected %q, got %q", expectedAbs, resultAbs)
				}
			}
		})
	}
}

func TestParseModulePath(t *testing.T) {
	tests := []struct {
		name     string
		modRoot  string
		expected string
		hasError bool
	}{
		{
			name:     "valid go.mod",
			modRoot:  "testdata/simple_module",
			expected: "example.com/project",
		},
		{
			name:     "nested module go.mod",
			modRoot:  "testdata/nested_module",
			expected: "example.com/nested",
		},
		{
			name:     "nonexistent go.mod",
			modRoot:  "/tmp/nonexistent",
			hasError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := parseModulePath(tt.modRoot)
			if tt.hasError {
				if err == nil {
					t.Errorf("expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if result != tt.expected {
					t.Errorf("expected %q, got %q", tt.expected, result)
				}
			}
		})
	}
}

// Test with the actual gotests module
func TestGetFullImportPath_CurrentModule(t *testing.T) {
	// Test with the gotests module itself
	result, err := GetFullImportPath("gomod.go")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := "github.com/cweill/gotests/internal/gomod"
	if result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}
}

func TestGetFullImportPath_EdgeCases(t *testing.T) {
	// Test with current directory (when running from internal/gomod)
	result, err := GetFullImportPath(".")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := "github.com/cweill/gotests/internal/gomod"
	if result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}

	// Test with absolute path
	absPath, _ := filepath.Abs("gomod.go")
	result, err = GetFullImportPath(absPath)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected = "github.com/cweill/gotests/internal/gomod"
	if result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}
}
