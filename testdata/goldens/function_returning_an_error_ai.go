package testdata

import "testing"

func TestFoo5(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		{
			name:    "test_case_1",
			want:    10,
			wantErr: false,
		},
		{
			name:    "test_case_2",
			want:    7,
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
