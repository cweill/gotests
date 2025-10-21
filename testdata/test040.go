package testdata

// Generic function with 'comparable' constraint
func GenericComparable[T comparable](a, b T) bool {
	return a == b
}
