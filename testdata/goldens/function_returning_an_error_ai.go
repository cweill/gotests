package testdata

import "testing"

func TestFoo5(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		{
			name:    "positive_numbers",
			want:    8,
			wantErr: false,
		},
		{
			name:    "zero_values",
			want:    0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Foo5()
			if (err != nil) != tt.wantErr {
				t.Fatalf("Foo5() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			if got != tt.want {
				t.Errorf("Foo5() = %v, want %v", got, tt.want)
			}
		})
	}
}
