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
			name: "test_contains_any",
			args: args{
				s:          "hello",
				substrings: []string{"world", "test"},
			},
			want: true,
		},
		{
			name: "test_contains_any_empty_string",
			args: args{
				s:          "",
				substrings: []string{"world", "test"},
			},
			want: false,
		},
		{
			name: "test_contains_any_no_substrings",
			args: args{
				s:          "hello",
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
