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
		want    []*testdata.Bar
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := testdata.FooFilter(tt.args.strs)
			if (err != nil) != tt.wantErr {
				t.Fatalf("FooFilter() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FooFilter() = %v, want %v", got, tt.want)
			}
		})
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
		t.Run(tt.name, func(t *testing.T) {
			b := &testdata.Bar{}
			if err := b.BarFilter(tt.args.i); (err != nil) != tt.wantErr {
				t.Errorf("Bar.BarFilter() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
