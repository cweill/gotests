package testdata

import "testing"

func TestFoo12(t *testing.T) {
	tests := []struct {
		name    string
		arg     string
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if err := Foo12(tt.arg); (err != nil) != tt.wantErr {
			t.Errorf("%q. Foo12() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}
