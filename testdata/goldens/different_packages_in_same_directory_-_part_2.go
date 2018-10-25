package foo

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type FooSuite struct {
	suite.Suite
	Bar string
}

func TestFooSuite(t *testing.T) {
	suite.Run(t, new(FooSuite))
}

func (s *FooSuite) SetupTest()    {}
func (s *FooSuite) TearDownTest() {}

func (s *FooSuite) SetupSuite()    {}
func (s *FooSuite) TearDownSuite() {}

func (s *FooSuite) TestFoo() {
	type fields struct {
		Bar string
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
		f := &Foo{
			Bar: tt.fields.Bar,
		}
		if err := f.Foo(tt.args.s); (err != nil) != tt.wantErr {
			t.Errorf("%q. Foo.Foo() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}
