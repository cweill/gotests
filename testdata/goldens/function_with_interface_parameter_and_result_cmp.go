package testdata

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFoo21(t *testing.T) {
	type args struct {
		i interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := Foo21(tt.args.i); !cmp.Equal(tt.want, got) {
			t.Errorf("%q. Foo21() = %v, want %v\ndiff=%s", tt.name, got, tt.want, cmp.Diff(tt.want, got))
		}
	}
}
