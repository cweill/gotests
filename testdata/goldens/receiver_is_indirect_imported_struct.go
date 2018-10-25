package testdata

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type SomeIndirectImportedStructSuite struct {
	suite.Suite
}

func TestSomeIndirectImportedStructSuite(t *testing.T) {
	suite.Run(t, new(SomeIndirectImportedStructSuite))
}

func (s *SomeIndirectImportedStructSuite) SetupTest()    {}
func (s *SomeIndirectImportedStructSuite) TearDownTest() {}

func (s *SomeIndirectImportedStructSuite) SetupSuite()    {}
func (s *SomeIndirectImportedStructSuite) TearDownSuite() {}

func (s *SomeIndirectImportedStructSuite) TestFoo037() {
	tests := []struct {
		name string
		smtg *someIndirectImportedStruct
	}{
		// TODO: Add test cases.
	}
	t := s.T()
	for _, tt := range tests {
		smtg := &someIndirectImportedStruct{}
		smtg.Foo037()
	}
}
