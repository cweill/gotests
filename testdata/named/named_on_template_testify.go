package testdata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFoo038(t *testing.T) {
	tests := map[string]struct {
		want bool
	}{
		// TODO: Add test cases.
	}
	for name, tt := range tests {
		assert.Equalf(t, tt.want, Foo038(), "%q. Foo038()", name)
	}
}
