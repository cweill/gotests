package undefinedtypes

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/suite"
)

type UndefinedSuite struct {
	suite.Suite
}

func TestUndefinedSuite(t *testing.T) {
	suite.Run(t, new(UndefinedSuite))
}

func (s *UndefinedSuite) SetupTest()    {}
func (s *UndefinedSuite) TearDownTest() {}

func (s *UndefinedSuite) SetupSuite()    {}
func (s *UndefinedSuite) TearDownSuite() {}

func (s *UndefinedSuite) TestDo() {
	type args struct {
		es Something
	}
	tests := []struct {
		name    string
		u       *Undefined
		args    args
		want    *Unknown
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	t := s.T()
	for _, tt := range tests {
		got, err := tt.u.Do(tt.args.es)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. Undefined.Do() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. Undefined.Do() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
