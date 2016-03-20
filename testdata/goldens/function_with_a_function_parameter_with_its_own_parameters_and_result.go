package testdata

import "testing"

func TestFoo14(t *testing.T) {
	tests := []struct {
		// Test description.
		name string
		// Parameters.
		f func(string, int) string
		// Expected results.
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if err := Foo14(tt.f); (err != nil) != tt.wantErr {
			t.Errorf("%q. Foo14() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}
