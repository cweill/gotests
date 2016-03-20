package testdata

import "testing"

func TestDoctorSayHello(t *testing.T) {
	tests := []struct {
		name        string
		person      *Person
		id          string
		numPatients int
		string      string
		r           *Person
		want        string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		d := &Doctor{
			Person:      tt.person,
			ID:          tt.id,
			numPatients: tt.numPatients,
			string:      tt.string,
		}
		if got := d.SayHello(tt.r); got != tt.want {
			t.Errorf("%q. Doctor.SayHello() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
