package stack

// Stack is used to implement the last-in first-out.
type Stack[T any] struct {
	top *node[T]
	l   int
}

type node[T any] struct {
	previous *node[T]
	value    T
}

// New is used to create a new instance of the stack.
func New[T any]() *Stack[T] {
	return new(Stack[T])
}

// Push is used to push an element to the stack.
func (s *Stack[T]) Push(v T) {
	n := &node[T]{
		previous: s.top,
		value:    v,
	}
	s.top = n
	s.l += 1
}

// Pop is used to remove the top-most element from the stack.
// It returns the value of the removed element.
//
// If there are no elements in the stack, then the default value of the type is returned and a boolean
// value false stating that no elements were popped.
func (s *Stack[T]) Pop() (a T, b bool) {
	if s.top == nil {
		return
	}
	n := s.top
	s.top = n.previous
	s.l--
	return n.value, true
}

// Peek is used to access the top-most element from the stack.
//
// If there are no elements in the stack, then the default value of the type is returned and a boolean
// value false stating that no elements were popped.
func (s *Stack[T]) Peek() (a T, b bool) {
	if s.top == nil {
		return
	}
	return s.top.value, true
}

// Length is used to return the number of elements in the stack.
func (s *Stack[T]) Length() int {
	return s.l
}
