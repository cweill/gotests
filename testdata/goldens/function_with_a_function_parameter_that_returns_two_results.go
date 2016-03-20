package testdata

import "testing"

func TestFoo15(t *testing.T) {
	tests := []struct {
		// Test description.
		name string
		// Parameters.
		f func(string) (string, error)
		// Expected results.
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if err := Foo15(tt.f); (err != nil) != tt.wantErr {
			t.Errorf("%q. Foo15() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}
