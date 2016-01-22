package testfiles

import "testing"

func TestBar100(t *testing.T) {
	tests := []struct {
		name    string
		b       *Bar
		i       interface{}
		wantErr bool
	}{
		{
			name:    "Basic test",
			b:       &Bar{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		if err := tt.b.Bar100(tt.i); (err != nil) != tt.wantErr {
			t.Errorf("%v. Bar100() error = %v, wantErr: %v", tt.name, err, tt.wantErr)
		}
	}
}

func TestBaz100(t *testing.T) {
	tests := []struct {
		name string
		f    *float64
		want float64
	}{
		{
			name: "Basic test",
			f:    func() *float64 { var x float64 = 64; return &x }(),
			want: 64,
		},
	}
	for _, tt := range tests {
		if got := baz100(tt.f); got != tt.want {
			t.Errorf("%v. baz100() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
