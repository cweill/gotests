package testdata

import "testing"

func TestPersonSayHello(t *testing.T) {
	tests := []struct {
		// Test description.
		name string
		// Receiver fields.
		rFirstName string
		rLastName  string
		rAge       int
		rGender    string
		rSiblings  []*Person
		// Parameters.
		r *Person
		// Expected results.
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		p := &Person{
			FirstName: tt.rFirstName,
			LastName:  tt.rLastName,
			Age:       tt.rAge,
			Gender:    tt.rGender,
			Siblings:  tt.rSiblings,
		}
		if got := p.SayHello(tt.r); got != tt.want {
			t.Errorf("%q. Person.SayHello() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
