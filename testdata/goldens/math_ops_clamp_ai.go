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
			name: "valid_input",
			args: args{
				value: 10,
				min:   5,
				max:   20,
			},
			want: 15,
		},
		{
			name: "empty_string",
			args: args{
				value: "",
				min:   5,
				max:   20,
			},
			want: 5,
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Clamp(tt.args.value, tt.args.min, tt.args.max); got != tt.want {
				t.Errorf("Clamp() = %v, want %v", got, tt.want)
			}
		})
	}
}
