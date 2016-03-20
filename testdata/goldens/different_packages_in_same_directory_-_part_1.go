package bar

import "testing"

func TestBarBar(t *testing.T) {
	tests := []struct {
		name    string
		foo     string
		s       string
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		b := &Bar{
			Foo: tt.foo,
		}
		if err := b.Bar(tt.s); (err != nil) != tt.wantErr {
			t.Errorf("%q. Bar.Bar() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}
