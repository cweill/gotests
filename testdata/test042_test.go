package testdata

import "testing"

func TestSet_Add(t *testing.T) {
	type args struct {
		v string
	}
	tests := []struct {
		name string
		s    *Set[string]
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Set[string]{}
			s.Add(tt.args.v)
		})
	}
}

func TestSet_Has(t *testing.T) {
	type args struct {
		v string
	}
	tests := []struct {
		name string
		s    *Set[string]
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Set[string]{}
			if got := s.Has(tt.args.v); got != tt.want {
				t.Errorf("Set[T].Has() = %v, want %v", got, tt.want)
			}
		})
	}
}
