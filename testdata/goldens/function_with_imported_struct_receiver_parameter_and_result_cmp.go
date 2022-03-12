package testdata

import (
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFoo18(t *testing.T) {
	type args struct {
		t *os.File
	}
	tests := []struct {
		name string
		args args
		want *os.File
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := Foo18(tt.args.t); !cmp.Equal(tt.want, got) {
			t.Errorf("%q. Foo18() = %v, want %v\ndiff=%s", tt.name, got, tt.want, cmp.Diff(tt.want, got))
		}
	}
}
