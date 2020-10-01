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

func TestSumAll(t *testing.T) {
	// Given
	// When
	result := SumAll([]int{1, 2}, []int{0, 9})

	// Then
	assert.ElementsMatch(t, []int{3, 9}, result)
}

func TestSumAllTails(t *testing.T) {
	// Given
	// When
	result := SumAllTails([]int{1, 2}, []int{0, 9})

	// Then
	assert.ElementsMatch(t, []int{2, 9}, result)
}

func TestSumAllTailsEmptySlice(t *testing.T) {
	// Given
	// When
	result := SumAllTails([]int{}, []int{0, 9})

	// Then
	assert.ElementsMatch(t, []int{0, 9}, result)
}
