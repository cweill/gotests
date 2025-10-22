package testdata

import (
	"path/filepath"
	"strings"
)

// GetExtension extracts the file extension from a filename
func GetExtension(filename string) string {
	if filename == "" {
		return ""
	}

	// Find the last dot
	lastDot := strings.LastIndex(filename, ".")
	if lastDot == -1 || lastDot == 0 {
		// No extension or hidden file with no extension
		return ""
	}

	// Check if the dot is in the filename part (not in directory path)
	lastSlash := strings.LastIndex(filename, "/")
	if lastSlash > lastDot {
		// Dot is in directory name, not filename
		return ""
	}

	return filename[lastDot:]
}

// IsValidPath checks if a path is valid and safe
func IsValidPath(path string) bool {
	if path == "" {
		return false
	}

	// Reject paths with null bytes
	if strings.Contains(path, "\x00") {
		return false
	}

	// Reject path traversal attempts
	if strings.Contains(path, "..") {
		return false
	}

	// Reject absolute paths on Unix
	if strings.HasPrefix(path, "/") {
		return false
	}

	// Reject Windows absolute paths
	if len(path) >= 2 && path[1] == ':' {
		return false
	}

	return true
}

// JoinPaths joins path components and cleans the result
func JoinPaths(base string, parts ...string) string {
	if base == "" && len(parts) == 0 {
		return ""
	}

	// Start with base
	components := []string{base}

	// Add non-empty parts
	for _, part := range parts {
		if part != "" {
			components = append(components, part)
		}
	}

	// Join and clean
	joined := filepath.Join(components...)

	// Normalize separators to forward slash
	return filepath.ToSlash(joined)
}

// GetBaseName returns the base name of a path without extension
func GetBaseName(path string) string {
	if path == "" {
		return ""
	}

	// Get the base filename
	base := filepath.Base(path)

	// Remove extension
	ext := filepath.Ext(base)
	if ext != "" {
		return base[:len(base)-len(ext)]
	}

	return base
}

// IsHiddenFile checks if a filename represents a hidden file
func IsHiddenFile(filename string) bool {
	if filename == "" {
		return false
	}

	// Get just the filename without path
	base := filepath.Base(filename)

	// Unix hidden files start with a dot
	if strings.HasPrefix(base, ".") && base != "." && base != ".." {
		return true
	}

	return false
}
