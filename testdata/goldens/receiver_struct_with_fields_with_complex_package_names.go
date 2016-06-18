package testdata

import (
	"go/types"
	"reflect"
	"testing"
)

func TestImporter_Foo35(t *testing.T) {
	tests := []struct {
		// Test description.
		name string
		// Receiver fields.
		rImporter types.Importer
		rField    *types.Var
		// Parameters.
		t types.Type
		// Expected results.
		want *types.Var
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		i := &Importer{
			Importer: tt.rImporter,
			Field:    tt.rField,
		}
		if got := i.Foo35(tt.t); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. Importer.Foo35() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
