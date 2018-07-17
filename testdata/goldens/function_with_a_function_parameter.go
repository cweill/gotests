package testdata

import "testing"

func TestFoo13(t *testing.T) {
	tests := []struct {
		name    string
		arg     func()
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if err := Foo13(tt.arg); (err != nil) != tt.wantErr {
			t.Errorf("%q. Foo13() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}
