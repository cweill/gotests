package testdata

type Bar struct{}

type Bazzar interface {
	Baz() string
}

type baz struct{}

func (b *baz) Baz() string { return "" }

type Celsius float64

type Fahrenheit float64

type Person struct {
	FirstName string
	LastName  string
	Age       int
	Gender    string
	Siblings  []*Person
}

type Doctor struct {
	*Person
	ID          string
	numPatients int
	string
}

type name string

type Name struct {
	Name string
}
