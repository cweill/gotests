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

func BenchmarkCelsius_ToFahrenheit(b *testing.B) {
	benchmarks := []struct {
		name string
		c    Celsius
		want Fahrenheit
	}{
	// TODO: Add benchmark cases.
	}
	for _, bb := range benchmarks {
		if got := tt.c.ToFahrenheit(); got != bb.want {
			b.Errorf("%q. Celsius.ToFahrenheit() = %v, want %v", tt.name, got, bb.want)
		}
	}
}

func TestHourToSecond(t *testing.T) {
	type args struct {
		h time.Duration
	}
	tests := []struct {
		name string
		args args
		want time.Duration
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := HourToSecond(tt.args.h); got != tt.want {
			t.Errorf("%q. HourToSecond() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func BenchmarkHourToSecond(b *testing.B) {
	type args struct {
		h time.Duration
	}
	benchmarks := []struct {
		name string
		args args
		want time.Duration
	}{
	// TODO: Add benchmark cases.
	}
	for _, bb := range benchmarks {
		if got := HourToSecond(tt.args.h); got != bb.want {
			b.Errorf("%q. HourToSecond() = %v, want %v", tt.name, got, bb.want)
		}
	}
}
