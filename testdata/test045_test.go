package testdata

import (
	"reflect"
	"testing"
)

func TestTransform(t *testing.T) {
	type args struct {
		input []int
		fn    func(int) int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Transform[int, int](tt.args.input, tt.args.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Transform() = %v, want %v", got, tt.want)
			}
		})
	}
}
