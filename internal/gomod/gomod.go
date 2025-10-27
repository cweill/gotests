package gomod

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// GetFullImportPath resolves the full Go import path for any file or directory
// within a Go module. Returns the complete import path like "github.com/user/repo/pkg".
//
// startAt can be either:
//   - A Go source file path (e.g., "/path/to/project/main.go")
//   - A directory path (e.g., "/path/to/project/pkg")
//   - An absolute or relative path
//
// Returns an error if:
//   - No go.mod found in the directory tree
//   - go.mod is malformed or missing module directive
//   - Path resolution fails
func GetFullImportPath(startAt string) (string, error) {
	absPath, err := filepath.Abs(startAt)
	if err != nil {
		return "", fmt.Errorf("failed to get absolute path for %s: %w", startAt, err)
	}

	// If it's a file, get its directory
	if info, err := os.Stat(absPath); err == nil && !info.IsDir() {
		absPath = filepath.Dir(absPath)
	}

	modRoot, err := findGoMod(absPath)
	if err != nil {
		return "", err
	}

	modulePath, err := parseModulePath(modRoot)
	if err != nil {
		return "", err
	}

	relPath, err := filepath.Rel(modRoot, absPath)
	if err != nil {
		return "", fmt.Errorf("failed to calculate relative path from %s to %s: %w", modRoot, absPath, err)
	}

	if relPath == "." {
		return modulePath, nil
	}

	return filepath.Join(modulePath, relPath), nil
}

// findGoMod walks up the directory tree from startDir to find a go.mod file.
// Returns the directory containing go.mod, or an error if not found.
func findGoMod(startDir string) (string, error) {
	current := startDir

	for {
		modPath := filepath.Join(current, "go.mod")
		if _, err := os.Stat(modPath); err == nil {
			return current, nil
		}

		parent := filepath.Dir(current)
		if parent == current {
			// Reached root directory
			break
		}
		current = parent
	}

	return "", fmt.Errorf("go.mod not found in %s or any parent directory", startDir)
}

// parseModulePath reads the go.mod file in modRoot and extracts the module path.
// Returns the module path or an error if parsing fails.
func parseModulePath(modRoot string) (string, error) {
	modFile := filepath.Join(modRoot, "go.mod")

	file, err := os.Open(modFile)
	if err != nil {
		return "", fmt.Errorf("failed to open go.mod at %s: %w", modFile, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "module ") {
			modulePath := strings.TrimSpace(line[7:]) // Remove "module " prefix
			if modulePath == "" {
				return "", fmt.Errorf("empty module path in %s", modFile)
			}
			return modulePath, nil
		}
	}

	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("error reading go.mod at %s: %w", modFile, err)
	}

	return "", fmt.Errorf("module directive not found in %s", modFile)
}
