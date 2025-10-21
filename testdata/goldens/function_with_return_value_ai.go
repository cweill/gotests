package testdata

import "testing"

func TestFoo4(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{
			name: "test_case_1",
			want: 10,
		},
		{
			name: "test_case_2",
			want: 7,
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
