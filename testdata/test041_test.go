package testdata

import (
	"reflect"
	"testing"
)

func TestSumIntsOrFloats(t *testing.T) {
	type args struct {
		m map[string]int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumIntsOrFloats[string, int64](tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SumIntsOrFloats() = %v, want %v", got, tt.want)
			}
		})
	}
}
