package linkedlist

type SimpleLinkedListNode[T comparable] struct {
	Next  *SimpleLinkedListNode[T]
	Value T
}

type SimpleLinkedList[T comparable] struct {
	Head *SimpleLinkedListNode[T]
}

func NewSimpleLinkedListNode[T comparable](value T) *SimpleLinkedListNode[T] {
	return &SimpleLinkedListNode[T]{
		Value: value,
	}
}

func NewSimpleLinkedList[T comparable]() *SimpleLinkedList[T] {
	return &SimpleLinkedList[T]{}
}

func (sll *SimpleLinkedList[T]) InsertAtStart(value T) *SimpleLinkedListNode[T] {
	newNode := NewSimpleLinkedListNode[T](value)
	if sll.Head == nil {
		sll.Head = newNode
		return newNode
	}

	newNode.Next = sll.Head
	sll.Head = newNode

	return newNode
}

func (sll *SimpleLinkedList[T]) InsertAtEnd(value T) *SimpleLinkedListNode[T] {
	newNode := NewSimpleLinkedListNode[T](value)
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

func (sll *SimpleLinkedList[T]) Size() int {
	size := 0
	listNode := sll.Head
	for listNode != nil {
		size += 1
		listNode = listNode.Next
	}
	return size
}

func (sll *SimpleLinkedList[T]) GetIndex(value T) int {
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

func (sll *SimpleLinkedList[T]) GetNodeAt(desiredIndex int) *SimpleLinkedListNode[T] {
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

func (sll *SimpleLinkedList[T]) DeleteValue(value T) *SimpleLinkedListNode[T] {
	var prevNode *SimpleLinkedListNode[T]
	listNode := sll.Head
	for listNode != nil {
		if value == listNode.Value {
			if prevNode == nil {
				sll.Head = listNode.Next
				return listNode
			}
			prevNode.Next = listNode.Next
			return listNode
		}

		prevNode = listNode
		listNode = listNode.Next
	}

	return nil
}

func (sll *SimpleLinkedList[T]) DeleteValueAt(desiredIndex int) *SimpleLinkedListNode[T] {
	index := 0
	var prevNode *SimpleLinkedListNode[T]
	listNode := sll.Head
	for listNode != nil {
		if index == desiredIndex {
			if prevNode == nil {
				sll.Head = listNode.Next
				return listNode
			}
			prevNode.Next = listNode.Next
			return listNode
		}

		prevNode = listNode
		listNode = listNode.Next
		index += 1
	}

	return nil
}
