package testdata

import "testing"

func TestFoo19(t *testing.T) {
	type args struct {
		in1 string
		in2 string
		in3 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := Foo19(tt.args.in1, tt.args.in2, tt.args.in3); got != tt.want {
			t.Errorf("%q. Foo19() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func BenchmarkFoo19(b *testing.B) {
	type args struct {
		in1 string
		in2 string
		in3 string
	}
	benchmarks := []struct {
		name string
		args args
		want string
	}{
	// TODO: Add benchmark cases.
	}
	for _, bb := range benchmarks {
		if got := Foo19(tt.args.in1, tt.args.in2, tt.args.in3); got != bb.want {
			b.Errorf("%q. Foo19() = %v, want %v", tt.name, got, bb.want)
		}
	}
}
