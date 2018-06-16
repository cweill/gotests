package testdata

import (
	"os"
	"reflect"
	"testing"
)

func TestFoo18(t *testing.T) {
	type args struct {
		t *os.File
	}
	tests := []struct {
		name string
		args args
		want *os.File
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := Foo18(tt.args.t); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. Foo18() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func BenchmarkFoo18(b *testing.B) {
	type args struct {
		t *os.File
	}
	benchmarks := []struct {
		name string
		args args
		want *os.File
	}{
	// TODO: Add benchmark cases.
	}
	for _, bb := range benchmarks {
		if got := Foo18(tt.args.t); !reflect.DeepEqual(got, bb.want) {
			b.Errorf("%q. Foo18() = %v, want %v", tt.name, got, bb.want)
		}
	}
}
