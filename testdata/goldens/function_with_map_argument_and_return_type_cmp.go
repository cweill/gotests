package testdata

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFoo10(t *testing.T) {
	type args struct {
		m map[string]int32
	}
	tests := []struct {
		name string
		args args
		want map[string]*Bar
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := Foo10(tt.args.m); !cmp.Equal(tt.want, got) {
			t.Errorf("%q. Foo10() = %v, want %v\ndiff=%s", tt.name, got, tt.want, cmp.Diff(tt.want, got))
		}
	}
}
