package testdata

import "testing"

func TestFoo20(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := Foo20(tt.args.strs...); got != tt.want {
			t.Errorf("%q. Foo20() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func BenchmarkFoo20(b *testing.B) {
	type args struct {
		strs []string
	}
	benchmarks := []struct {
		name string
		args args
		want string
	}{
	// TODO: Add benchmark cases.
	}
	for _, bb := range benchmarks {
		if got := Foo20(tt.args.strs...); got != bb.want {
			b.Errorf("%q. Foo20() = %v, want %v", tt.name, got, bb.want)
		}
	}
}
