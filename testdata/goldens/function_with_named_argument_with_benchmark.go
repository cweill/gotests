package testdata

import "testing"

func TestFoo3(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		Foo3(tt.args.s)
	}
}

func BenchmarkFoo3(b *testing.B) {
	type args struct {
		s string
	}
	benchmarks := []struct {
		name string
		args args
	}{
	// TODO: Add benchmark cases.
	}
	for _, bb := range benchmarks {
		Foo3(tt.args.s)
	}
}
