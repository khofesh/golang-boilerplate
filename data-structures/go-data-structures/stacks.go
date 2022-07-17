package main

import "errors"

// Stack
// stack representation
type Stack struct {
	Elements []int
}

// Push
// handles pushing new element on top of
// the stack
func (s *Stack) Push(elem int) {
	s.Elements = append(s.Elements, elem)
}

// Pop
// removes the top element from the stack or
// returns an error if the stack is empty
func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("stack is empty")
	} else {
		lastElementIndex := len(s.Elements) - 1

		var lastElement int
		lastElement, s.Elements = s.Elements[lastElementIndex], s.Elements[:lastElementIndex]

		return lastElement, nil
	}
}

// Peek
// returns the top element from the stack
// without updating the stack
func (s *Stack) Peek() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("stack is empty")
	} else {
		return s.Elements[len(s.Elements)-1], nil
	}
}

// Length
// return the size of the stack
func (s *Stack) Length() int {
	return len(s.Elements)
}

// IsEmpty
// returns true if length of stack is 0
func (s *Stack) IsEmpty() bool {
	return len(s.Elements) == 0
}
