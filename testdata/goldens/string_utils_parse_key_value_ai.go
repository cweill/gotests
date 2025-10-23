package testdata
import (
	"reflect"
	"testing"
)
func TestParseKeyValue(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]string
		wantErr bool
	}{
		{
			name: "valid_input",
			args: args{
				input: "hello,world",
			},
			want:    map[string]string{"hello": "world"},
			wantErr: false,
		},
		{
			name: "empty_string",
			args: args{
				input: "",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "negative_value",
			args: args{
				input: "-1,23",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseKeyValue(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Fatalf("ParseKeyValue() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseKeyValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
