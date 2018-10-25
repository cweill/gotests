package testdata

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type InitFuncStructSuite struct {
	suite.Suite
	field int
}

func TestInitFuncStructSuite(t *testing.T) {
	suite.Run(t, new(InitFuncStructSuite))
}

func (s *InitFuncStructSuite) SetupTest()    {}
func (s *InitFuncStructSuite) TearDownTest() {}

func (s *InitFuncStructSuite) SetupSuite()    {}
func (s *InitFuncStructSuite) TearDownSuite() {}

type InitFieldStructSuite struct {
	suite.Suite
	init int
}

func TestInitFieldStructSuite(t *testing.T) {
	suite.Run(t, new(InitFieldStructSuite))
}

func (s *InitFieldStructSuite) SetupTest()    {}
func (s *InitFieldStructSuite) TearDownTest() {}

func (s *InitFieldStructSuite) SetupSuite()    {}
func (s *InitFieldStructSuite) TearDownSuite() {}

func (s *InitFuncStructSuite) Test_init() {
	type fields struct {
		field int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	t := s.T()
	for _, tt := range tests {
		i := initFuncStruct{
			field: tt.fields.field,
		}
		if got := i.init(); got != tt.want {
			t.Errorf("%q. initFuncStruct.init() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func (s *InitFieldStructSuite) Test_getInit() {
	type fields struct {
		init int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	t := s.T()
	for _, tt := range tests {
		i := initFieldStruct{
			init: tt.fields.init,
		}
		if got := i.getInit(); got != tt.want {
			t.Errorf("%q. initFieldStruct.getInit() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
