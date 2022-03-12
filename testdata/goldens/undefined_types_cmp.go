package undefinedtypes

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestUndefined_Do(t *testing.T) {
	type args struct {
		es Something
	}
	tests := []struct {
		name    string
		u       *Undefined
		args    args
		want    *Unknown
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, err := tt.u.Do(tt.args.es)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. Undefined.Do() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !cmp.Equal(tt.want, got) {
			t.Errorf("%q. Undefined.Do() = %v, want %v\ndiff=%s", tt.name, got, tt.want, cmp.Diff(tt.want, got))
		}
	}
}
