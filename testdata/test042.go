package testdata

// Generic type with method (from Issue #165)
type Set[T comparable] struct {
	m map[T]struct{}
}

// Add adds a value to the set
func (s *Set[T]) Add(v T) {
	if s.m == nil {
		s.m = make(map[T]struct{})
	}
	s.m[v] = struct{}{}
}

// Has checks if a value is in the set
func (s *Set[T]) Has(v T) bool {
	_, ok := s.m[v]
	return ok
}
