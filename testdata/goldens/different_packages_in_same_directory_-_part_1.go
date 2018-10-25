package bar

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type BarSuite struct {
	suite.Suite
	Foo string
}

func TestBarSuite(t *testing.T) {
	suite.Run(t, new(BarSuite))
}

func (s *BarSuite) SetupTest()    {}
func (s *BarSuite) TearDownTest() {}

func (s *BarSuite) SetupSuite()    {}
func (s *BarSuite) TearDownSuite() {}

func (s *BarSuite) TestBar() {
	type fields struct {
		Foo string
	}
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	t := s.T()
	for _, tt := range tests {
		b := &Bar{
			Foo: tt.fields.Foo,
		}
		if err := b.Bar(tt.args.s); (err != nil) != tt.wantErr {
			t.Errorf("%q. Bar.Bar() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}
