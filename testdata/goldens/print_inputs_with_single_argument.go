package testdata

import "testing"

func TestBarFoo7(t *testing.T) {
	tests := []struct {
		// Parameters.
		i int
		// Expected results.
		want    string
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		b := &Bar{}
		got, err := b.Foo7(tt.i)
		if (err != nil) != tt.wantErr {
			t.Errorf("Bar.Foo7(%v) error = %v, wantErr %v", tt.i, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("Bar.Foo7(%v) = %v, want %v", tt.i, got, tt.want)
		}
	}
}
