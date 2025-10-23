package testdata
import "testing"
func TestReverse(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "valid_input",
			args: args{
				s: "hello",
			},
			want: "olleh",
		},
		{
			name: "empty_string",
			args: args{
				s: "",
			},
			want: "",
		},
		{
			name: "negative_value",
			args: args{
				s: "-123",
			},
			want: "-123",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reverse(tt.args.s); got != tt.want {
				t.Errorf("Reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
