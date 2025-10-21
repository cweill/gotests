package testdata

import (
	"reflect"
	"testing"
)

func TestGenericAny(t *testing.T) {
	type args struct {
		val int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenericAny[int](tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenericAny() = %v, want %v", got, tt.want)
			}
		})
	}
}
