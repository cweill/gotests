package syntaxtest

import (
	"os"
	"testing"
)

// Plural all the types.
func Foo(s strings) errors {
	// Incorrect return type.
	return ""
}

func TestNot(t *testing.T) {
	tests := []struct {
		name string
		this *os.File
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := Not(tt.this); got != tt.want {
			t.Errorf("%q. Not() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
