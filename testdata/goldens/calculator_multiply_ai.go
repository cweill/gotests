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
			name: "Multiply_1",
			c:    nil, // TODO: Initialize receiver from AI case
			args: args{
				n: 5,
				d: 3,
			},
			want: 15,
		},
		{
			name: "Multiply_2",
			c:    nil, // TODO: Initialize receiver from AI case
			args: args{
				n: -4,
				d: 6,
			},
			want: -24,
		},
		{
			name: "Multiply_3",
			c:    nil, // TODO: Initialize receiver from AI case
			args: args{
				n: 0,
				d: 5,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Calculator{}
			if got := c.Multiply(tt.args.n, tt.args.d); got != tt.want {
				t.Errorf("Calculator.Multiply() = %v, want %v", got, tt.want)
			}
		})
	}
}
