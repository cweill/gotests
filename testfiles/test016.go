package test16

type Bazzar interface {
	Baz() string
}

type baz struct{}

func (b *baz) Baz() string { return "" }

func (i Bazzar) Foo16(in Bazzar) Bazzar { return &baz{} }
