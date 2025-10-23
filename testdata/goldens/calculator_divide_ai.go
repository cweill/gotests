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
			name: "divide_by_zero",
			c:    nil, // TODO: Initialize receiver from AI case
			args: args{
				n: 10,
				d: 0,
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "positive_division",
			c:    nil, // TODO: Initialize receiver from AI case
			args: args{
				n: 5,
				d: 3,
			},
			want:    1.67,
			wantErr: false,
		},
		{
			name: "negative_division",
			c:    nil, // TODO: Initialize receiver from AI case
			args: args{
				n: -5,
				d: 3,
			},
			want:    -1.67,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Calculator{}
			got, err := c.Divide(tt.args.n, tt.args.d)
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
