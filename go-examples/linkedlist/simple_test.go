package linkedlist_test

import (
	"github.com/stretchr/testify/assert"
	"go-examples/linkedlist"
	"testing"
)

func TestNewSimpleLinkedListNode(t *testing.T) {
	node1 := linkedlist.NewSimpleLinkedListNode[int](1)

	assert.NotNil(t, node1)
	assert.Equal(t, 1, node1.Value)
}

func TestNewSimpleLinkedList(t *testing.T) {
	list := linkedlist.NewSimpleLinkedList[int]()

	assert.NotNil(t, list)
	assert.Nil(t, list.Head)
}

func TestNewSimpleLinkedList_InsertAtStart(t *testing.T) {
	list := linkedlist.NewSimpleLinkedList[int]()
	assert.NotNil(t, list)

	node1 := list.InsertAtStart(1)
	assert.NotNil(t, node1)
	assert.Nil(t, node1.Next)
	assert.Equal(t, 1, node1.Value)
	assert.Equal(t, list.Head, node1)

	node2 := list.InsertAtStart(2)
	assert.NotNil(t, node2)
	assert.Equal(t, node1, node2.Next)
	assert.Equal(t, 2, node2.Value)
	assert.Equal(t, list.Head, node2)
}

func TestNewSimpleLinkedList_InsertAtEnd(t *testing.T) {
	list := linkedlist.NewSimpleLinkedList[int]()
	assert.NotNil(t, list)

	node1 := list.InsertAtEnd(1)
	assert.NotNil(t, node1)
	assert.Nil(t, node1.Next)
	assert.Equal(t, 1, node1.Value)

	node2 := list.InsertAtEnd(2)
	assert.NotNil(t, node2)
	assert.Nil(t, node2.Next)
	assert.Equal(t, node2, node1.Next)
	assert.Equal(t, 2, node2.Value)

	node3 := list.InsertAtEnd(3)
	assert.NotNil(t, node3)
	assert.Nil(t, node3.Next)
	assert.Equal(t, node3, node2.Next)
	assert.Equal(t, 3, node3.Value)
}

func TestNewSimpleLinkedList_Size(t *testing.T) {
	list := linkedlist.NewSimpleLinkedList[int]()
	assert.NotNil(t, list)

	list.InsertAtStart(1)
	list.InsertAtStart(2)

	assert.Equal(t, 2, list.Size())
}

func TestNewSimpleLinkedList_GetIndex(t *testing.T) {
	list := linkedlist.NewSimpleLinkedList[int]()
	assert.NotNil(t, list)

	list.InsertAtEnd(1)
	list.InsertAtEnd(2)

	assert.Equal(t, 0, list.GetIndex(1))
	assert.Equal(t, 1, list.GetIndex(2))
	assert.Equal(t, -1, list.GetIndex(3))
}

func TestNewSimpleLinkedList_GetValueAt(t *testing.T) {
	list := linkedlist.NewSimpleLinkedList[int]()
	assert.NotNil(t, list)

	node1 := list.InsertAtEnd(1)
	node2 := list.InsertAtEnd(2)

	assert.Equal(t, node1, list.GetNodeAt(0))
	assert.Equal(t, node2, list.GetNodeAt(1))
	assert.Nil(t, list.GetNodeAt(2))
}

func TestNewSimpleLinkedList_DeleteValue(t *testing.T) {
	list := linkedlist.NewSimpleLinkedList[int]()
	assert.NotNil(t, list)

	node1 := list.InsertAtEnd(1)
	node2 := list.InsertAtEnd(2)

	assert.Equal(t, 2, list.Size())
	assert.Equal(t, node2, list.DeleteValue(2))
	assert.Equal(t, 1, list.Size())
	assert.Equal(t, node1, list.DeleteValue(1))
	assert.Equal(t, 0, list.Size())
	assert.Nil(t, list.DeleteValue(0))
}

func TestNewSimpleLinkedList_DeleteValueAt(t *testing.T) {
	list := linkedlist.NewSimpleLinkedList[int]()
	assert.NotNil(t, list)

	node1 := list.InsertAtEnd(1)
	node2 := list.InsertAtEnd(2)

	assert.Equal(t, 2, list.Size())
	assert.Equal(t, node2, list.DeleteValueAt(1))
	assert.Equal(t, 1, list.Size())
	assert.Equal(t, node1, list.DeleteValueAt(0))
	assert.Equal(t, 0, list.Size())
	assert.Nil(t, list.DeleteValueAt(0))
}
