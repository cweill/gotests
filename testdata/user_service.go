package testdata

import (
	"errors"
	"regexp"
	"strings"
)

// User represents a user in the system
type User struct {
	ID       int
	Email    string
	Name     string
	Password string
}

// ValidateEmail checks if an email address is valid
func ValidateEmail(email string) error {
	if email == "" {
		return errors.New("email cannot be empty")
	}

	// Simple email regex pattern
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(email) {
		return errors.New("invalid email format")
	}

	return nil
}

// HashPassword simulates password hashing with validation
func HashPassword(password string) (string, error) {
	if len(password) < 8 {
		return "", errors.New("password must be at least 8 characters")
	}

	if len(password) > 72 {
		return "", errors.New("password too long (max 72 characters)")
	}

	// Check for at least one letter and one number
	hasLetter := false
	hasNumber := false
	for _, ch := range password {
		if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') {
			hasLetter = true
		}
		if ch >= '0' && ch <= '9' {
			hasNumber = true
		}
	}

	if !hasLetter || !hasNumber {
		return "", errors.New("password must contain both letters and numbers")
	}

	// Simulate hashing by prefixing (not real crypto!)
	return "hashed_" + password, nil
}

// FindUserByID searches for a user in a slice by ID
func FindUserByID(users []*User, id int) (*User, error) {
	if users == nil {
		return nil, errors.New("users list cannot be nil")
	}

	if id <= 0 {
		return nil, errors.New("invalid user ID")
	}

	for _, user := range users {
		if user != nil && user.ID == id {
			return user, nil
		}
	}

	return nil, errors.New("user not found")
}

// SanitizeUsername cleans and validates a username
func SanitizeUsername(username string) (string, error) {
	// Trim whitespace
	username = strings.TrimSpace(username)

	if username == "" {
		return "", errors.New("username cannot be empty")
	}

	if len(username) < 3 {
		return "", errors.New("username too short (min 3 characters)")
	}

	if len(username) > 20 {
		return "", errors.New("username too long (max 20 characters)")
	}

	// Convert to lowercase and remove special characters
	username = strings.ToLower(username)
	sanitized := strings.Builder{}
	for _, ch := range username {
		if (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9') || ch == '_' {
			sanitized.WriteRune(ch)
		}
	}

	result := sanitized.String()
	if result == "" {
		return "", errors.New("username contains no valid characters")
	}

	return result, nil
}
