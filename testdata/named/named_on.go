package testdata

import "testing"

func TestFoo038(t *testing.T) {
	tests := map[string]struct {
		want bool
	}{
		// TODO: Add test cases.
	}
	for name, tt := range tests {
		if got := Foo038(); got != tt.want {
			t.Errorf("%q. Foo038() = %v, want %v", name, got, tt.want)
		}
	}
}
