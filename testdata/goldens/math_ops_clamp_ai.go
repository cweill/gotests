package testdata

import "testing"

func TestClamp(t *testing.T) {
	type args struct {
		value int
		min   int
		max   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "positive_value",
			args: args{
				value: 10,
				min:   5,
				max:   20,
			},
			want: 10,
		},
		{
			name: "negative_value",
			args: args{
				value: -10,
				min:   5,
				max:   20,
			},
			want: -10,
		},
		{
			name: "within_range",
			args: args{
				value: 15,
				min:   5,
				max:   20,
			},
			want: 15,
		},
	}
	for _, tt := range tests {
		if got := Clamp(tt.args.value, tt.args.min, tt.args.max); got != tt.want {
			t.Errorf("%q. Clamp() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
