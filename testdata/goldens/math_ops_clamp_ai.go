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
			name: "valid input",
			args: args{
				value: 10,
				min:   5,
				max:   20,
			},
			want: 10,
		},
		{
			name: "edge case zero",
			args: args{
				value: 0,
				min:   5,
				max:   20,
			},
			want: 5,
		},
		{
			name: "edge case empty",
			args: args{
				value: "",
				min:   5,
				max:   20,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Clamp(tt.args.value, tt.args.min, tt.args.max); got != tt.want {
				t.Errorf("Clamp() = %v, want %v", got, tt.want)
			}
		})
	}
}
