package testdata

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/suite"
)

type BarSuite struct {
	suite.Suite
}

func TestBarSuite(t *testing.T) {
	suite.Run(t, new(BarSuite))
}

func (s *BarSuite) SetupTest()    {}
func (s *BarSuite) TearDownTest() {}

func (s *BarSuite) SetupSuite()    {}
func (s *BarSuite) TearDownSuite() {}

func (s *BarSuite) TestFoo9() {
	tests := []struct {
		name string
		b    Bar
		want Bar
	}{
		// TODO: Add test cases.
	}
	t := s.T()
	for _, tt := range tests {
		b := Bar{}
		if got := b.Foo9(); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. Bar.Foo9() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
