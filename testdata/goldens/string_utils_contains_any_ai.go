package testdata
import "testing"
func TestContainsAny(t *testing.T) {
	type args struct {
		s          string
		substrings []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "valid_input",
			args: args{
				s:          "hello",
				substrings: []string{"world", "test"},
			},
			want: true,
		},
		{
			name: "edge_case_1",
			args: args{
				s:          "",
				substrings: []string{},
			},
			want: false,
		},
		{
			name: "boundary_value",
			args: args{
				s:          "hello",
				substrings: []string{"world", "test", "end"},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsAny(tt.args.s, tt.args.substrings); got != tt.want {
				t.Errorf("ContainsAny() = %v, want %v", got, tt.want)
			}
		})
	}
}
