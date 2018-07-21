package testdata

import (
	"go/types"
	"reflect"
	"testing"
)

func TestImporter_Foo35(t *testing.T) {
	type fields struct {
		Importer types.Importer
		Field    *types.Var
	}
	tests := []struct {
		name   string
		fields fields
		arg    types.Type
		want   *types.Var
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		i := &Importer{
			Importer: tt.fields.Importer,
			Field:    tt.fields.Field,
		}
		if got := i.Foo35(tt.arg); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. Importer.Foo35() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
