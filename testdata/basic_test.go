package testdata

import "testing"

func TestSum(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "valid input positive numbers",
			args: args{
				a: 1,
				b: 2,
			},
			want: 3,
		},
		{
			name: "valid input negative numbers",
			args: args{
				a: -1,
				b: -2,
			},
			want: -3,
		},
		{
			name: "one positive one negative",
			args: args{
				a: 5,
				b: -3,
			},
			want: 2,
		},
		{
			name: "zero and positive",
			args: args{
				a: 0,
				b: 7,
			},
			want: 7,
		},
		{
			name: "zero and negative",
			args: args{
				a: 0,
				b: -4,
			},
			want: -4,
		},
		{
			name: "both zero",
			args: args{
				a: 0,
				b: 0,
			},
			want: 0,
		},
		{
			name: "large positive numbers",
			args: args{
				a: 1000000,
				b: 2000000,
			},
			want: 3000000,
		},
		{
			name: "large negative numbers",
			args: args{
				a: -1000000,
				b: -2000000,
			},
			want: -3000000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}
