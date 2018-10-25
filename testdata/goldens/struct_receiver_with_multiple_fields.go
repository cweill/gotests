package testdata

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type PersonSuite struct {
	suite.Suite
	FirstName string
	LastName  string
	Age       int
	Gender    string
	Siblings  []*Person
}

func TestPersonSuite(t *testing.T) {
	suite.Run(t, new(PersonSuite))
}

func (s *PersonSuite) SetupTest()    {}
func (s *PersonSuite) TearDownTest() {}

func (s *PersonSuite) SetupSuite()    {}
func (s *PersonSuite) TearDownSuite() {}

func (s *PersonSuite) TestSayHello() {
	type fields struct {
		FirstName string
		LastName  string
		Age       int
		Gender    string
		Siblings  []*Person
	}
	type args struct {
		r *Person
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
	}
	t := s.T()
	for _, tt := range tests {
		p := &Person{
			FirstName: tt.fields.FirstName,
			LastName:  tt.fields.LastName,
			Age:       tt.fields.Age,
			Gender:    tt.fields.Gender,
			Siblings:  tt.fields.Siblings,
		}
		if got := p.SayHello(tt.args.r); got != tt.want {
			t.Errorf("%q. Person.SayHello() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
