package hashtable_test

import (
	"github.com/stretchr/testify/assert"
	"go-examples/dsa/hashtable"
	"testing"
)

func TestNewHashTable(t *testing.T) {
	cht := hashtable.NewHashTable[string]()
	assert.NotNil(t, cht)
}

func TestHashTable_Insert(t *testing.T) {
	cht := hashtable.NewHashTable[int]()

	item1 := &hashtable.Item[int]{
		Key:   "one",
		Value: 1,
	}

	item2 := &hashtable.Item[int]{
		Key:   "two",
		Value: 2,
	}

	item3 := &hashtable.Item[int]{
		Key:   "owt",
		Value: 3,
	}

	cht.Insert(item1)
	cht.Insert(item2)
	cht.Insert(item3)

	index1 := cht.Hash("one")
	index2 := cht.Hash("two")

	node1 := cht.Table[index1].GetNodeAt(0)
	node2 := cht.Table[index2].GetNodeAt(0)
	node3 := cht.Table[index2].GetNodeAt(1)

	assert.Equal(t, item1, node1.Value)
	assert.Equal(t, item2, node2.Value)
	assert.Equal(t, item3, node3.Value)
}

func TestHashTable_Get(t *testing.T) {
	cht := hashtable.NewHashTable[int]()

	item1 := &hashtable.Item[int]{
		Key:   "one",
		Value: 1,
	}
	cht.Insert(item1)

	item2 := &hashtable.Item[int]{
		Key:   "two",
		Value: 2,
	}
	cht.Insert(item2)

	assert.Equal(t, item1, cht.Get("one"))
	assert.Equal(t, item2, cht.Get("two"))
	assert.Nil(t, cht.Get("three"))
	assert.Nil(t, cht.Get("owt"))
}

func TestHashTable_Delete(t *testing.T) {
	cht := hashtable.NewHashTable[int]()

	item1 := &hashtable.Item[int]{
		Key:   "one",
		Value: 1,
	}
	cht.Insert(item1)

	item2 := &hashtable.Item[int]{
		Key:   "two",
		Value: 2,
	}
	cht.Insert(item2)

	assert.Equal(t, item1, cht.Get("one"))
	assert.Equal(t, item2, cht.Get("two"))

	assert.Equal(t, item1, cht.Delete("one"))
	assert.Equal(t, item2, cht.Delete("two"))

	assert.Nil(t, cht.Delete("one"))
	assert.Nil(t, cht.Delete("two"))
	assert.Nil(t, cht.Delete("three"))
	assert.Nil(t, cht.Delete("owt"))
}
