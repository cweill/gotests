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
		if got := b.Foo9(); !cmp.Equal(got, tt.want) {
			t.Errorf("%q. Bar.Foo9() = %v, want %v\ndiff=%v", tt.name, got, tt.want, cmp.Diff(got, tt.want))
		}
	}
}
