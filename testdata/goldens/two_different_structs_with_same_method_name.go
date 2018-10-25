package testdata

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type BookSuite struct {
	suite.Suite
}

func TestBookSuite(t *testing.T) {
	suite.Run(t, new(BookSuite))
}

func (s *BookSuite) SetupTest()    {}
func (s *BookSuite) TearDownTest() {}

func (s *BookSuite) SetupSuite()    {}
func (s *BookSuite) TearDownSuite() {}

type DoorSuite struct {
	suite.Suite
}

func TestDoorSuite(t *testing.T) {
	suite.Run(t, new(DoorSuite))
}

func (s *DoorSuite) SetupTest()    {}
func (s *DoorSuite) TearDownTest() {}

func (s *DoorSuite) SetupSuite()    {}
func (s *DoorSuite) TearDownSuite() {}

type XmlSuite struct {
	suite.Suite
}

func TestXmlSuite(t *testing.T) {
	suite.Run(t, new(XmlSuite))
}

func (s *XmlSuite) SetupTest()    {}
func (s *XmlSuite) TearDownTest() {}

func (s *XmlSuite) SetupSuite()    {}
func (s *XmlSuite) TearDownSuite() {}

func (s *BookSuite) TestOpen() {
	tests := []struct {
		name    string
		b       *Book
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	t := s.T()
	for _, tt := range tests {
		b := &Book{}
		if err := b.Open(); (err != nil) != tt.wantErr {
			t.Errorf("%q. Book.Open() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}

func (s *DoorSuite) TestOpen() {
	tests := []struct {
		name    string
		d       *door
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	t := s.T()
	for _, tt := range tests {
		d := &door{}
		if err := d.Open(); (err != nil) != tt.wantErr {
			t.Errorf("%q. door.Open() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}

func (s *XmlSuite) TestOpen() {
	tests := []struct {
		name    string
		x       *xml
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	t := s.T()
	for _, tt := range tests {
		x := &xml{}
		if err := x.Open(); (err != nil) != tt.wantErr {
			t.Errorf("%q. xml.Open() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}
