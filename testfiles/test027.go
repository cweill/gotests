package testfiles

type Opener interface {
	Open() error
}

type Book struct{}

func (b *Book) Open() error { return nil }

type door struct{}

func (d *door) Open() error { return nil }
