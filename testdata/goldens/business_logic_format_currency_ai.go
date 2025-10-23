package testdata
import "testing"
func TestFormatCurrency(t *testing.T) {
	type args struct {
		amount float64
		code   string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "descriptive_test_name",
			args: args{
				amount: 123.456789,
				code:   "USD",
			},
			want:    "$123.46",
			wantErr: false,
		},
		{
			name: "descriptive_test_name",
			args: args{
				amount: 987.654321,
				code:   "EUR",
			},
			want:    "€987.66",
			wantErr: false,
		},
		{
			name: "descriptive_test_name",
			args: args{
				amount: 0.0001,
				code:   "JPY",
			},
			want:    "¥0.01",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FormatCurrency(tt.args.amount, tt.args.code)
			if (err != nil) != tt.wantErr {
				t.Fatalf("FormatCurrency() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			if got != tt.want {
				t.Errorf("FormatCurrency() = %v, want %v", got, tt.want)
			}
		})
	}
}
