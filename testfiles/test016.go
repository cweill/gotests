package test16

type Bazzer interface {
	Baz()
}

type baz struct

func (b *baz) Baz() {}

func (i Interface)Foo16(in Interface) Interface {return &baz{}}
