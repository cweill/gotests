package testdata

import (
	"reflect"
	"testing"
)

func TestFooFilter(t *testing.T) {
	tests := []struct {
		// Test description.
		name string
		// Parameters.
		strs []string
		// Expected results.
		want    []*Bar
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, err := FooFilter(tt.strs)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. FooFilter() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. FooFilter() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestBazFilter(t *testing.T) {
	tests := []struct {
		// Test description.
		name string
		// Parameters.
		f *float64
		// Expected results.
		want float64
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := bazFilter(tt.f); got != tt.want {
			t.Errorf("%q. bazFilter() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
