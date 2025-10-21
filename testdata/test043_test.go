package testdata

import (
	"reflect"
	"testing"
)

func TestPair(t *testing.T) {
	type args struct {
		first  int
		second int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Pair[int, int](tt.args.first, tt.args.second)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Pair() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Pair() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
