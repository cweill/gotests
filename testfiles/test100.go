package testfiles

import "errors"

func Foo100(strs []string) ([]*Bar, error) { return nil, nil }

func (b *Bar) Bar100(i interface{}) error {
	if i == nil {
		return errors.New("i is nil")
	}
	return nil
}

func baz100(f *float64) float64 { return *f }
