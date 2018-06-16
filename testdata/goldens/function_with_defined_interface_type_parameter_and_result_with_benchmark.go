package testdata

import (
	"reflect"
	"testing"
)

func TestFoo16(t *testing.T) {
	type args struct {
		in Bazzar
	}
	tests := []struct {
		name string
		args args
		want Bazzar
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := Foo16(tt.args.in); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. Foo16() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func BenchmarkFoo16(b *testing.B) {
	type args struct {
		in Bazzar
	}
	benchmarks := []struct {
		name string
		args args
		want Bazzar
	}{
	// TODO: Add benchmark cases.
	}
	for _, bb := range benchmarks {
		if got := Foo16(tt.args.in); !reflect.DeepEqual(got, bb.want) {
			b.Errorf("%q. Foo16() = %v, want %v", tt.name, got, bb.want)
		}
	}
}
