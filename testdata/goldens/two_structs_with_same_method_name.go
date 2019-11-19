package testdata

import "testing"

func TestCelsiusString(t *testing.T) {
	tests := []struct {
		name string
		c    Celsius
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := tt.c.String(); got != tt.want {
			t.Errorf("%q. Celsius.String() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestFahrenheitString(t *testing.T) {
	tests := []struct {
		name string
		f    Fahrenheit
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := tt.f.String(); got != tt.want {
			t.Errorf("%q. Fahrenheit.String() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
