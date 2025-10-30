package testdata

import "testing"

func TestCalculator_Multiply(t *testing.T) {
	type args struct {
		n int
		d int
	}
	tests := []struct {
		name string
		c    *Calculator
		args args
		want int
	}{
		{
			name: "valid input",
			c:    nil, // TODO: Initialize receiver from AI case
			args: args{
				n: 5,
				d: 3,
			},
			want: 15,
		},
		{
			name: "edge case zero",
			c:    nil, // TODO: Initialize receiver from AI case
			args: args{
				n: 0,
				d: 3,
			},
			want: 0,
		},
		{
			name: "edge case empty",
			c:    nil, // TODO: Initialize receiver from AI case
			args: args{
				n: 0,
				d: "",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Multiply(tt.args.n, tt.args.d); got != tt.want {
				t.Errorf("Calculator.Multiply() = %v, want %v", got, tt.want)
			}
		})
	}
}
