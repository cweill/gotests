package testdata

import (
	"reflect"
	"testing"
)

func TestFoo10(t *testing.T) {
	tests := []struct {
		// Test description.
		name string
		// Parameters.
		m map[string]int32
		// Expected results.
		want map[string]*Bar
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := Foo10(tt.m); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. Foo10() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
