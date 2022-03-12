package testdata

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFoo16(t *testing.T) {
	type args struct {
		in Bazzar
	}
	tests := []struct {
		name string
		args args
		want Bazzar
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := Foo16(tt.args.in); !cmp.Equal(tt.want, got) {
			t.Errorf("%q. Foo16() = %v, want %v\ndiff=%s", tt.name, got, tt.want, cmp.Diff(tt.want, got))
		}
	}
}
