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
			name: "valid input",
			args: args{
				password: "secure_password123",
			},
			want:    "hashed_secure_password123",
			wantErr: false,
		},
		{
			name: "short password",
			args: args{
				password: "shortPassword",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "long password",
			args: args{
				password: "a" * 73,
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HashPassword(tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Fatalf("HashPassword() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			if got != tt.want {
				t.Errorf("HashPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}
