package testdata

import (
	ht "html/template"
	"reflect"
	"testing"
)

func TestFoo22(t *testing.T) {
	tests := []struct {
		name string
		t    *ht.Template
		want *ht.Template
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := Foo22(tt.t); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. Foo22() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
