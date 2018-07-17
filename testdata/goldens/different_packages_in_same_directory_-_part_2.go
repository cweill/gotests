package foo

import "testing"

func TestFoo_Foo(t *testing.T) {
	type fields struct {
		Bar string
	}
	tests := []struct {
		name    string
		fields  fields
		arg     string
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		f := &Foo{
			Bar: tt.fields.Bar,
		}
		if err := f.Foo(tt.arg); (err != nil) != tt.wantErr {
			t.Errorf("%q. Foo.Foo() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}
