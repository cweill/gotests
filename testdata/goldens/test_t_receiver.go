package testdata

import "testing"

func TestTestReceiver_FooMethod(t *testing.T) {
	tests := []struct {
		name    string
		tr      *TestReceiver
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		tr := &TestReceiver{}
		if err := tr.FooMethod(); (err != nil) != tt.wantErr {
			t.Errorf("%q. TestReceiver.FooMethod() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}
