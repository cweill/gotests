package testdata

import (
	"os"
	"reflect"
	"testing"
)

func TestFoo18(t *testing.T) {
	tests := []struct {
		name string
		arg  *os.File
		want *os.File
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := Foo18(tt.arg); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. Foo18() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
