package testdata

import "testing"

func TestFoo19(t *testing.T) {
	tests := []struct {
		// Test description.
		name string
		// Parameters.
		in1 string
		in2 string
		in3 string
		// Expected results.
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := Foo19(tt.in1, tt.in2, tt.in3); got != tt.want {
			t.Errorf("%q. Foo19() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
