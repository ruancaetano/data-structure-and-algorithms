package hashtable

import (
	"log"

	"github.com/ruancaetano/data-structure-and-algorithms/go-examples/dsa/linkedlist"
)

const PrimeBlockSize = 33

type Item[T comparable] struct {
	Key   string
	Value T
}

type Chaining[T comparable] struct {
	*linkedlist.SinglyLinkedList[*Item[T]]
}

// HashTable
// Using chaining conflict resolution
type HashTable[T comparable] struct {
	Table []*Chaining[T]
}

func NewHashTable[T comparable]() *HashTable[T] {
	return &HashTable[T]{
		Table: make([]*Chaining[T], PrimeBlockSize),
	}
}

func (ht *HashTable[T]) Hash(key string) int {
	var index int

	for _, letter := range key {
		index += int(letter)
	}

	index = index % PrimeBlockSize
	log.Printf("Index %d for key %s", index, key)
	return index
}

func (ht *HashTable[T]) Insert(item *Item[T]) {
	index := ht.Hash(item.Key)

	if ht.Table[index] == nil {
		ht.Table[index] = &Chaining[T]{
			linkedlist.NewSinglyLinkedList[*Item[T]](),
		}
	}

	list := ht.Table[index]
	list.Append(item)
}

func (ht *HashTable[T]) Delete(key string) *Item[T] {
	index := ht.Hash(key)

	if ht.Table[index] == nil {
		return nil
	}

	list := ht.Table[index]

	node := list.Head
	for node != nil {
		if node.Value.Key == key {
			list.DeleteValue(node.Value)
			return node.Value
		}

		node = node.Next
	}
	return nil
}

func (ht *HashTable[T]) Get(key string) *Item[T] {
	index := ht.Hash(key)

	if ht.Table[index] == nil {
		return nil
	}

	list := ht.Table[index]

	node := list.Head
	for node != nil {
		if node.Value.Key == key {
			return node.Value
		}

		node = node.Next
	}
	return nil
}
