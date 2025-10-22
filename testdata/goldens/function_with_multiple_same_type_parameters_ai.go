package testdata

import "testing"

func TestFoo19(t *testing.T) {
	type args struct {
		in1 string
		in2 string
		in3 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "empty strings",
			args: args{
				in1: "",
				in2: "",
				in3: "",
			},
			want: "",
		},
		{
			name: "single empty string",
			args: args{
				in1: " ",
				in2: "",
				in3: "",
			},
			want: " ",
		},
		{
			name: "multiple empty strings",
			args: args{
				in1: "",
				in2: "",
				in3: "",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		if got := Foo19(tt.args.in1, tt.args.in2, tt.args.in3); got != tt.want {
			t.Errorf("%q. Foo19() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
