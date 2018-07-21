package testdata

import "testing"

func TestBar_Foo7(t *testing.T) {
	tests := []struct {
		name    string
		b       *Bar
		arg     int
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		b := &Bar{}
		got, err := b.Foo7(tt.arg)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. Bar.Foo7(%v) error = %v, wantErr %v", tt.name, tt.arg, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("%q. Bar.Foo7(%v) = %v, want %v", tt.name, tt.arg, got, tt.want)
		}
	}
}
