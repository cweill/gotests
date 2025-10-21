package testdata

import "testing"

func TestFoo6(t *testing.T) {
	type args struct {
		i int
		b bool
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "test_Foo6",
			args: args{
				i: 10,
				b: true,
			},
			want:    "true",
			wantErr: false,
		},
		{
			name: "test_Foo6",
			args: args{
				i: -5,
				b: false,
			},
			want:    "-5",
			wantErr: false,
		},
		{
			name: "test_Foo6",
			args: args{
				i: 0,
				b: true,
			},
			want:    "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		got, err := Foo6(tt.args.i, tt.args.b)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. Foo6() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if tt.wantErr {
			return
		}
		if got != tt.want {
			t.Errorf("%q. Foo6() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
