package testdata

import (
	"errors"
	"strings"
)

// TrimAndLower trims whitespace and converts to lowercase
func TrimAndLower(s string) string {
	return strings.ToLower(strings.TrimSpace(s))
}

// Join joins strings with a separator, handling empty inputs
func Join(parts ...string) string {
	if len(parts) == 0 {
		return ""
	}

	// Filter out empty strings
	nonEmpty := make([]string, 0, len(parts))
	for _, part := range parts {
		if part != "" {
			nonEmpty = append(nonEmpty, part)
		}
	}

	if len(nonEmpty) == 0 {
		return ""
	}

	return strings.Join(nonEmpty, " ")
}

// ParseKeyValue parses "key=value" pairs from input
func ParseKeyValue(input string) (map[string]string, error) {
	if input == "" {
		return nil, errors.New("input cannot be empty")
	}

	result := make(map[string]string)
	pairs := strings.Split(input, ",")

	for _, pair := range pairs {
		pair = strings.TrimSpace(pair)
		if pair == "" {
			continue
		}

		parts := strings.SplitN(pair, "=", 2)
		if len(parts) != 2 {
			return nil, errors.New("invalid format: expected key=value")
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		if key == "" {
			return nil, errors.New("key cannot be empty")
		}

		result[key] = value
	}

	if len(result) == 0 {
		return nil, errors.New("no valid key-value pairs found")
	}

	return result, nil
}

// Reverse reverses a string
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// ContainsAny checks if a string contains any of the substrings
func ContainsAny(s string, substrings []string) bool {
	if s == "" || len(substrings) == 0 {
		return false
	}

	for _, substr := range substrings {
		if substr != "" && strings.Contains(s, substr) {
			return true
		}
	}

	return false
}

// TruncateWithEllipsis truncates a string to maxLen with ellipsis
func TruncateWithEllipsis(s string, maxLen int) string {
	if maxLen <= 0 {
		return ""
	}

	if maxLen <= 3 {
		// Too short for ellipsis
		if len(s) <= maxLen {
			return s
		}
		return s[:maxLen]
	}

	if len(s) <= maxLen {
		return s
	}

	return s[:maxLen-3] + "..."
}
