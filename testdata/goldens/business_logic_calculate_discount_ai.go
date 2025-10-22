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
			name: "valid_price",
			args: args{
				price:      10.5,
				percentage: 20,
			},
			want:    8.5,
			wantErr: false,
		},
		{
			name: "invalid_percentage",
			args: args{
				price:      10.5,
				percentage: -10,
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "valid_price_and_percentage",
			args: args{
				price:      10.5,
				percentage: 20,
			},
			want:    8.5,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		got, err := CalculateDiscount(tt.args.price, tt.args.percentage)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. CalculateDiscount() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if tt.wantErr {
			return
		}
		if got != tt.want {
			t.Errorf("%q. CalculateDiscount() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
