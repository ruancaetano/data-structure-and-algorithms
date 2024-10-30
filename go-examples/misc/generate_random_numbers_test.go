package misc_test

import (
	"testing"

	"github.com/ruancaetano/data-structure-and-algorithms/go-examples/misc"
	"github.com/stretchr/testify/assert"
)

func TestGenerateRandomNumbers(t *testing.T) {
	size := 10
	array := misc.GenerateRandomNumbers(size)
	assert.Len(t, array, size)
}
