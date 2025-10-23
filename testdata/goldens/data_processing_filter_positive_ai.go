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
			name: "valid input",
			args: args{
				numbers: []int{1, 2, -3, 4},
			},
			want: []int{1, 2, 4},
		},
		{
			name: "edge case zero",
			args: args{
				numbers: []int{},
			},
			want: []int{},
		},
		{
			name: "edge case empty",
			args: args{
				numbers: []int{0},
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
