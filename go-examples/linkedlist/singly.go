package linkedlist

type SinglyLinkedListNode[T comparable] struct {
	Next  *SinglyLinkedListNode[T]
	Value T
}

type SinglyLinkedList[T comparable] struct {
	Head *SinglyLinkedListNode[T]
}

func NewSinglyLinkedListNode[T comparable](value T) *SinglyLinkedListNode[T] {
	return &SinglyLinkedListNode[T]{
		Value: value,
	}
}

func NewSinglyLinkedList[T comparable]() *SinglyLinkedList[T] {
	return &SinglyLinkedList[T]{}
}

// Prepend
// O(1)
func (sll *SinglyLinkedList[T]) Prepend(value T) *SinglyLinkedListNode[T] {
	newNode := NewSinglyLinkedListNode[T](value)
	if sll.Head == nil {
		sll.Head = newNode
		return newNode
	}

	newNode.Next = sll.Head
	sll.Head = newNode

	return newNode
}

// Append
// O(n)
func (sll *SinglyLinkedList[T]) Append(value T) *SinglyLinkedListNode[T] {
	newNode := NewSinglyLinkedListNode[T](value)
	if sll.Head == nil {
		sll.Head = newNode
		return newNode
	}

	listNode := sll.Head
	for {
		if listNode.Next == nil {
			listNode.Next = newNode
			return newNode
		}

		listNode = listNode.Next
	}
}

// Size
// O(n)
func (sll *SinglyLinkedList[T]) Size() int {
	size := 0
	listNode := sll.Head
	for listNode != nil {
		size += 1
		listNode = listNode.Next
	}
	return size
}

// GetIndex
// O(n)
func (sll *SinglyLinkedList[T]) GetIndex(value T) int {
	index := 0
	listNode := sll.Head
	for listNode != nil {
		if value == listNode.Value {
			return index
		}

		index += 1
		listNode = listNode.Next
	}

	return -1
}

// GetNodeAt
// O(n)
func (sll *SinglyLinkedList[T]) GetNodeAt(desiredIndex int) *SinglyLinkedListNode[T] {
	index := 0
	listNode := sll.Head
	for listNode != nil {
		if index == desiredIndex {
			return listNode
		}

		index += 1
		listNode = listNode.Next
	}

	return nil
}

// DeleteValue
// O(n)
func (sll *SinglyLinkedList[T]) DeleteValue(value T) *SinglyLinkedListNode[T] {
	var nextNode *SinglyLinkedListNode[T]
	listNode := sll.Head
	for listNode != nil {
		if value == listNode.Value {
			if nextNode == nil {
				sll.Head = listNode.Next
				return listNode
			}
			nextNode.Next = listNode.Next
			return listNode
		}

		nextNode = listNode
		listNode = listNode.Next
	}

	return nil
}

// DeleteValueAt
// O(n)
func (sll *SinglyLinkedList[T]) DeleteValueAt(desiredIndex int) *SinglyLinkedListNode[T] {
	index := 0
	var nextNode *SinglyLinkedListNode[T]
	listNode := sll.Head
	for listNode != nil {
		if index == desiredIndex {
			if nextNode == nil {
				sll.Head = listNode.Next
				return listNode
			}
			nextNode.Next = listNode.Next
			return listNode
		}

		nextNode = listNode
		listNode = listNode.Next
		index += 1
	}

	return nil
}

func (sll *SinglyLinkedList[T]) ToArray() []T {
	var array []T

	listNode := sll.Head
	for listNode != nil {
		array = append(array, listNode.Value)
		listNode = listNode.Next
	}

	return array
}
