package stack

import "fmt"

// Stack defines a stack structure.
type Stack struct {
	elements []interface{}
}

// NewStack creates and returns a new Stack.
func NewStack() *Stack {
	return &Stack{elements: []interface{}{}}
}

// Push adds an element to the stack.
func (s *Stack) Push(value interface{}) {
	s.elements = append(s.elements, value)
}

// Pop removes and returns the last element from the stack.
// Returns an error if the stack is empty.
func (s *Stack) Pop() (interface{}, error) {
	if len(s.elements) == 0 {
		return 0, fmt.Errorf("stack is empty")
	}
	value := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return value, nil
}

// Peek returns the top element of the stack without removing it.
func (s *Stack) Peek() (interface{}, error) {
	if len(s.elements) == 0 {
		return 0, fmt.Errorf("stack is empty")
	}
	return s.elements[len(s.elements)-1], nil
}

// IsEmpty checks if the stack is empty.
func (s *Stack) IsEmpty() bool {
	return len(s.elements) == 0
}

func (s *Stack) Len() int {
	return len(s.elements)
}

