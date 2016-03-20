package testdata

import "testing"

func TestFoo20(t *testing.T) {
	tests := []struct {
		// Test description.
		name string
		// Parameters.
		strs []string
		// Expected results.
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := Foo20(tt.strs...); got != tt.want {
			t.Errorf("%q. Foo20() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
