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
			name: "valid input",
			args: args{
				s:          "hello",
				substrings: []string{"world"},
			},
			want: true,
		},
		{
			name: "empty string",
			args: args{
				s:          "",
				substrings: []string{},
			},
			want: false,
		},
		{
			name: "nil slice",
			args: args{
				s:          nil,
				substrings: []string{},
			},
			want: false,
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
