package testdata

import (
	"go/ast"
	"go/types"
	"io"
	"testing"
)

func TestFoo24(t *testing.T) {
	tests := []struct {
		// Test description.
		name string
		// Parameters.
		r io.Reader
		x ast.Expr
		t types.Type
		// Expected results.
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if err := Foo24(tt.r, tt.x, tt.t); (err != nil) != tt.wantErr {
			t.Errorf("%q. Foo24() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}
