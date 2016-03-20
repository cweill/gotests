package testdata

import (
	"reflect"
	"testing"
)

func TestFoo21(t *testing.T) {
	tests := []struct {
		name string
		i    interface{}
		want interface{}
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := Foo21(tt.i); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. Foo21() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
