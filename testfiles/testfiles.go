package testfiles

type Bar struct{}

type Bazzar interface {
	Baz() string
}

type baz struct{}

func (b *baz) Baz() string { return "" }
