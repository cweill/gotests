package blanktest

import (
	"os"
	"testing"
)

func TestNot(t *testing.T) {
	tests := []struct {
		name string
		arg  *os.File
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := Not(tt.arg); got != tt.want {
			t.Errorf("%q. Not() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
