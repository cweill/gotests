package testdata

import (
	"reflect"
	"testing"
)

func TestFoo26(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantI   int
		want2   []byte
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, gotI, got2, err := Foo26(tt.args.v)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. Foo26() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("%q. Foo26() got = %v, want %v", tt.name, got, tt.want)
		}
		if gotI != tt.wantI {
			t.Errorf("%q. Foo26() gotI = %v, want %v", tt.name, gotI, tt.wantI)
		}
		if !reflect.DeepEqual(got2, tt.want2) {
			t.Errorf("%q. Foo26() got2 = %v, want %v", tt.name, got2, tt.want2)
		}
	}
}

func BenchmarkFoo26(b *testing.B) {
	type args struct {
		v interface{}
	}
	benchmarks := []struct {
		name    string
		args    args
		want    string
		wantI   int
		want2   []byte
		wantErr bool
	}{
	// TODO: Add benchmark cases.
	}
	for _, bb := range benchmarks {
		got, gotI, got2, err := Foo26(tt.args.v)
		if (err != nil) != bb.wantErr {
			b.Errorf("%q. Foo26() error = %v, wantErr %v", tt.name, err, bb.wantErr)
			continue
		}
		if got != bb.want {
			b.Errorf("%q. Foo26() got = %v, want %v", tt.name, got, bb.want)
		}
		if gotI != bb.wantI {
			b.Errorf("%q. Foo26() gotI = %v, want %v", tt.name, gotI, bb.wantI)
		}
		if !reflect.DeepEqual(got2, bb.want2) {
			b.Errorf("%q. Foo26() got2 = %v, want %v", tt.name, got2, bb.want2)
		}
	}
}
