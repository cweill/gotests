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

func TestBar_BarFilter(t *testing.T) {
	tests := []struct {
		name    string
		b       *Bar
		arg     interface{}
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		b := &Bar{}
		if err := b.BarFilter(tt.arg); (err != nil) != tt.wantErr {
			t.Errorf("%q. Bar.BarFilter() error = %v, wantErr %v", tt.name, err, tt.wantErr)
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
