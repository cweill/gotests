package testdata

import "testing"

func TestHashPassword(t *testing.T) {
	type args struct {
		password string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "short_password",
			args: args{
				password: "a",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "long_password",
			args: args{
				password: "a1234567890",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "valid_password",
			args: args{
				password: "StrongPassword123!",
			},
			want:    "hashed_StrongPassword123!",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		got, err := HashPassword(tt.args.password)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. HashPassword() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if tt.wantErr {
			return
		}
		if got != tt.want {
			t.Errorf("%q. HashPassword() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
