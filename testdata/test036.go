package testdata

func SameName() error {
	return nil
}

func sameName() error {
	return nil
}

func (t *SameTypeName) SameName() error {
	return nil
}

func (t *SameTypeName) sameName() error {
	return nil
}

func (t *sameTypeName) SameName() error {
	return nil
}

func (t *sameTypeName) sameName() error {
	return nil
}
