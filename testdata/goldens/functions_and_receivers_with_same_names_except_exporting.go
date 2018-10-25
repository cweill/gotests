package testdata

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type SameTypeNameSuite struct {
	suite.Suite
}

func TestSameTypeNameSuite(t *testing.T) {
	suite.Run(t, new(SameTypeNameSuite))
}

func (s *SameTypeNameSuite) SetupTest()    {}
func (s *SameTypeNameSuite) TearDownTest() {}

func (s *SameTypeNameSuite) SetupSuite()    {}
func (s *SameTypeNameSuite) TearDownSuite() {}

func TestSameName(t *testing.T) {
	tests := []struct {
		name    string
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, err := SameName()
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. SameName() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("%q. SameName() = %v, want %v", tt.name, got, tt.want)
		}
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
		got, err := sameName()
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. sameName() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("%q. sameName() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func (s *SameTypeNameSuite) TestSameName() {
	tests := []struct {
		name    string
		t       *SameTypeName
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	t := s.T()
	for _, tt := range tests {
		t := &SameTypeName{}
		got, err := t.SameName()
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. SameTypeName.SameName() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("%q. SameTypeName.SameName() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func (s *SameTypeNameSuite) Test_sameName() {
	tests := []struct {
		name    string
		t       *SameTypeName
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	t := s.T()
	for _, tt := range tests {
		t := &SameTypeName{}
		got, err := t.sameName()
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. SameTypeName.sameName() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("%q. SameTypeName.sameName() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func (s *SameTypeNameSuite) TestSameName() {
	tests := []struct {
		name    string
		t       *sameTypeName
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	t := s.T()
	for _, tt := range tests {
		t := &sameTypeName{}
		got, err := t.SameName()
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. sameTypeName.SameName() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("%q. sameTypeName.SameName() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func (s *SameTypeNameSuite) Test_sameName() {
	tests := []struct {
		name    string
		t       *sameTypeName
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	t := s.T()
	for _, tt := range tests {
		t := &sameTypeName{}
		got, err := t.sameName()
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. sameTypeName.sameName() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("%q. sameTypeName.sameName() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
