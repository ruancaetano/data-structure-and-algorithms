package binary_search_tree

import "golang.org/x/exp/constraints"

type BinarySearchTreeNode[T constraints.Ordered] struct {
	Value T
	Left  *BinarySearchTreeNode[T]
	Right *BinarySearchTreeNode[T]
}

type BinarySearchTree[T constraints.Ordered] struct {
	Root *BinarySearchTreeNode[T]
}

func NewBinarySearchTreeNode[T constraints.Ordered](value T) *BinarySearchTreeNode[T] {
	return &BinarySearchTreeNode[T]{
		Value: value,
	}
}

func NewBinarySearchTree[T constraints.Ordered]() *BinarySearchTree[T] {
	return &BinarySearchTree[T]{}
}

func (t *BinarySearchTree[T]) Insert(value T) *BinarySearchTreeNode[T] {
	node := NewBinarySearchTreeNode[T](value)

	t.Root = t.insertNode(t.Root, node)

	return node
}

func (t *BinarySearchTree[T]) Find(value T) *BinarySearchTreeNode[T] {
	return t.findNode(t.Root, value)
}

func (t *BinarySearchTree[T]) FindParent(value T) *BinarySearchTreeNode[T] {
	return t.findNodeParent(t.Root, value, nil)
}

func (t *BinarySearchTree[T]) CountNodes() int {
	return t.countNodes(t.Root)
}

func (t *BinarySearchTree[T]) insertNode(ref, node *BinarySearchTreeNode[T]) *BinarySearchTreeNode[T] {
	if ref == nil {
		return node
	}

	if node.Value <= ref.Value {
		ref.Left = t.insertNode(ref.Left, node)
		return ref
	}
	ref.Right = t.insertNode(ref.Right, node)
	return ref
}

func (t *BinarySearchTree[T]) findNode(ref *BinarySearchTreeNode[T], value T) *BinarySearchTreeNode[T] {
	if ref == nil {
		return nil
	}

	if ref.Value == value {
		return ref
	}

	if value <= ref.Value {
		return t.findNode(ref.Left, value)
	}
	return t.findNode(ref.Right, value)
}

func (t *BinarySearchTree[T]) findNodeParent(ref *BinarySearchTreeNode[T], value T, parent *BinarySearchTreeNode[T]) *BinarySearchTreeNode[T] {
	if ref == nil {
		return nil
	}

	if ref.Value == value {
		return parent
	}

	if value < ref.Value {
		return t.findNodeParent(ref.Left, value, ref)
	}
	return t.findNodeParent(ref.Right, value, ref)
}

func (t *BinarySearchTree[T]) countNodes(ref *BinarySearchTreeNode[T]) int {
	if ref == nil {
		return 0
	}
	return t.countNodes(ref.Left) + t.countNodes(ref.Right) + 1
}
