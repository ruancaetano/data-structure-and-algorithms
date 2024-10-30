package misc

func MaxInt(numbers ...int) int {
	max := 0
	for i, number := range numbers {
		if i == 0 {
			max = number
			continue
		}

		if number > max {
			max = number
		}
	}

	return max
}
