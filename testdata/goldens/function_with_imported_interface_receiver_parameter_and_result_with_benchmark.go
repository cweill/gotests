package testdata

import (
	"io"
	"reflect"
	"testing"
)

func TestFoo17(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name string
		args args
		want io.Reader
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := Foo17(tt.args.r); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. Foo17() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func BenchmarkFoo17(b *testing.B) {
	type args struct {
		r io.Reader
	}
	benchmarks := []struct {
		name string
		args args
		want io.Reader
	}{
	// TODO: Add benchmark cases.
	}
	for _, bb := range benchmarks {
		if got := Foo17(tt.args.r); !reflect.DeepEqual(got, bb.want) {
			b.Errorf("%q. Foo17() = %v, want %v", tt.name, got, bb.want)
		}
	}
}
