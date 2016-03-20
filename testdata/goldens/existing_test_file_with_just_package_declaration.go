package testdata

import "testing"

func TestFoo101(t *testing.T) {
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
		if got := Foo101(tt.s); got != tt.want {
			t.Errorf("%q. Foo101() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
