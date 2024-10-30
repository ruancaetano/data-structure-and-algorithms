package bigonotation_test

import (
	"testing"

	"github.com/ruancaetano/data-structure-and-algorithms/go-examples/dsa/bigonotation"

	"github.com/stretchr/testify/assert"
)

func TestNewBigONotationExample(t *testing.T) {
	arraySize := 10
	result := bigonotation.NewNotationExample(arraySize)

	assert.NotNil(t, result)
	assert.Len(t, result.GetArray(), arraySize)
}

func TestBigONotationExample_GetArray(t *testing.T) {
	arraySize := 10
	example := bigonotation.NewNotationExample(arraySize)

	array := example.GetArray()
	assert.Len(t, array, arraySize)
}

func TestBigONotationExample_SetArray(t *testing.T) {
	arraySize := 10
	example := bigonotation.NewNotationExample(arraySize)
	assert.Len(t, example.GetArray(), arraySize)

	newArray := []int{1, 2, 3, 4, 5}
	example.SetArray(newArray)
	assert.Equal(t, newArray, newArray)
}

func TestBigONotationExample_GetItemAtIndex(t *testing.T) {
	arraySize := 10
	example := bigonotation.NewNotationExample(arraySize)

	index := 5
	result, err := example.GetItemAtIndex(index)
	assert.Nil(t, err)
	assert.Equal(t, example.GetArray()[index], result)

	index = 1000
	result, err = example.GetItemAtIndex(index)
	assert.EqualError(t, err, "invalid index")
	assert.Empty(t, result)
}

func TestBigONotationExample_LinearGetValueIndex(t *testing.T) {
	arraySize := 10
	example := bigonotation.NewNotationExample(arraySize)

	index := arraySize / 2
	value := example.GetArray()[index]
	result, err := example.LinearGetValueIndex(value)
	assert.Nil(t, err)
	assert.Equal(t, index, result)

	result, err = example.LinearGetValueIndex(0)
	assert.EqualError(t, err, "value not found")
	assert.Empty(t, result)
}

func TestBigONotationExample_BinaryGetValueIndex(t *testing.T) {
	arraySize := 10
	example := bigonotation.NewNotationExample(arraySize)
	example.BubbleSort()

	for i := 0; i < arraySize; i++ {
		value := example.GetArray()[i]
		result, err := example.BinaryGetValueIndex(value)
		assert.Nil(t, err)
		assert.Equal(t, i, result)
	}

	result, err := example.BinaryGetValueIndex(0)
	assert.EqualError(t, err, "value not found")
	assert.Empty(t, result)
}

func TestBigONotationExample_BubbleSort(t *testing.T) {
	arraySize := 10
	example := bigonotation.NewNotationExample(arraySize)

	array := []int{5, 4, 3, 2, 1}
	sortedArray := []int{1, 2, 3, 4, 5}

	example.SetArray(array)
	example.BubbleSort()

	assert.Equal(t, sortedArray, example.GetArray())
}
