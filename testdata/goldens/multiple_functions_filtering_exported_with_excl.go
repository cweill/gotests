package testdata

import (
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

func (s *BarSuite) TestBarFilter() {
	type args struct {
		i interface{}
	}
	tests := []struct {
		name    string
		b       *Bar
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	t := s.T()
	for _, tt := range tests {
		b := &Bar{}
		if err := b.BarFilter(tt.args.i); (err != nil) != tt.wantErr {
			t.Errorf("%q. Bar.BarFilter() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}
