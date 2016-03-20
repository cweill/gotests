package testdata

import "testing"

func TestFoo6(t *testing.T) {
	tests := []struct {
		i       int
		b       bool
		want    string
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, err := Foo6(tt.i, tt.b)
		if (err != nil) != tt.wantErr {
			t.Errorf("Foo6(%v, %v) error = %v, wantErr %v", tt.i, tt.b, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("Foo6(%v, %v) = %v, want %v", tt.i, tt.b, got, tt.want)
		}
	}
}
