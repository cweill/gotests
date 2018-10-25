package testdata

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type DoctorSuite struct {
	suite.Suite
	Person      *Person
	ID          string
	numPatients int
	string      string
}

func TestDoctorSuite(t *testing.T) {
	suite.Run(t, new(DoctorSuite))
}

func (s *DoctorSuite) SetupTest()    {}
func (s *DoctorSuite) TearDownTest() {}

func (s *DoctorSuite) SetupSuite()    {}
func (s *DoctorSuite) TearDownSuite() {}

func (s *DoctorSuite) TestSayHello() {
	type fields struct {
		Person      *Person
		ID          string
		numPatients int
		string      string
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
		d := &Doctor{
			Person:      tt.fields.Person,
			ID:          tt.fields.ID,
			numPatients: tt.fields.numPatients,
			string:      tt.fields.string,
		}
		if got := d.SayHello(tt.args.r); got != tt.want {
			t.Errorf("%q. Doctor.SayHello() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
