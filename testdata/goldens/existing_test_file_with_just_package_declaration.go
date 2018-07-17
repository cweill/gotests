package testdata

import "testing"

func TestFoo101(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := Foo101(tt.arg); got != tt.want {
			t.Errorf("%q. Foo101() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
