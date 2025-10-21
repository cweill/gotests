package testdata

import "testing"

func TestFoo6(t *testing.T) {
	type args struct {
		i int
		b bool
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Foo6(tt.args.i, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Fatalf("Foo6() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			if got != tt.want {
				t.Errorf("Foo6() = %v, want %v", got, tt.want)
			}
		})
	}
}
