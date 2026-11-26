package external_test

import (
	"reflect"
	"testing"

	"github.com/cweill/gotests/testdata/packagetest/external"
)

func TestExisting(t *testing.T) {
	if got := external.Existing(); got != 0 {
		t.Errorf("Existing() = %v, want 0", got)
	}
}

func TestFind(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name string
		args args
		want *external.Item
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := external.Find(tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Find() = %v, want %v", got, tt.want)
			}
		})
	}
}
