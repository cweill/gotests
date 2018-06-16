package testdata

import (
	"reflect"
	"testing"
)

func TestFoo10(t *testing.T) {
	type args struct {
		m map[string]int32
	}
	tests := []struct {
		name string
		args args
		want map[string]*Bar
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := Foo10(tt.args.m); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. Foo10() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func BenchmarkFoo10(b *testing.B) {
	type args struct {
		m map[string]int32
	}
	benchmarks := []struct {
		name string
		args args
		want map[string]*Bar
	}{
	// TODO: Add benchmark cases.
	}
	for _, bb := range benchmarks {
		if got := Foo10(tt.args.m); !reflect.DeepEqual(got, bb.want) {
			b.Errorf("%q. Foo10() = %v, want %v", tt.name, got, bb.want)
		}
	}
}
