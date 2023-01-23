package queue

import (
	"go-examples/dsa/linkedlist"
)

type Queue[T comparable] struct {
	list *linkedlist.DoublyLinkedList[T]
}

func NewQueue[T comparable](list *linkedlist.DoublyLinkedList[T]) *Queue[T] {
	return &Queue[T]{
		list: list,
	}
}

// Size
// O(n)
func (q *Queue[T]) Size() int {
	return q.list.Size()
}

// IsEmpty
// O(1)
func (q *Queue[T]) IsEmpty() bool {
	return q.list.Head == nil
}

// ToArray
// O(n)
func (q *Queue[T]) ToArray() []T {
	return q.list.ToArray()
}

// Enqueue
// O(1)
func (q *Queue[T]) Enqueue(value T) {
	q.list.Append(value)
}

// Peek
// O(1)
func (q *Queue[T]) Peek() any {
	first := q.list.GetFirst()

	if first == nil {
		return nil
	}

	return first.Value
}

// Dequeue
// O(1)
func (q *Queue[T]) Dequeue() any {
	value := q.list.DeleteFirst()

	if value == nil {
		return nil
	}

	return value.Value
}
