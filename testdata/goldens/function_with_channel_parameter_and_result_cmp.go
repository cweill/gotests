package testdata

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFoo23(t *testing.T) {
	type args struct {
		ch chan bool
	}
	tests := []struct {
		name string
		args args
		want chan string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := Foo23(tt.args.ch); !cmp.Equal(tt.want, got) {
			t.Errorf("%q. Foo23() = %v, want %v\ndiff=%s", tt.name, got, tt.want, cmp.Diff(tt.want, got))
		}
	}
}
