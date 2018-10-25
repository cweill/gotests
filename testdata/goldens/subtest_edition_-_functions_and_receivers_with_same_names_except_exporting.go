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
		t.Run(tt.name, func(t *testing.T) {
			got, err := SameName()
			if (err != nil) != tt.wantErr {
				t.Errorf("SameName() error = %v, wantErr %v", err, tt.wantErr)
				return
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
				t.Errorf("sameName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("sameName() = %v, want %v", got, tt.want)
			}
		})
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
		t.Run(tt.name, func(t *testing.T) {
			t := &SameTypeName{}
			got, err := t.SameName()
			if (err != nil) != tt.wantErr {
				t.Errorf("SameTypeName.SameName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SameTypeName.SameName() = %v, want %v", got, tt.want)
			}
		})
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
		t.Run(tt.name, func(t *testing.T) {
			t := &SameTypeName{}
			got, err := t.sameName()
			if (err != nil) != tt.wantErr {
				t.Errorf("SameTypeName.sameName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SameTypeName.sameName() = %v, want %v", got, tt.want)
			}
		})
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
		t.Run(tt.name, func(t *testing.T) {
			t := &sameTypeName{}
			got, err := t.SameName()
			if (err != nil) != tt.wantErr {
				t.Errorf("sameTypeName.SameName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("sameTypeName.SameName() = %v, want %v", got, tt.want)
			}
		})
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
		t.Run(tt.name, func(t *testing.T) {
			t := &sameTypeName{}
			got, err := t.sameName()
			if (err != nil) != tt.wantErr {
				t.Errorf("sameTypeName.sameName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("sameTypeName.sameName() = %v, want %v", got, tt.want)
			}
		})
	}
}
