package testdata_test

import (
	"reflect"
	"testing"

	testdata "github.com/cweill/gotests/testdata"
)

func TestFooFilter(t *testing.T) {
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
		got, err := testdata.FooFilter(tt.args.strs)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. FooFilter() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if tt.wantErr {
			return
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. FooFilter() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestBar_BarFilter(t *testing.T) {
	type args struct {
		i interface{}
	}
	tests := []struct {
		name    string
		b       *testdata.Bar
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		b := &testdata.Bar{}
		if err := b.BarFilter(tt.args.i); (err != nil) != tt.wantErr {
			t.Errorf("%q. Bar.BarFilter() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}
