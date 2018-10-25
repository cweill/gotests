package testdata

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type CelsiusSuite struct {
	suite.Suite
}

func TestCelsiusSuite(t *testing.T) {
	suite.Run(t, new(CelsiusSuite))
}

func (s *CelsiusSuite) SetupTest()    {}
func (s *CelsiusSuite) TearDownTest() {}

func (s *CelsiusSuite) SetupSuite()    {}
func (s *CelsiusSuite) TearDownSuite() {}

type FahrenheitSuite struct {
	suite.Suite
}

func TestFahrenheitSuite(t *testing.T) {
	suite.Run(t, new(FahrenheitSuite))
}

func (s *FahrenheitSuite) SetupTest()    {}
func (s *FahrenheitSuite) TearDownTest() {}

func (s *FahrenheitSuite) SetupSuite()    {}
func (s *FahrenheitSuite) TearDownSuite() {}

func (s *CelsiusSuite) TestString() {
	tests := []struct {
		name string
		c    Celsius
		want string
	}{
		// TODO: Add test cases.
	}
	t := s.T()
	for _, tt := range tests {
		if got := tt.c.String(); got != tt.want {
			t.Errorf("%q. Celsius.String() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func (s *FahrenheitSuite) TestString() {
	tests := []struct {
		name string
		f    Fahrenheit
		want string
	}{
		// TODO: Add test cases.
	}
	t := s.T()
	for _, tt := range tests {
		if got := tt.f.String(); got != tt.want {
			t.Errorf("%q. Fahrenheit.String() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
