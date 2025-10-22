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
			name: "ParseKeyValue",
			args: args{
				input: "key1=value2,key3=value4",
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "ParseKeyValue",
			args: args{
				input: "key1,value2,key3=value4,key5=value6",
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "ParseKeyValue",
			args: args{
				input: "",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		got, err := ParseKeyValue(tt.args.input)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. ParseKeyValue() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if tt.wantErr {
			return
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. ParseKeyValue() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
