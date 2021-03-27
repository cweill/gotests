package testdata

import "testing"

func TestFoo038(t *testing.T) {
	tests := map[string]struct {
		want bool
	}{
		// TODO: Add test cases.
	}
	for name, tt := range tests {
		tt := tt
		name := name
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if got := Foo038(); got != tt.want {
				t.Errorf("Foo038() = %v, want %v", got, tt.want)
			}
		})
	}
}
