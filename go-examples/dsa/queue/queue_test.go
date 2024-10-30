package queue_test

import (
	"testing"

	"github.com/ruancaetano/data-structure-and-algorithms/go-examples/dsa/linkedlist"
	"github.com/ruancaetano/data-structure-and-algorithms/go-examples/dsa/queue"
	"github.com/stretchr/testify/assert"
)

func TestNewQueue(t *testing.T) {
	list := linkedlist.NewDoublyLinkedList[int]()
	q := queue.NewQueue[int](list)
	assert.NotNil(t, q)
}

func TestQueue_Size(t *testing.T) {
	list := linkedlist.NewDoublyLinkedList[int]()

	list.Append(1)
	list.Append(2)

	q := queue.NewQueue[int](list)
	assert.Equal(t, 2, q.Size())
}

func TestQueue_ToArray(t *testing.T) {
	list := linkedlist.NewDoublyLinkedList[int]()

	list.Append(1)
	list.Append(2)

	q := queue.NewQueue[int](list)
	assert.Equal(t, []int{1, 2}, q.ToArray())
}

func TestQueue_IsEmpty(t *testing.T) {
	list := linkedlist.NewDoublyLinkedList[int]()

	q := queue.NewQueue[int](list)
	assert.True(t, q.IsEmpty())

	list.Append(1)
	list.Append(2)

	assert.False(t, q.IsEmpty())
}

func TestQueue_Enqueue(t *testing.T) {
	list := linkedlist.NewDoublyLinkedList[int]()
	q := queue.NewQueue[int](list)

	assert.True(t, q.IsEmpty())
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	assert.Equal(t, []int{1, 2, 3}, q.ToArray())

}

func TestQueue_Peek(t *testing.T) {
	list := linkedlist.NewDoublyLinkedList[int]()
	q := queue.NewQueue[int](list)

	assert.True(t, q.IsEmpty())
	assert.Nil(t, q.Peek())

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	assert.Equal(t, 1, q.Peek())
	list.DeleteFirst()

	assert.Equal(t, 2, q.Peek())
	list.DeleteFirst()

	assert.Equal(t, 3, q.Peek())
	list.DeleteFirst()

	assert.Nil(t, q.Peek())
}

func TestQueue_Dequeue(t *testing.T) {
	list := linkedlist.NewDoublyLinkedList[int]()
	q := queue.NewQueue[int](list)

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	assert.Equal(t, 3, q.Size())
	assert.Equal(t, 1, q.Dequeue())

	assert.Equal(t, 2, q.Size())
	assert.Equal(t, 2, q.Dequeue())

	assert.Equal(t, 1, q.Size())
	assert.Equal(t, 3, q.Dequeue())

	assert.True(t, q.IsEmpty())
	assert.Nil(t, q.Dequeue())
}
