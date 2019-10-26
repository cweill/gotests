package testdata

import "testing"

func TestFoo4(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		assert.Equalf(t, tt.want, Foo4(), "%q. Foo4()", tt.name)
	}
}
