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
			name: "valid_input",
			args: args{
				numbers: []int{1, 2, -3, 4},
			},
			want: []int{1, 2, 4},
		},
		{
			name: "edge_case_1",
			args: args{
				numbers: []int{},
			},
			want: []int{},
		},
		{
			name: "boundary_value_1",
			args: args{
				numbers: []int{-5, -3, -1},
			},
			want: []int{-5, -3, -1},
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
