package testdata

import "testing"

func TestPersonSayHello(t *testing.T) {
	tests := []struct {
		name      string
		firstName string
		lastName  string
		age       int
		gender    string
		siblings  []*Person
		r         *Person
		want      string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		p := &Person{
			FirstName: tt.firstName,
			LastName:  tt.lastName,
			Age:       tt.age,
			Gender:    tt.gender,
			Siblings:  tt.siblings,
		}
		if got := p.SayHello(tt.r); got != tt.want {
			t.Errorf("%q. Person.SayHello() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
