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
			name: "valid_input",
			args: args{
				n: 5,
			},
			want:    120,
			wantErr: false,
		},
		{
			name: "empty_string",
			args: args{
				n: "",
			},
			want:    1,
			wantErr: true,
		},
		{
			name: "negative_value",
			args: args{
				n: -3,
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
