package misc_test

import (
	"go-examples/misc"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxInt(t *testing.T) {
	numbers := []int{10, 2, 3, 4, 5}
	assert.Equal(t, 10, misc.MaxInt(numbers...))

	numbers = []int{500, 100, 1000, 1, 1}
	assert.Equal(t, 1000, misc.MaxInt(numbers...))

	numbers = []int{500, 100, 1000, 1, 2001}
	assert.Equal(t, 2001, misc.MaxInt(numbers...))
}
