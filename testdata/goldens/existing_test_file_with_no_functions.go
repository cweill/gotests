package testdata

import (
	"fmt"
	"testing"
)

var example102 = fmt.Sprintf("test%v", 1)

func TestFoo102(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := Foo102(tt.arg); got != tt.want {
			t.Errorf("%q. Foo102() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
