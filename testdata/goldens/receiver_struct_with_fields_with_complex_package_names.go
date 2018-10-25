package testdata

import (
	"go/types"
	"reflect"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ImporterSuite struct {
	suite.Suite
	Importer types.Importer
	Field    *types.Var
}

func TestImporterSuite(t *testing.T) {
	suite.Run(t, new(ImporterSuite))
}

func (s *ImporterSuite) SetupTest()    {}
func (s *ImporterSuite) TearDownTest() {}

func (s *ImporterSuite) SetupSuite()    {}
func (s *ImporterSuite) TearDownSuite() {}

func (s *ImporterSuite) TestFoo35() {
	type fields struct {
		Importer types.Importer
		Field    *types.Var
	}
	type args struct {
		t types.Type
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *types.Var
	}{
		// TODO: Add test cases.
	}
	t := s.T()
	for _, tt := range tests {
		i := &Importer{
			Importer: tt.fields.Importer,
			Field:    tt.fields.Field,
		}
		if got := i.Foo35(tt.args.t); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. Importer.Foo35() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
