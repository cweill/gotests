package testdata

import "testing"

func TestFoo038(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := Foo038(); got != tt.want {
			t.Errorf("%q. Foo038() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
