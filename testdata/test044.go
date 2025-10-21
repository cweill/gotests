package testdata

// Generic type with multiple type parameters
type Map[K comparable, V any] struct {
	m map[K]V
}

// Set sets a key-value pair
func (m *Map[K, V]) Set(key K, value V) {
	if m.m == nil {
		m.m = make(map[K]V)
	}
	m.m[key] = value
}

// Get retrieves a value by key
func (m *Map[K, V]) Get(key K) (V, bool) {
	v, ok := m.m[key]
	return v, ok
}
