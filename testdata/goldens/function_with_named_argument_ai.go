package testdata

import "testing"

func TestFoo3(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "empty_string",
			args: args{
				s: "",
			},
		},
		{
			name: "single_char",
			args: args{
				s: "a",
			},
		},
		{
			name: "longer_string",
			args: args{
				s: "hello world",
			},
		},
	}
	for _, tt := range tests {
		Foo3(tt.args.s)
	}
}
