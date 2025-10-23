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
			name: "valid input",
			args: args{
				n: 5,
			},
			want:    120,
			wantErr: false,
		},
		{
			name: "edge case 1",
			args: args{
				n: -1,
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "error case 1",
			args: args{
				n: 21,
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Factorial(tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Fatalf("Factorial() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			if got != tt.want {
				t.Errorf("Factorial() = %v, want %v", got, tt.want)
			}
		})
	}
}
