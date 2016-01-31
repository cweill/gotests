package testfiles

type Bar struct{}

type Bazzar interface {
	Baz() string
}

type baz struct{}

func (b *baz) Baz() string { return "" }

type Celsius float64

type Fahrenheit float64

type Person struct {
	Name   string
	Age    int
	Gender string
}
