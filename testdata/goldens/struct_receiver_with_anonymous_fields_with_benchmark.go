package testdata

import "testing"

func TestDoctor_SayHello(t *testing.T) {
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

func BenchmarkDoctor_SayHello(b *testing.B) {
	type fields struct {
		Person      *Person
		ID          string
		numPatients int
		string      string
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
		d := &Doctor{
			Person:      bb.fields.Person,
			ID:          bb.fields.ID,
			numPatients: bb.fields.numPatients,
			string:      bb.fields.string,
		}
		if got := d.SayHello(tt.args.r); got != bb.want {
			b.Errorf("%q. Doctor.SayHello() = %v, want %v", tt.name, got, bb.want)
		}
	}
}
