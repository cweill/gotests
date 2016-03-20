package testdata

import "testing"

func TestBarBarFilter(t *testing.T) {
	tests := []struct {
		// Test description.
		name string
		// Parameters.
		i interface{}
		// Expected results.
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		b := &Bar{}
		if err := b.BarFilter(tt.i); (err != nil) != tt.wantErr {
			t.Errorf("%q. Bar.BarFilter() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}
