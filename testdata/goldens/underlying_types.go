package testdata

import (
	"testing"
	"time"
)

func TestCelsius_ToFahrenheit(t *testing.T) {
	tests := []struct {
		name string
		c    Celsius
		want Fahrenheit
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := tt.c.ToFahrenheit(); got != tt.want {
			t.Errorf("%q. Celsius.ToFahrenheit() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestHourToSecond(t *testing.T) {
	tests := []struct {
		name string
		arg  time.Duration
		want time.Duration
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := HourToSecond(tt.arg); got != tt.want {
			t.Errorf("%q. HourToSecond() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
