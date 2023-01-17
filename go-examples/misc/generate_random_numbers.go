package misc

import "math/rand"

func GenerateRandomNumbers(arraySize int) []int {
	var result []int
	for i := 0; i < arraySize; i++ {
		result = append(result, rand.Int())
	}

	return result
}
