package testdata

import "testing"

func TestFoo1(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "should return an empty slice when the input is empty",
		},
		{
			name: "should return a non-empty slice with one element when the input contains one element",
		},
		{
			name: "should return a non-empty slice with multiple elements when the input contains multiple elements",
		},
	}
	for range tests {
		Foo1()
	}
}
