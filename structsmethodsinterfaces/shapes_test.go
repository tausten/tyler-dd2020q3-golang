package structsmethodsinterfaces

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPerimiter(t *testing.T) {
	// Given
	rect := Rectangle{10.0, 10.0}

	// When
	result := Perimiter(rect)

	// Then
	assert.Equal(t, 40.0, result)
}

func TestArea(t *testing.T) {

	checkArea := func(t *testing.T, shape Shape, expected float64) {
		t.Helper()

		// When
		result := shape.Area()

		// Then
		assert.Equal(t, expected, result)
	}

	t.Run("rectangles", func(t *testing.T) {
		// Given
		rect := Rectangle{12.0, 6.0}

		// When + Then
		checkArea(t, rect, 72.00)
	})

	t.Run("circles", func(t *testing.T) {
		// Given
		circle := Circle{10}

		// When + Then
		checkArea(t, circle, 314.1592653589793)
	})
}
