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
