package testdata

import (
	ht "html/template"
	"reflect"
	"testing"
)

func TestFoo22(t *testing.T) {
	type args struct {
		t *ht.Template
	}
	tests := []struct {
		name string
		args args
		want *ht.Template
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := Foo22(tt.args.t); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. Foo22() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func BenchmarkFoo22(b *testing.B) {
	type args struct {
		t *ht.Template
	}
	benchmarks := []struct {
		name string
		args args
		want *ht.Template
	}{
	// TODO: Add benchmark cases.
	}
	for _, bb := range benchmarks {
		if got := Foo22(tt.args.t); !reflect.DeepEqual(got, bb.want) {
			b.Errorf("%q. Foo22() = %v, want %v", tt.name, got, bb.want)
		}
	}
}
