package ex1

import "errors"

type Bar struct {
	Field int
}

func (b *Bar) Foo() (string, error) {
	if b.Field <= 0 {
		return "", errors.New("error")
	}
	return "foo", nil
}
