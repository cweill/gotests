package testdata

import "testing"

func TestFoo4(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{
			name: "positive_numbers",
			want: 8,
		},
		{
			name: "zero_values",
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Foo4(); got != tt.want {
				t.Errorf("Foo4() = %v, want %v", got, tt.want)
			}
		})
	}
}
