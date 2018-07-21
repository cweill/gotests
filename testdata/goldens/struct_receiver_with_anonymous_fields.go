package testdata

import "testing"

func TestDoctor_SayHello(t *testing.T) {
	type fields struct {
		Person      *Person
		ID          string
		numPatients int
		string      string
	}
	tests := []struct {
		name   string
		fields fields
		arg    *Person
		want   string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		d := &Doctor{
			Person:      tt.fields.Person,
			ID:          tt.fields.ID,
			numPatients: tt.fields.numPatients,
			string:      tt.fields.string,
		}
		if got := d.SayHello(tt.arg); got != tt.want {
			t.Errorf("%q. Doctor.SayHello() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
