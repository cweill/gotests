package foo

import "testing"

func TestFooFoo(t *testing.T) {
	tests := []struct {
		name    string
		bar     string
		s       string
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		f := &Foo{
			Bar: tt.bar,
		}
		if err := f.Foo(tt.s); (err != nil) != tt.wantErr {
			t.Errorf("%q. Foo.Foo() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}
