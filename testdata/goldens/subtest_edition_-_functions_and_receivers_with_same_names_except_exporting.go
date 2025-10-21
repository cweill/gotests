package testdata

import "testing"

func TestSameName(t *testing.T) {
	tests := []struct {
		name    string
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SameName()
			if (err != nil) != tt.wantErr {
				t.Fatalf("SameName() error = %v, wantErr %v", err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("SameName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sameName(t *testing.T) {
	tests := []struct {
		name    string
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := sameName()
			if (err != nil) != tt.wantErr {
				t.Fatalf("sameName() error = %v, wantErr %v", err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("sameName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSameTypeName_SameName(t *testing.T) {
	tests := []struct {
		name    string
		tr      *SameTypeName
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &SameTypeName{}
			got, err := tr.SameName()
			if (err != nil) != tt.wantErr {
				t.Fatalf("SameTypeName.SameName() error = %v, wantErr %v", err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("SameTypeName.SameName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSameTypeName_sameName(t *testing.T) {
	tests := []struct {
		name    string
		tr      *SameTypeName
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &SameTypeName{}
			got, err := tr.sameName()
			if (err != nil) != tt.wantErr {
				t.Fatalf("SameTypeName.sameName() error = %v, wantErr %v", err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("SameTypeName.sameName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sameTypeName_SameName(t *testing.T) {
	tests := []struct {
		name    string
		tr      *sameTypeName
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &sameTypeName{}
			got, err := tr.SameName()
			if (err != nil) != tt.wantErr {
				t.Fatalf("sameTypeName.SameName() error = %v, wantErr %v", err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("sameTypeName.SameName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sameTypeName_sameName(t *testing.T) {
	tests := []struct {
		name    string
		tr      *sameTypeName
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &sameTypeName{}
			got, err := tr.sameName()
			if (err != nil) != tt.wantErr {
				t.Fatalf("sameTypeName.sameName() error = %v, wantErr %v", err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("sameTypeName.sameName() = %v, want %v", got, tt.want)
			}
		})
	}
}
