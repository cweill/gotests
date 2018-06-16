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

func BenchmarkPerson_SayHello(b *testing.B) {
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
	benchmarks := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
	// TODO: Add benchmark cases.
	}
	for _, bb := range benchmarks {
		p := &Person{
			FirstName: bb.fields.FirstName,
			LastName:  bb.fields.LastName,
			Age:       bb.fields.Age,
			Gender:    bb.fields.Gender,
			Siblings:  bb.fields.Siblings,
		}
		if got := p.SayHello(tt.args.r); got != bb.want {
			b.Errorf("%q. Person.SayHello() = %v, want %v", tt.name, got, bb.want)
		}
	}
}
