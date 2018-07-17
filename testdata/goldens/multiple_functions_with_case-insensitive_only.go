package testdata

import (
	"reflect"
	"testing"
)

func TestFooFilter(t *testing.T) {
	tests := []struct {
		name    string
		arg     []string
		want    []*Bar
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, err := FooFilter(tt.arg)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. FooFilter() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. FooFilter() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func Test_bazFilter(t *testing.T) {
	tests := []struct {
		name string
		arg  *float64
		want float64
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := bazFilter(tt.arg); got != tt.want {
			t.Errorf("%q. bazFilter() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
