package testdata

import "testing"

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
