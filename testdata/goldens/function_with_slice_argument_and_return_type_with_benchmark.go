package testdata

import (
	"reflect"
	"testing"
)

func TestFoo11(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name    string
		args    args
		want    []*Bar
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, err := Foo11(tt.args.strs)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. Foo11() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. Foo11() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func BenchmarkFoo11(b *testing.B) {
	type args struct {
		strs []string
	}
	benchmarks := []struct {
		name    string
		args    args
		want    []*Bar
		wantErr bool
	}{
	// TODO: Add benchmark cases.
	}
	for _, bb := range benchmarks {
		got, err := Foo11(tt.args.strs)
		if (err != nil) != bb.wantErr {
			b.Errorf("%q. Foo11() error = %v, wantErr %v", tt.name, err, bb.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, bb.want) {
			b.Errorf("%q. Foo11() = %v, want %v", tt.name, got, bb.want)
		}
	}
}
