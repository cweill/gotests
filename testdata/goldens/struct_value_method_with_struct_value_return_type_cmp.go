package testdata

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestBar_Foo9(t *testing.T) {
	tests := []struct {
		name string
		b    Bar
		want Bar
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		b := Bar{}
		if got := b.Foo9(); !cmp.Equal(tt.want, got) {
			t.Errorf("%q. Bar.Foo9() = %v, want %v\ndiff=%s", tt.name, got, tt.want, cmp.Diff(tt.want, got))
		}
	}
}
