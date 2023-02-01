package tree_test

import (
	"github.com/stretchr/testify/assert"
	"go-examples/dsa/tree"
	"testing"
)

func TestNewBinarySearchTreeNode(t *testing.T) {
	node := tree.NewBinarySearchTreeNode[int](1)
	assert.NotNil(t, node)
	assert.Equal(t, 1, node.Value)
}

func TestNewBinarySearchTree(t *testing.T) {
	searchTree := tree.NewBinarySearchTree[int]()
	assert.NotNil(t, searchTree)
}

func TestNewBinarySearchTree_Insert(t *testing.T) {
	searchTree := tree.NewBinarySearchTree[int]()

	node1 := searchTree.Insert(5)
	assert.Equal(t, 5, node1.Value)

	node2 := searchTree.Insert(1)
	assert.Equal(t, 1, node2.Value)
	assert.Equal(t, node2, node1.Left)

	node3 := searchTree.Insert(10)
	assert.Equal(t, 10, node3.Value)
	assert.Equal(t, node3, node1.Right)
}

func TestNewBinarySearchTree_Find(t *testing.T) {
	searchTree := tree.NewBinarySearchTree[int]()

	node1 := searchTree.Insert(5)
	node2 := searchTree.Insert(1)
	node3 := searchTree.Insert(10)

	found1 := searchTree.Find(5)
	assert.Equal(t, node1, found1)

	found2 := searchTree.Find(1)
	assert.Equal(t, node2, found2)

	found3 := searchTree.Find(10)
	assert.Equal(t, node3, found3)

	assert.Nil(t, searchTree.Find(1000))
}

func TestNewBinarySearchTree_FindParent(t *testing.T) {
	searchTree := tree.NewBinarySearchTree[int]()

	node1 := searchTree.Insert(5)
	searchTree.Insert(1)
	searchTree.Insert(10)

	found1 := searchTree.FindParent(5)
	assert.Nil(t, found1)

	found2 := searchTree.FindParent(1)
	assert.Equal(t, node1, found2)

	found3 := searchTree.FindParent(10)
	assert.Equal(t, node1, found3)

	assert.Nil(t, searchTree.FindParent(1000))
}

func TestNewBinarySearchTree_CountNodes(t *testing.T) {
	searchTree := tree.NewBinarySearchTree[int]()

	assert.Equal(t, 0, searchTree.CountNodes())
	searchTree.Insert(5)
	assert.Equal(t, 1, searchTree.CountNodes())
	searchTree.Insert(1)
	assert.Equal(t, 2, searchTree.CountNodes())
	searchTree.Insert(10)
	assert.Equal(t, 3, searchTree.CountNodes())
}
