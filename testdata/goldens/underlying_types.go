package testdata

import (
	"testing"
	"time"
)

func TestCelsiusToFahrenheit(t *testing.T) {
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
		h    time.Duration
		want time.Duration
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := HourToSecond(tt.h); got != tt.want {
			t.Errorf("%q. HourToSecond() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
