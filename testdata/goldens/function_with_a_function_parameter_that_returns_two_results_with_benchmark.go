package testdata

import "testing"

func TestFoo15(t *testing.T) {
	type args struct {
		f func(string) (string, error)
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if err := Foo15(tt.args.f); (err != nil) != tt.wantErr {
			t.Errorf("%q. Foo15() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}

func BenchmarkFoo15(b *testing.B) {
	type args struct {
		f func(string) (string, error)
	}
	benchmarks := []struct {
		name    string
		args    args
		wantErr bool
	}{
	// TODO: Add benchmark cases.
	}
	for _, bb := range benchmarks {
		if err := Foo15(tt.args.f); (err != nil) != bb.wantErr {
			b.Errorf("%q. Foo15() error = %v, wantErr %v", tt.name, err, bb.wantErr)
		}
	}
}
