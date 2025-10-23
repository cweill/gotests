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
			name: "valid input",
			args: args{
				s: "hello",
			},
			want: "olleh",
		},
		{
			name: "empty input",
			args: args{
				s: "",
			},
			want: "",
		},
		{
			name: "nil input",
			args: args{
				s: nil,
			},
			want: "",
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
