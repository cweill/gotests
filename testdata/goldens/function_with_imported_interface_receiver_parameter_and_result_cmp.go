package testdata

import (
	"io"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFoo17(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name string
		args args
		want io.Reader
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := Foo17(tt.args.r); !cmp.Equal(tt.want, got) {
			t.Errorf("%q. Foo17() = %v, want %v\ndiff=%s", tt.name, got, tt.want, cmp.Diff(tt.want, got))
		}
	}
}
