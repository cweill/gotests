package testdata

import "testing"

func TestFoo4(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{
			name: "should return false for an empty slice",
			want: false,
		},
		{
			name: "should return true for a slice with one element",
			want: true,
		},
		{
			name: "should return true for a slice with multiple elements",
			want: true,
		},
	}
	for _, tt := range tests {
		if got := Foo4(); got != tt.want {
			t.Errorf("%q. Foo4() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
