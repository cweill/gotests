package testdata
import "testing"
func TestCalculateDiscount(t *testing.T) {
	type args struct {
		price      float64
		percentage int
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{
			name: "valid_input",
			args: args{
				price:      10.5,
				percentage: 20,
			},
			want:    8.5,
			wantErr: false,
		},
		{
			name: "empty_string",
			args: args{
				price:      "",
				percentage: 30,
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "negative_value",
			args: args{
				price:      -10,
				percentage: 20,
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CalculateDiscount(tt.args.price, tt.args.percentage)
			if (err != nil) != tt.wantErr {
				t.Fatalf("CalculateDiscount() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			if got != tt.want {
				t.Errorf("CalculateDiscount() = %v, want %v", got, tt.want)
			}
		})
	}
}
