package testdata

import "testing"

func TestValidateEmail(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "valid_email",
			args: args{
				email: "example@example.com",
			},
			wantErr: false,
		},
		{
			name: "invalid_email",
			args: args{
				email: "invalid-email@domain.com",
			},
			wantErr: true,
		},
		{
			name: "empty_email",
			args: args{
				email: "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		if err := ValidateEmail(tt.args.email); (err != nil) != tt.wantErr {
			t.Errorf("%q. ValidateEmail() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}
