package testdata

import "testing"

func TestFoo6(t *testing.T) {
	type args struct {
		i int
		b bool
	}
	tests := []struct {
		args    args
		want    string
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, err := Foo6(tt.args.i, tt.args.b)
		if (err != nil) != tt.wantErr {
			t.Errorf("Foo6(%v, %v) error = %v, wantErr %v", tt.args.i, tt.args.b, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("Foo6(%v, %v) = %v, want %v", tt.args.i, tt.args.b, got, tt.want)
		}
	}
}
