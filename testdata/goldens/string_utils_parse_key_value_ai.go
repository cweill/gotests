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
			name: "valid input",
			args: args{
				input: "key1=value1,key2=value2",
			},
			want:    map[string]string{"key1": "value1", "key2": "value2"},
			wantErr: false,
		},
		{
			name: "empty input",
			args: args{
				input: "",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "invalid format",
			args: args{
				input: "key=value,key3=value4",
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
