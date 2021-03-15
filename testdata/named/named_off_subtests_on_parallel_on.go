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
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := Foo038(); got != tt.want {
				t.Errorf("Foo038() = %v, want %v", got, tt.want)
			}
		})
	}
}
