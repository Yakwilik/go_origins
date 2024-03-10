package utils

import "golang.org/x/exp/slices"

type Stack[T comparable] struct {
	container []T
}

func (s *Stack[T]) PushBack(value T) {
	s.container = append(s.container, value)
}

func (s *Stack[T]) Pop() {
	length := len(s.container)
	if length == 0 {
		return
	}
	s.container = s.container[:length-1]
}
func (s *Stack[T]) Empty() bool {
	return s.Size() < 1
}

func (s *Stack[T]) Size() int {
	return len(s.container)
}

func (s *Stack[T]) Top() T {
	return s.container[s.Size()-1]
}

func (s *Stack[T]) Has(value T) bool {
	return slices.Contains(s.container, value)
}

func (s *Stack[T]) GetTopOrDefault() T {
	var value T
	if s.Empty() {
		return value
	} else {
		return s.Top()
	}
}
