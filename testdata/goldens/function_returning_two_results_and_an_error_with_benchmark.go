package testdata

import (
	"reflect"
	"testing"
)

func TestFoo25(t *testing.T) {
	type args struct {
		in0 interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    string
		want1   []byte
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, got1, err := Foo25(tt.args.in0)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. Foo25() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("%q. Foo25() got = %v, want %v", tt.name, got, tt.want)
		}
		if !reflect.DeepEqual(got1, tt.want1) {
			t.Errorf("%q. Foo25() got1 = %v, want %v", tt.name, got1, tt.want1)
		}
	}
}

func BenchmarkFoo25(b *testing.B) {
	type args struct {
		in0 interface{}
	}
	benchmarks := []struct {
		name    string
		args    args
		want    string
		want1   []byte
		wantErr bool
	}{
	// TODO: Add benchmark cases.
	}
	for _, bb := range benchmarks {
		got, got1, err := Foo25(tt.args.in0)
		if (err != nil) != bb.wantErr {
			b.Errorf("%q. Foo25() error = %v, wantErr %v", tt.name, err, bb.wantErr)
			continue
		}
		if got != bb.want {
			b.Errorf("%q. Foo25() got = %v, want %v", tt.name, got, bb.want)
		}
		if !reflect.DeepEqual(got1, bb.want1) {
			b.Errorf("%q. Foo25() got1 = %v, want %v", tt.name, got1, bb.want1)
		}
	}
}
