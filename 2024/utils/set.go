package utils

type Set[T comparable] struct {
	elements map[T]struct{}
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{
		elements: make(map[T]struct{}),
	}
}

func (s *Set[T]) Add(value T) {
	s.elements[value] = struct{}{}
}

func (s *Set[T]) Remove(value T) {
	delete(s.elements, value)
}

func (s *Set[T]) Contains(value T) bool {
	_, exists := s.elements[value]
	return exists
}

func (s *Set[T]) Size() int {
	return len(s.elements)
}

func (s *Set[T]) Union(other *Set[T]) *Set[T] {
	var result = NewSet[T]()
	for element := range s.elements {
		result.Add(element)
	}
	for element := range other.elements {
		result.Add(element)
	}
	return result
}

func (s *Set[T]) Intersection(other *Set[T]) *Set[T] {
	var result = NewSet[T]()
	for element := range s.elements {
		if other.Contains(element) {
			result.Add(element)
		}
	}
	return result
}

func (s *Set[T]) Difference(other *Set[T]) *Set[T] {
	var result = NewSet[T]()
	for element := range s.elements {
		if !other.Contains(element) {
			result.Add(element)
		}
	}
	return result
}

func (s *Set[T]) IsEmpty() bool {
	return len(s.elements) == 0
}

func (s *Set[T]) Clear() {
	s.elements = make(map[T]struct{})
}

func (s *Set[T]) Clone() *Set[T] {
	var clone = NewSet[T]()
	for element := range s.elements {
		clone.Add(element)
	}
	return clone
}

func (s *Set[T]) ForEach(action func(T)) {
	for element := range s.elements {
		action(element)
	}
}
