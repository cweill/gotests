package testdata

import (
	"go/types"
	"reflect"
	"testing"
)

func TestImporterFoo35(t *testing.T) {
	tests := []struct {
		name     string
		importer types.Importer
		field    *types.Var
		t        types.Type
		want     *types.Var
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		i := &Importer{
			Importer: tt.importer,
			Field:    tt.field,
		}
		if got := i.Foo35(tt.t); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. Importer.Foo35() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
