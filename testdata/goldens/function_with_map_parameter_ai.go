package testdata

import (
	"reflect"
	"testing"
)

func TestFoo10(t *testing.T) {
	type args struct {
		m map[string]int32
	}
	tests := []struct {
		name string
		args args
		want map[string]*Bar
	}{
		{
			name: "test_Foo10",
			args: args{
				m: nil,
			},
			want: nil,
		},
		{
			name: "test_Foo10_with_zero_value",
			args: args{
				m: nil,
			},
			want: nil,
		},
		{
			name: "test_Foo10_with_negative_value",
			args: args{
				m: nil,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		if got := Foo10(tt.args.m); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. Foo10() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
