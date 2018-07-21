package testdata

import (
	"reflect"
	"testing"
)

func TestFoo23(t *testing.T) {
	tests := []struct {
		name string
		arg  chan bool
		want chan string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := Foo23(tt.arg); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. Foo23() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
