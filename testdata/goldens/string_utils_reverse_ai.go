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
			name: "Reverse",
			args: args{
				s: "hello",
			},
			want: "olleh",
		},
		{
			name: "Reverse",
			args: args{
				s: "world",
			},
			want: "dlrow",
		},
		{
			name: "Reverse",
			args: args{
				s: "",
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
