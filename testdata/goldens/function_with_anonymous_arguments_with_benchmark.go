package testdata

import "testing"

func TestFoo2(t *testing.T) {
	type args struct {
		in0 string
		in1 int
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		Foo2(tt.args.in0, tt.args.in1)
	}
}

func BenchmarkFoo2(b *testing.B) {
	type args struct {
		in0 string
		in1 int
	}
	benchmarks := []struct {
		name string
		args args
	}{
	// TODO: Add benchmark cases.
	}
	for _, bb := range benchmarks {
		Foo2(tt.args.in0, tt.args.in1)
	}
}
