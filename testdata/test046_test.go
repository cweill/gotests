package testdata

import "testing"

func TestFindFirst(t *testing.T) {
	type args struct {
		slice  []string
		target string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FindFirst[string](tt.args.slice, tt.args.target)
			if (err != nil) != tt.wantErr {
				t.Fatalf("FindFirst() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			if got != tt.want {
				t.Errorf("FindFirst() = %v, want %v", got, tt.want)
			}
		})
	}
}
