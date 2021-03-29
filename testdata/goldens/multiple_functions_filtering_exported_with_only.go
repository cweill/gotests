package testdata

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFooFilter(t *testing.T) {
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
		got, err := FooFilter(tt.args.strs)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. FooFilter() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !cmp.Equal(got, tt.want) {
			t.Errorf("%q. FooFilter() = %v, want %v\ndiff=%s", tt.name, got, tt.want, cmp.Diff(got, tt.want))
		}
	}
}
