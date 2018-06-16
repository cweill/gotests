package testdata

import (
	"reflect"
	"testing"
)

func TestFoo21(t *testing.T) {
	type args struct {
		i interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := Foo21(tt.args.i); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. Foo21() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func BenchmarkFoo21(b *testing.B) {
	type args struct {
		i interface{}
	}
	benchmarks := []struct {
		name string
		args args
		want interface{}
	}{
	// TODO: Add benchmark cases.
	}
	for _, bb := range benchmarks {
		if got := Foo21(tt.args.i); !reflect.DeepEqual(got, bb.want) {
			b.Errorf("%q. Foo21() = %v, want %v", tt.name, got, bb.want)
		}
	}
}
