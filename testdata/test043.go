package testdata

// Generic function with multiple type parameters
func Pair[T, U any](first T, second U) (T, U) {
	return first, second
}
