package testdata

import (
	"fmt"
	"testing"
)

var example102 = fmt.Sprintf("test%v", 1)

func TestFoo102(t *testing.T) {
	tests := []struct {
		// Test description.
		name string
		// Parameters.
		s string
		// Expected results.
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := Foo102(tt.s); got != tt.want {
			t.Errorf("%q. Foo102() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
