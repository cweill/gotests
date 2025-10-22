package testdata

import "testing"

func TestFactorial(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "descriptive_test_name",
			args: args{
				n: 0,
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "descriptive_test_name",
			args: args{
				n: 1,
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "descriptive_test_name",
			args: args{
				n: 2,
			},
			want:    2,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		got, err := Factorial(tt.args.n)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. Factorial() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if tt.wantErr {
			return
		}
		if got != tt.want {
			t.Errorf("%q. Factorial() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
