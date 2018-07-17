package testdata

import "testing"

func TestPerson_SayHello(t *testing.T) {
	type fields struct {
		FirstName string
		LastName  string
		Age       int
		Gender    string
		Siblings  []*Person
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
		p := &Person{
			FirstName: tt.fields.FirstName,
			LastName:  tt.fields.LastName,
			Age:       tt.fields.Age,
			Gender:    tt.fields.Gender,
			Siblings:  tt.fields.Siblings,
		}
		if got := p.SayHello(tt.arg); got != tt.want {
			t.Errorf("%q. Person.SayHello() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
