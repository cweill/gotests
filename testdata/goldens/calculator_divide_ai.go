package testdata

import "testing"

func TestCalculator_Divide(t *testing.T) {
	type args struct {
		n int
		d int
	}
	tests := []struct {
		name    string
		c       *Calculator
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "valid input",
			c:    nil, // TODO: Initialize receiver from AI case
			args: args{
				n: 10,
				d: 2,
			},
			want:    5,
			wantErr: false,
		},
		{
			name: "edge case: division by zero",
			c:    nil, // TODO: Initialize receiver from AI case
			args: args{
				n: 10,
				d: 0,
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "error case: invalid input",
			c:    nil, // TODO: Initialize receiver from AI case
			args: args{
				n: -5,
				d: 2,
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Divide(tt.args.n, tt.args.d)
			if (err != nil) != tt.wantErr {
				t.Fatalf("Calculator.Divide() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			if got != tt.want {
				t.Errorf("Calculator.Divide() = %v, want %v", got, tt.want)
			}
		})
	}
}
