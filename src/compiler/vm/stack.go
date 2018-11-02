package vm

import (
	"ape/src/data"
	"fmt"
)

// Stack contains the definition of the VM stack
type Stack struct {
	size    int
	items   []data.Data
	pointer int
}

// NewStack creates a new Stack
func NewStack(size int) *Stack {
	return &Stack{size: size, items: make([]data.Data, size), pointer: 0}
}

// top returns the item on the top of the stack
func (s *Stack) top() data.Data {
	if s.pointer == 0 {
		return nil
	}

	return s.items[s.pointer-1]
}

// push adds an item on the stack
func (s *Stack) push(item data.Data) error {
	if s.pointer >= s.size {
		return fmt.Errorf("stack overflow")
	}

	s.items[s.pointer] = item
	s.pointer++

	return nil
}

// pop takes an item off the stack
func (s *Stack) pop() data.Data {
	item := s.top()
	s.pointer--

	return item
}

// popped returns the last popped item (like top but uses the pointer as an index)
func (s *Stack) popped() data.Data {
	return s.items[s.pointer]
}
