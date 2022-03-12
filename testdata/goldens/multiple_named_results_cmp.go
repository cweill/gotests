package testdata

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFoo26(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantI   int
		want2   []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, gotI, got2, err := Foo26(tt.args.v)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. Foo26() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("%q. Foo26() got = %v, want %v", tt.name, got, tt.want)
		}
		if gotI != tt.wantI {
			t.Errorf("%q. Foo26() gotI = %v, want %v", tt.name, gotI, tt.wantI)
		}
		if !cmp.Equal(tt.want2, got2) {
			t.Errorf("%q. Foo26() got2 = %v, want %v\ndiff=%s", tt.name, got2, tt.want2, cmp.Diff(tt.want2, got2))
		}
	}
}
