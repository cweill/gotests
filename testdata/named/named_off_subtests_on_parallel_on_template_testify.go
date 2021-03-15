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
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.want, Foo038())
		})
	}
}
