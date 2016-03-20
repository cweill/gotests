package undefinedtypes

import (
	"reflect"
	"testing"
)

func TestUndefinedDo(t *testing.T) {
	tests := []struct {
		name    string
		u       *Undefined
		es      Something
		want    *Unknown
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, err := tt.u.Do(tt.es)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. Undefined.Do() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. Undefined.Do() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
