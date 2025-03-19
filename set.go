package set

type Set[T comparable] map[T]struct{}

func New[T comparable]() Set[T] {
	return make(Set[T])
}

func (s Set[T]) Copy() Set[T] {
	result := New[T]()
	for value := range s {
		result.Add(value)
	}
	return result
}

func (s Set[T]) Add(value T) {
	s[value] = struct{}{}
}

func (s Set[T]) Remove(value T) {
	delete(s, value)
}

func (s Set[T]) Contains(value T) bool {
	_, ok := s[value]
	return ok
}

func (s Set[T]) ToSlice() []T {
	result := make([]T, len(s))
	i := 0
	for value := range s {
		result[i] = value
		i++
	}
	return result
}

func (s Set[T]) Union(other Set[T]) Set[T] {
	result := New[T]()
	for value := range s {
		result.Add(value)
	}
	for value := range other {
		result.Add(value)
	}
	return result
}

func (s Set[T]) Difference(other Set[T]) Set[T] {
	result := New[T]()
	for value := range s {
		if !other.Contains(value) {
			result.Add(value)
		}
	}
	return result
}
