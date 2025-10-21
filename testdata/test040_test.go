package testdata

import "testing"

func TestGenericComparable(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenericComparable[string](tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("GenericComparable() = %v, want %v", got, tt.want)
			}
		})
	}
}
