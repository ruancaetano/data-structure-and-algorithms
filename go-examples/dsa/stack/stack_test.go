package stack_test

import (
	"testing"

	"github.com/ruancaetano/data-structure-and-algorithms/go-examples/dsa/linkedlist"
	"github.com/ruancaetano/data-structure-and-algorithms/go-examples/dsa/stack"
	"github.com/stretchr/testify/assert"
)

func TestNewStack(t *testing.T) {
	list := linkedlist.NewDoublyLinkedList[int]()
	s := stack.NewStack[int](list)
	assert.NotNil(t, s)
}

func TestStack_Size(t *testing.T) {
	list := linkedlist.NewDoublyLinkedList[int]()

	list.Append(1)
	list.Append(2)

	s := stack.NewStack[int](list)
	assert.Equal(t, 2, s.Size())
}

func TestStack_ToArray(t *testing.T) {
	list := linkedlist.NewDoublyLinkedList[int]()

	list.Append(1)
	list.Append(2)

	s := stack.NewStack[int](list)
	assert.Equal(t, []int{1, 2}, s.ToArray())
}

func TestStack_IsEmpty(t *testing.T) {
	list := linkedlist.NewDoublyLinkedList[int]()

	s := stack.NewStack[int](list)
	assert.True(t, s.IsEmpty())

	list.Append(1)
	list.Append(2)

	assert.False(t, s.IsEmpty())
}

func TestStack_Push(t *testing.T) {
	list := linkedlist.NewDoublyLinkedList[int]()
	s := stack.NewStack[int](list)

	assert.True(t, s.IsEmpty())
	s.Push(1)
	s.Push(2)
	s.Push(3)
	assert.Equal(t, []int{1, 2, 3}, s.ToArray())

}

func TestStack_Peek(t *testing.T) {
	list := linkedlist.NewDoublyLinkedList[int]()
	s := stack.NewStack[int](list)

	assert.True(t, s.IsEmpty())
	assert.Nil(t, s.Peek())

	s.Push(1)
	assert.Equal(t, 1, s.Peek())

	s.Push(2)
	assert.Equal(t, 2, s.Peek())

	s.Push(3)
	assert.Equal(t, 3, s.Peek())

	assert.Equal(t, []int{1, 2, 3}, s.ToArray())
}

func TestStack_Pop(t *testing.T) {
	list := linkedlist.NewDoublyLinkedList[int]()
	s := stack.NewStack[int](list)

	s.Push(1)
	s.Push(2)
	s.Push(3)

	assert.Equal(t, 3, s.Size())
	assert.Equal(t, 3, s.Pop())

	assert.Equal(t, 2, s.Size())
	assert.Equal(t, 2, s.Pop())

	assert.Equal(t, 1, s.Size())
	assert.Equal(t, 1, s.Pop())

	assert.True(t, s.IsEmpty())
	assert.Nil(t, s.Pop())
}
