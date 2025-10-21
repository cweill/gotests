package testdata

// Generic function with error return
func FindFirst[T comparable](slice []T, target T) (int, error) {
	for i, v := range slice {
		if v == target {
			return i, nil
		}
	}
	return -1, ErrNotFound
}

// ErrNotFound is returned when an element is not found
var ErrNotFound = error(nil)
