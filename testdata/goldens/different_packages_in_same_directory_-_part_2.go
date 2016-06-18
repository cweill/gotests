package foo

import "testing"

func TestFoo_Foo(t *testing.T) {
	tests := []struct {
		// Test description.
		name string
		// Receiver fields.
		rBar string
		// Parameters.
		s string
		// Expected results.
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		f := &Foo{
			Bar: tt.rBar,
		}
		if err := f.Foo(tt.s); (err != nil) != tt.wantErr {
			t.Errorf("%q. Foo.Foo() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}
