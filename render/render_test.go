package render

import (
	"io"
	"testing"

	"github.com/cweill/gotests/models"
)

func TestTestCases(t *testing.T) {
	tests := []struct {
		name    string
		w       io.Writer
		f       *models.Function
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if err := TestCases(tt.w, tt.f); (err != nil) != tt.wantErr {
			t.Errorf("%v. TestCases() error = %v, wantErr: %v", tt.name, err, tt.wantErr)
		}
	}
}
