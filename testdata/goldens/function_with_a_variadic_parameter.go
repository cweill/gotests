package testdata

import "testing"

func TestFoo20(t *testing.T) {
	tests := []struct {
		name string
		arg  []string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := Foo20(tt.arg...); got != tt.want {
			t.Errorf("%q. Foo20() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
