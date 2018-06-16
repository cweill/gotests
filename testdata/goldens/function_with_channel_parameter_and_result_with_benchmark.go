package testdata

import (
	"reflect"
	"testing"
)

func TestFoo23(t *testing.T) {
	type args struct {
		ch chan bool
	}
	tests := []struct {
		name string
		args args
		want chan string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := Foo23(tt.args.ch); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. Foo23() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func BenchmarkFoo23(b *testing.B) {
	type args struct {
		ch chan bool
	}
	benchmarks := []struct {
		name string
		args args
		want chan string
	}{
	// TODO: Add benchmark cases.
	}
	for _, bb := range benchmarks {
		if got := Foo23(tt.args.ch); !reflect.DeepEqual(got, bb.want) {
			b.Errorf("%q. Foo23() = %v, want %v", tt.name, got, bb.want)
		}
	}
}
