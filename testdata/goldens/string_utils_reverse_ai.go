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
			name: "edge_case_1",
			args: args{
				s: "",
			},
			want: "",
		},
		{
			name: "boundary_value",
			args: args{
				s: "a",
			},
			want: "a",
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
