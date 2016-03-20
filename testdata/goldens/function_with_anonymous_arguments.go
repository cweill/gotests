package testdata

import "testing"

func TestFoo2(t *testing.T) {
	tests := []struct {
		name string
		in0  string
		in1  int
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		Foo2(tt.in0, tt.in1)
	}
}
