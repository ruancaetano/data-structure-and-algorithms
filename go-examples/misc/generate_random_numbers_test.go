package misc_test

import (
	"github.com/stretchr/testify/assert"
	"go-examples/misc"
	"testing"
)

func TestGenerateRandomNumbers(t *testing.T) {
	size := 10
	array := misc.GenerateRandomNumbers(size)
	assert.Len(t, array, size)
}
