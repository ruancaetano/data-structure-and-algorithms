package linkedlist_test

import (
	"testing"

	"github.com/ruancaetano/data-structure-and-algorithms/go-examples/dsa/linkedlist"
	"github.com/stretchr/testify/assert"
)

func TestNewSinglyLinkedListNode(t *testing.T) {
	node1 := linkedlist.NewSinglyLinkedListNode[int](1)

	assert.NotNil(t, node1)
	assert.Equal(t, 1, node1.Value)
}

func TestNewSinglyLinkedList(t *testing.T) {
	list := linkedlist.NewSinglyLinkedList[int]()

	assert.NotNil(t, list)
	assert.Nil(t, list.Head)
}

func TestNewSinglyLinkedList_Prepend(t *testing.T) {
	list := linkedlist.NewSinglyLinkedList[int]()
	assert.NotNil(t, list)

	node1 := list.Prepend(1)
	assert.NotNil(t, node1)
	assert.Nil(t, node1.Next)
	assert.Equal(t, 1, node1.Value)
	assert.Equal(t, list.Head, node1)

	node2 := list.Prepend(2)
	assert.NotNil(t, node2)
	assert.Equal(t, node1, node2.Next)
	assert.Equal(t, 2, node2.Value)
	assert.Equal(t, list.Head, node2)
}

func TestNewSinglyLinkedList_Append(t *testing.T) {
	list := linkedlist.NewSinglyLinkedList[int]()
	assert.NotNil(t, list)

	node1 := list.Append(1)
	assert.NotNil(t, node1)
	assert.Nil(t, node1.Next)
	assert.Equal(t, 1, node1.Value)

	node2 := list.Append(2)
	assert.NotNil(t, node2)
	assert.Nil(t, node2.Next)
	assert.Equal(t, node2, node1.Next)
	assert.Equal(t, 2, node2.Value)

	node3 := list.Append(3)
	assert.NotNil(t, node3)
	assert.Nil(t, node3.Next)
	assert.Equal(t, node3, node2.Next)
	assert.Equal(t, 3, node3.Value)
}

func TestNewSinglyLinkedList_Size(t *testing.T) {
	list := linkedlist.NewSinglyLinkedList[int]()
	assert.NotNil(t, list)

	list.Prepend(1)
	list.Prepend(2)

	assert.Equal(t, 2, list.Size())
}

func TestNewSinglyLinkedList_GetIndex(t *testing.T) {
	list := linkedlist.NewSinglyLinkedList[int]()
	assert.NotNil(t, list)

	list.Append(1)
	list.Append(2)

	assert.Equal(t, 0, list.GetIndex(1))
	assert.Equal(t, 1, list.GetIndex(2))
	assert.Equal(t, -1, list.GetIndex(3))
}

func TestNewSinglyLinkedList_GetValueAt(t *testing.T) {
	list := linkedlist.NewSinglyLinkedList[int]()
	assert.NotNil(t, list)

	node1 := list.Append(1)
	node2 := list.Append(2)

	assert.Equal(t, node1, list.GetNodeAt(0))
	assert.Equal(t, node2, list.GetNodeAt(1))
	assert.Nil(t, list.GetNodeAt(2))
}

func TestNewSinglyLinkedList_DeleteValue(t *testing.T) {
	list := linkedlist.NewSinglyLinkedList[int]()
	assert.NotNil(t, list)

	node1 := list.Append(1)
	node2 := list.Append(2)

	assert.Equal(t, 2, list.Size())
	assert.Equal(t, node2, list.DeleteValue(2))
	assert.Equal(t, 1, list.Size())
	assert.Equal(t, node1, list.DeleteValue(1))
	assert.Equal(t, 0, list.Size())
	assert.Nil(t, list.DeleteValue(0))
	assert.Nil(t, list.Head)
}

func TestNewSinglyLinkedList_DeleteValueAt(t *testing.T) {
	list := linkedlist.NewSinglyLinkedList[int]()
	assert.NotNil(t, list)

	node1 := list.Append(1)
	node2 := list.Append(2)

	assert.Equal(t, 2, list.Size())
	assert.Equal(t, node2, list.DeleteValueAt(1))
	assert.Equal(t, 1, list.Size())
	assert.Equal(t, node1, list.DeleteValueAt(0))
	assert.Equal(t, 0, list.Size())
	assert.Nil(t, list.DeleteValueAt(0))
	assert.Nil(t, list.Head)
}

func TestNewSinglyLinkedList_GetFirst(t *testing.T) {
	list := linkedlist.NewSinglyLinkedList[int]()
	assert.NotNil(t, list)

	node1 := list.Prepend(1)
	assert.Equal(t, node1, list.GetFirst())
	node2 := list.Prepend(2)
	assert.Equal(t, node2, list.GetFirst())
}

func TestNewSinglyLinkedList_GetLast(t *testing.T) {
	list := linkedlist.NewSinglyLinkedList[int]()
	assert.NotNil(t, list)

	node1 := list.Append(1)
	assert.Equal(t, node1, list.GetLast())
	node2 := list.Append(2)
	assert.Equal(t, node2, list.GetLast())
}

func TestNewSinglyLinkedList_DeleteFirst(t *testing.T) {
	list := linkedlist.NewSinglyLinkedList[int]()
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

func TestNewSinglyLinkedList_DeleteLast(t *testing.T) {
	list := linkedlist.NewSinglyLinkedList[int]()
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

func TestNewSinglyLinkedList_ToArray(t *testing.T) {
	list := linkedlist.NewSinglyLinkedList[int]()
	assert.NotNil(t, list)

	list.Append(1)
	list.Append(2)
	list.Append(3)

	assert.Equal(t, []int{1, 2, 3}, list.ToArray())
}
