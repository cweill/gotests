package testdata

import (
	"reflect"
	"testing"
)

func TestFilterPositive(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "FilterPositive",
			args: args{
				numbers: nil,
			},
			want: nil,
		},
		{
			name: "FilterPositive",
			args: args{
				numbers: nil,
			},
			want: nil,
		},
		{
			name: "FilterPositive",
			args: args{
				numbers: nil,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		if got := FilterPositive(tt.args.numbers); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. FilterPositive() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
