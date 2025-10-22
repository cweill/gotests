package testdata

import "testing"

func TestGetExtension(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "empty_string",
			args: args{
				filename: "",
			},
			want: "",
		},
		{
			name: "single_dot_file",
			args: args{
				filename: "file.txt",
			},
			want: "txt",
		},
		{
			name: "multiple_dot_file",
			args: args{
				filename: "document.docx",
			},
			want: "docx",
		},
	}
	for _, tt := range tests {
		if got := GetExtension(tt.args.filename); got != tt.want {
			t.Errorf("%q. GetExtension() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
