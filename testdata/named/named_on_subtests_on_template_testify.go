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
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tt.want, Foo038())
		})
	}
}
