package linkedlist_test

import (
	"github.com/stretchr/testify/assert"
	"go-examples/linkedlist"
	"testing"
)

func TestNewDoublyLinkedListNode(t *testing.T) {
	node1 := linkedlist.NewDoublyLinkedListNode[int](1)

	assert.NotNil(t, node1)
	assert.Equal(t, 1, node1.Value)
}

func TestNewDoublyLinkedList(t *testing.T) {
	list := linkedlist.NewDoublyLinkedList[int]()

	assert.NotNil(t, list)
	assert.Nil(t, list.Head)
	assert.Nil(t, list.Tail)
}

func TestNewDoublyLinkedList_Prepend(t *testing.T) {
	list := linkedlist.NewDoublyLinkedList[int]()
	assert.NotNil(t, list)

	node1 := list.Prepend(1)
	assert.NotNil(t, node1)
	assert.Nil(t, node1.Next)
	assert.Nil(t, node1.Prev)
	assert.Equal(t, 1, node1.Value)
	assert.Equal(t, list.Head, node1)
	assert.Equal(t, list.Tail, node1)

	node2 := list.Prepend(2)
	assert.NotNil(t, node2)
	assert.Equal(t, node1, node2.Next)
	assert.Nil(t, node2.Prev)
	assert.Equal(t, node1.Prev, node2)
	assert.Equal(t, 2, node2.Value)
	assert.Equal(t, list.Head, node2)
	assert.Equal(t, list.Tail, node1)
}

func TestNewDoublyLinkedList_Append(t *testing.T) {
	list := linkedlist.NewDoublyLinkedList[int]()
	assert.NotNil(t, list)

	node1 := list.Append(1)
	assert.NotNil(t, node1)
	assert.Nil(t, node1.Next)
	assert.Nil(t, node1.Prev)
	assert.Equal(t, 1, node1.Value)
	assert.Equal(t, node1, list.Head)
	assert.Equal(t, node1, list.Tail)

	node2 := list.Append(2)
	assert.NotNil(t, node2)
	assert.Nil(t, node2.Next)
	assert.Equal(t, node1, node2.Prev)
	assert.Equal(t, node2, node1.Next)
	assert.Equal(t, 2, node2.Value)
	assert.Equal(t, node1, list.Head)
	assert.Equal(t, node2, list.Tail)

	node3 := list.Append(3)
	assert.NotNil(t, node3)
	assert.Nil(t, node3.Next)
	assert.Equal(t, node2, node3.Prev)
	assert.Equal(t, node3, node2.Next)
	assert.Equal(t, 3, node3.Value)
	assert.Equal(t, node1, list.Head)
	assert.Equal(t, node3, list.Tail)
}

func TestNewDoublyLinkedList_Size(t *testing.T) {
	list := linkedlist.NewDoublyLinkedList[int]()
	assert.NotNil(t, list)

	list.Prepend(1)
	list.Append(2)

	assert.Equal(t, 2, list.Size())
}

func TestNewDoublyLinkedList_GetIndex(t *testing.T) {
	list := linkedlist.NewDoublyLinkedList[int]()
	assert.NotNil(t, list)

	list.Append(1)
	list.Append(2)

	assert.Equal(t, 0, list.GetIndex(1))
	assert.Equal(t, 1, list.GetIndex(2))
	assert.Equal(t, -1, list.GetIndex(3))
}

func TestNewDoublyLinkedList_GetValueAt(t *testing.T) {
	list := linkedlist.NewDoublyLinkedList[int]()
	assert.NotNil(t, list)

	node1 := list.Append(1)
	node2 := list.Append(2)

	assert.Equal(t, node1, list.GetNodeAt(0))
	assert.Equal(t, node2, list.GetNodeAt(1))
	assert.Nil(t, list.GetNodeAt(2))
}

func TestNewDoublyLinkedList_DeleteValue(t *testing.T) {
	list := linkedlist.NewDoublyLinkedList[int]()
	assert.NotNil(t, list)

	node1 := list.Append(1)
	node2 := list.Append(2)

	assert.Equal(t, 2, list.Size())
	assert.Equal(t, node2, list.DeleteValue(2))
	assert.Equal(t, 1, list.Size())
	assert.Equal(t, node1, list.DeleteValue(1))
	assert.Equal(t, 0, list.Size())
	assert.Nil(t, list.DeleteValue(0))
}

func TestNewDoublyLinkedList_DeleteValueAt(t *testing.T) {
	list := linkedlist.NewDoublyLinkedList[int]()
	assert.NotNil(t, list)

	node1 := list.Append(1)
	node2 := list.Append(2)

	assert.Equal(t, 2, list.Size())
	assert.Equal(t, node2, list.DeleteValueAt(1))
	assert.Equal(t, 1, list.Size())
	assert.Equal(t, node1, list.DeleteValueAt(0))
	assert.Equal(t, 0, list.Size())
	assert.Nil(t, list.DeleteValueAt(0))
}

func TestNewDoublyLinkedList_GetFirst(t *testing.T) {
	list := linkedlist.NewDoublyLinkedList[int]()
	assert.NotNil(t, list)

	node1 := list.Prepend(1)
	assert.Equal(t, node1, list.GetFirst())
	node2 := list.Prepend(2)
	assert.Equal(t, node2, list.GetFirst())
}

func TestNewDoublyLinkedList_GetLast(t *testing.T) {
	list := linkedlist.NewDoublyLinkedList[int]()
	assert.NotNil(t, list)

	node1 := list.Append(1)
	assert.Equal(t, node1, list.GetLast())
	node2 := list.Append(2)
	assert.Equal(t, node2, list.GetLast())
}

func TestNewDoublyLinkedList_DeleteFirst(t *testing.T) {
	list := linkedlist.NewDoublyLinkedList[int]()
	assert.NotNil(t, list)

	assert.Nil(t, list.DeleteFirst())

	node1 := list.Append(1)
	node2 := list.Append(2)
	assert.Equal(t, 2, list.Size())

	assert.Equal(t, node1, list.DeleteFirst())
	assert.Equal(t, 1, list.Size())

	assert.Equal(t, node2, list.DeleteFirst())
	assert.Equal(t, 0, list.Size())
}

func TestNewDoublyLinkedList_DeleteLast(t *testing.T) {
	list := linkedlist.NewDoublyLinkedList[int]()
	assert.NotNil(t, list)

	assert.Nil(t, list.DeleteLast())

	node1 := list.Append(1)
	node2 := list.Append(2)

	assert.Equal(t, 2, list.Size())
	assert.Equal(t, node2, list.DeleteLast())
	assert.Equal(t, 1, list.Size())
	assert.Equal(t, node1, list.DeleteLast())
	assert.Equal(t, 0, list.Size())
}

func TestNewDoublyLinkedList_ToArray(t *testing.T) {
	list := linkedlist.NewDoublyLinkedList[int]()
	assert.NotNil(t, list)

	list.Append(1)
	list.Append(2)
	list.Append(3)

	assert.Equal(t, []int{1, 2, 3}, list.ToArray())
}
