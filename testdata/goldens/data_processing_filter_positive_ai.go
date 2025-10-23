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
				numbers: []int{-1, 2, -3, 4, 5},
			},
			want: []int{2, 4, 5},
		},
		{
			name: "FilterPositive",
			args: args{
				numbers: []int{},
			},
			want: []int{},
		},
		{
			name: "FilterPositive",
			args: args{
				numbers: []int{-10, -20, -30},
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FilterPositive(tt.args.numbers); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilterPositive() = %v, want %v", got, tt.want)
			}
		})
	}
}
