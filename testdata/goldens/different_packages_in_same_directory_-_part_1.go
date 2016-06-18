package bar

import "testing"

func TestBar_Bar(t *testing.T) {
	tests := []struct {
		// Test description.
		name string
		// Receiver fields.
		rFoo string
		// Parameters.
		s string
		// Expected results.
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		b := &Bar{
			Foo: tt.rFoo,
		}
		if err := b.Bar(tt.s); (err != nil) != tt.wantErr {
			t.Errorf("%q. Bar.Bar() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}
