package testdata

import "testing"

func TestSameName(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for range tests {
		if err := SameName(); (err != nil) != tt.wantErr {
			t.Errorf("%q. SameName() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}

func Test_sameName(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for range tests {
		if err := sameName(); (err != nil) != tt.wantErr {
			t.Errorf("%q. sameName() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}

func TestSameTypeName_SameName(t *testing.T) {
	tests := []struct {
		name    string
		t       *SameTypeName
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t := &SameTypeName{}
		if err := t.SameName(); (err != nil) != tt.wantErr {
			t.Errorf("%q. SameTypeName.SameName() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}

func TestSameTypeName_sameName(t *testing.T) {
	tests := []struct {
		name    string
		t       *SameTypeName
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t := &SameTypeName{}
		if err := t.sameName(); (err != nil) != tt.wantErr {
			t.Errorf("%q. SameTypeName.sameName() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}

func Test_sameTypeName_SameName(t *testing.T) {
	tests := []struct {
		name    string
		t       *sameTypeName
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t := &sameTypeName{}
		if err := t.SameName(); (err != nil) != tt.wantErr {
			t.Errorf("%q. sameTypeName.SameName() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}

func Test_sameTypeName_sameName(t *testing.T) {
	tests := []struct {
		name    string
		t       *sameTypeName
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t := &sameTypeName{}
		if err := t.sameName(); (err != nil) != tt.wantErr {
			t.Errorf("%q. sameTypeName.sameName() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}
