package testdata

// Generic function with 'any' constraint
func GenericAny[T any](val T) T {
	return val
}
