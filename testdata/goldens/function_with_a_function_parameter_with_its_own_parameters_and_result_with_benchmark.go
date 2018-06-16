package testdata

import "testing"

func TestFoo14(t *testing.T) {
	type args struct {
		f func(string, int) string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if err := Foo14(tt.args.f); (err != nil) != tt.wantErr {
			t.Errorf("%q. Foo14() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}

func BenchmarkFoo14(b *testing.B) {
	type args struct {
		f func(string, int) string
	}
	benchmarks := []struct {
		name    string
		args    args
		wantErr bool
	}{
	// TODO: Add benchmark cases.
	}
	for _, bb := range benchmarks {
		if err := Foo14(tt.args.f); (err != nil) != bb.wantErr {
			b.Errorf("%q. Foo14() error = %v, wantErr %v", tt.name, err, bb.wantErr)
		}
	}
}
