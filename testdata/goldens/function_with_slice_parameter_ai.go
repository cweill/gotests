package testdata

import (
	"reflect"
	"testing"
)

func TestFoo11(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name    string
		args    args
		want    []*Bar
		wantErr bool
	}{
		{
			name: "test_Foo11",
			args: args{
				strs: nil,
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "test_Foo11_with_empty_string",
			args: args{
				strs: nil,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "test_Foo11_with_non_string_value",
			args: args{
				strs: nil,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		got, err := Foo11(tt.args.strs)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. Foo11() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if tt.wantErr {
			return
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. Foo11() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
