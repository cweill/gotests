package testdata

import (
	ht "html/template"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFoo22(t *testing.T) {
	type args struct {
		t *ht.Template
	}
	tests := []struct {
		name string
		args args
		want *ht.Template
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := Foo22(tt.args.t); !cmp.Equal(got, tt.want) {
			t.Errorf("%q. Foo22() = %v, want %v\ndiff=%v", tt.name, got, tt.want, cmp.Diff(got, tt.want))
		}
	}
}
