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

func TestFooFilter(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name    string
		args    args
		want    []*Bar
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, err := FooFilter(tt.args.strs)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. FooFilter() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. FooFilter() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

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
