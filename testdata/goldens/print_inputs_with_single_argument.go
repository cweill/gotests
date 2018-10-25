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

func (s *BarSuite) TestFoo7() {
	type args struct {
		i int
	}
	tests := []struct {
		name    string
		b       *Bar
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	t := s.T()
	for _, tt := range tests {
		b := &Bar{}
		got, err := b.Foo7(tt.args.i)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. Bar.Foo7(%v) error = %v, wantErr %v", tt.name, tt.args.i, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("%q. Bar.Foo7(%v) = %v, want %v", tt.name, tt.args.i, got, tt.want)
		}
	}
}
