package stack

import (
	"go-examples/linkedlist"
)

type Stack[T comparable] struct {
	list *linkedlist.DoublyLinkedList[T]
}

func NewStack[T comparable](list *linkedlist.DoublyLinkedList[T]) *Stack[T] {
	return &Stack[T]{
		list: list,
	}
}

// Size
// O(n)
func (s *Stack[T]) Size() int {
	return s.list.Size()
}

// IsEmpty
// O(n)
func (s *Stack[T]) IsEmpty() bool {
	return s.list.Size() == 0
}

// ToArray
// O(n)
func (s *Stack[T]) ToArray() []T {
	return s.list.ToArray()
}

// Push
// O(1)
func (s *Stack[T]) Push(value T) {
	s.list.Append(value)
}

// Peek
// O(1)
func (s *Stack[T]) Peek() interface{} {
	node := s.list.GetLast()

	if node == nil {
		return nil
	}

	return node.Value
}

// Pop
// O(1)
func (s *Stack[T]) Pop() interface{} {
	node := s.list.DeleteLast()

	if node == nil {
		return nil
	}

	return node.Value
}
