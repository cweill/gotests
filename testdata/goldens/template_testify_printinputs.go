package testdata

import (
	"fmt"
	"testing"
)

func TestFoo201a(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		assert.Equalf(t, tt.want, Foo201a(), "%q. Foo201a()", tt.name)
	}
}

func TestFoo201b(t *testing.T) {
	tests := []struct {
		name      string
		assertion assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for range tests {
		tt.assertion(t, Foo201b(), fmt.Sprintf("%q. Foo201b()", tt.name))
	}
}

func TestFoo201c(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name      string
		args      args
		assertion assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		tt.assertion(t, Foo201c(tt.args.n), fmt.Sprintf("%q. Foo201c(%v)", tt.name, tt.args.n))
	}
}

func TestFoo201d(t *testing.T) {
	tests := []struct {
		name      string
		want      bool
		assertion assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, err := Foo201d()
		tt.assertion(t, err, fmt.Sprintf("%q. Foo201d()", tt.name))
		assert.Equalf(t, tt.want, got, "%q. Foo201d()", tt.name)
	}
}

func TestFoo201e(t *testing.T) {
	type args struct {
		n int
		s string
	}
	tests := []struct {
		name      string
		args      args
		want      bool
		assertion assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, err := Foo201e(tt.args.n, tt.args.s)
		tt.assertion(t, err, fmt.Sprintf("%q. Foo201e(%v, %v)", tt.name, tt.args.n, tt.args.s))
		assert.Equalf(t, tt.want, got, "%q. Foo201e(%v, %v)", tt.name, tt.args.n, tt.args.s)
	}
}
