package testdata

import "testing"

func TestCelsius_String(t *testing.T) {
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

func BenchmarkCelsius_String(b *testing.B) {
	benchmarks := []struct {
		name string
		c    Celsius
		want string
	}{
	// TODO: Add benchmark cases.
	}
	for _, bb := range benchmarks {
		if got := tt.c.String(); got != bb.want {
			b.Errorf("%q. Celsius.String() = %v, want %v", tt.name, got, bb.want)
		}
	}
}

func TestFahrenheit_String(t *testing.T) {
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

func BenchmarkFahrenheit_String(b *testing.B) {
	benchmarks := []struct {
		name string
		f    Fahrenheit
		want string
	}{
	// TODO: Add benchmark cases.
	}
	for _, bb := range benchmarks {
		if got := tt.f.String(); got != bb.want {
			b.Errorf("%q. Fahrenheit.String() = %v, want %v", tt.name, got, bb.want)
		}
	}
}
