package testdata

import "testing"

func TestFoo13(t *testing.T) {
	type args struct {
		f func()
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if err := Foo13(tt.args.f); (err != nil) != tt.wantErr {
			t.Errorf("%q. Foo13() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}

func BenchmarkFoo13(b *testing.B) {
	type args struct {
		f func()
	}
	benchmarks := []struct {
		name    string
		args    args
		wantErr bool
	}{
	// TODO: Add benchmark cases.
	}
	for _, bb := range benchmarks {
		if err := Foo13(tt.args.f); (err != nil) != bb.wantErr {
			b.Errorf("%q. Foo13() error = %v, wantErr %v", tt.name, err, bb.wantErr)
		}
	}
}
