package arraysandslices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumSlice(t *testing.T) {
	// Given
	numbers := []int{1, 2, 3, 4, 5}

	// When
	result := Sum(numbers)

	// Then
	assert.Equal(t, 15, result)
}
