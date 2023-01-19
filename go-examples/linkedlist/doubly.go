package linkedlist

type DoublyLinkedListNode[T comparable] struct {
	Next  *DoublyLinkedListNode[T]
	Prev  *DoublyLinkedListNode[T]
	Value T
}

type DoublyLinkedList[T comparable] struct {
	Head *DoublyLinkedListNode[T]
	Tail *DoublyLinkedListNode[T]
}

func NewDoublyLinkedListNode[T comparable](value T) *DoublyLinkedListNode[T] {
	return &DoublyLinkedListNode[T]{
		Value: value,
	}
}

func NewDoublyLinkedList[T comparable]() *DoublyLinkedList[T] {
	return &DoublyLinkedList[T]{}
}

// Prepend
// O(1)
func (sll *DoublyLinkedList[T]) Prepend(value T) *DoublyLinkedListNode[T] {
	newNode := NewDoublyLinkedListNode[T](value)
	if sll.Head == nil {
		sll.Head = newNode
		sll.Tail = newNode
		return newNode
	}

	sll.Head.Prev = newNode
	newNode.Next = sll.Head
	newNode.Prev = nil
	sll.Head = newNode

	return newNode
}

// Append
// O(1)
func (sll *DoublyLinkedList[T]) Append(value T) *DoublyLinkedListNode[T] {
	newNode := NewDoublyLinkedListNode[T](value)
	if sll.Head == nil {
		sll.Head = newNode
		sll.Tail = newNode
		return newNode
	}

	newNode.Prev = sll.Tail
	sll.Tail.Next = newNode
	sll.Tail = newNode

	return newNode
}

// Size
// O(n)
func (sll *DoublyLinkedList[T]) Size() int {
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
func (sll *DoublyLinkedList[T]) GetIndex(value T) int {
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
func (sll *DoublyLinkedList[T]) GetNodeAt(desiredIndex int) *DoublyLinkedListNode[T] {
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
func (sll *DoublyLinkedList[T]) DeleteValue(value T) *DoublyLinkedListNode[T] {
	var nextNode *DoublyLinkedListNode[T]
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
func (sll *DoublyLinkedList[T]) DeleteValueAt(desiredIndex int) *DoublyLinkedListNode[T] {
	index := 0
	var nextNode *DoublyLinkedListNode[T]
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
