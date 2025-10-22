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
		got, err := FormatCurrency(tt.args.amount, tt.args.code)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. FormatCurrency() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if tt.wantErr {
			return
		}
		if got != tt.want {
			t.Errorf("%q. FormatCurrency() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
