package testdata

import (
	"reflect"
	"testing"
)

func TestFoo16(t *testing.T) {
	tests := []struct {
		name string
		arg  Bazzar
		want Bazzar
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := Foo16(tt.arg); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. Foo16() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
