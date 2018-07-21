package testdata

import "testing"

func TestFoo14(t *testing.T) {
	tests := []struct {
		name    string
		arg     func(string, int) string
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if err := Foo14(tt.arg); (err != nil) != tt.wantErr {
			t.Errorf("%q. Foo14() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}
