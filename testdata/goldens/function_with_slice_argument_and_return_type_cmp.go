package testdata

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFoo11(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name    string
		args    args
		want    []*Bar
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, err := Foo11(tt.args.strs)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. Foo11() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !cmp.Equal(tt.want, got) {
			t.Errorf("%q. Foo11() = %v, want %v\ndiff=%s", tt.name, got, tt.want, cmp.Diff(tt.want, got))
		}
	}
}
