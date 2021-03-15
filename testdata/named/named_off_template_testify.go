package testdata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFoo038(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		assert.Equalf(t, tt.want, Foo038(), "%q. Foo038()", tt.name)
	}
}
