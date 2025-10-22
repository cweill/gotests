package testdata

import "testing"

func TestFoo20(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "empty_string",
			args: args{
				strs: nil,
			},
			want: "",
		},
		{
			name: "single_element_string",
			args: args{
				strs: nil,
			},
			want: "hello",
		},
		{
			name: "multiple_elements_string",
			args: args{
				strs: nil,
			},
			want: "worldgo",
		},
	}
	for _, tt := range tests {
		if got := Foo20(tt.args.strs...); got != tt.want {
			t.Errorf("%q. Foo20() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
