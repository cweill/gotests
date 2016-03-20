package testdata

import "testing"

func TestDoctorSayHello(t *testing.T) {
	tests := []struct {
		// Test description.
		name string
		// Receiver fields.
		rPerson      *Person
		rID          string
		rnumPatients int
		rstring      string
		// Parameters.
		r *Person
		// Expected results.
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		d := &Doctor{
			Person:      tt.rPerson,
			ID:          tt.rID,
			numPatients: tt.rnumPatients,
			string:      tt.rstring,
		}
		if got := d.SayHello(tt.r); got != tt.want {
			t.Errorf("%q. Doctor.SayHello() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
