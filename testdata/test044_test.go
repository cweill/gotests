package testdata

import (
	"reflect"
	"testing"
)

func TestMap_Set(t *testing.T) {
	type args struct {
		key   string
		value int
	}
	tests := []struct {
		name string
		m    *Map[string, int]
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Map[string, int]{}
			m.Set(tt.args.key, tt.args.value)
		})
	}
}

func TestMap_Get(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name  string
		m     *Map[string, int]
		args  args
		want  int
		want1 bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Map[string, int]{}
			got, got1 := m.Get(tt.args.key)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map[K, V].Get() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Map[K, V].Get() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
