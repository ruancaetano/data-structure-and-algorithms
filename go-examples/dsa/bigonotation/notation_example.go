package bigonotation

import (
	"errors"

	"github.com/ruancaetano/data-structure-and-algorithms/go-examples/misc"
)

type NotationExample struct {
	array []int
}

func NewNotationExample(arraySize int) *NotationExample {
	return &NotationExample{
		misc.GenerateRandomNumbers(arraySize),
	}
}

// GetArray
// returns the complete list
// O(1)
func (ne *NotationExample) GetArray() []int {
	return ne.array
}

// SetArray
// only set array value
// O(1)
func (ne *NotationExample) SetArray(newArray []int) {
	ne.array = newArray
}

// GetItemAtIndex
// returns the item on specified index value, if not found throws error
// O(1)
func (ne *NotationExample) GetItemAtIndex(index int) (int, error) {
	if index >= len(ne.array) {
		return 0, errors.New("invalid index")
	}

	return ne.array[index], nil
}

// LinearGetValueIndex
// returns the index of the specified value, if not found throws error
// O(n)
func (ne *NotationExample) LinearGetValueIndex(value int) (int, error) {
	for i := 0; i < len(ne.array); i++ {
		if ne.array[i] == value {
			return i, nil
		}
	}

	return 0, errors.New("value not found")
}

// BinaryGetValueIndex
// returns the index of the specified value, if not found throws error
// O(log n)
func (ne *NotationExample) BinaryGetValueIndex(value int) (int, error) {
	return ne.recursiveBinarySearch(value, 0, len(ne.array)-1)
}

func (ne *NotationExample) recursiveBinarySearch(value int, lowIndex, highIndex int) (int, error) {
	if lowIndex > highIndex {
		return 0, errors.New("value not found")
	}

	middleIndex := (lowIndex + highIndex) / 2

	if ne.array[middleIndex] == value {
		return middleIndex, nil
	}

	if ne.array[middleIndex] > value {
		return ne.recursiveBinarySearch(value, lowIndex, middleIndex-1)
	}

	return ne.recursiveBinarySearch(value, middleIndex+1, highIndex)
}

// BubbleSort
// sorts array with bubble sort algorithm
// O(nˆ2)
func (ne *NotationExample) BubbleSort() {
	for i := 0; i < len(ne.array)-1; i++ {
		for j := 0; j < len(ne.array)-1; j++ {
			if ne.array[j] > ne.array[j+1] {
				ne.swapArrayValues(j, j+1)
			}
		}
	}
}

func (ne *NotationExample) swapArrayValues(indexA, indexB int) {
	temp := ne.array[indexA]
	ne.array[indexA] = ne.array[indexB]
	ne.array[indexB] = temp
}
