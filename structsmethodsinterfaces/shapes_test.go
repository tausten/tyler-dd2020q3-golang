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
	// Given
	areaTests := []struct {
		name     string
		shape    Shape
		expected float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, expected: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, expected: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, expected: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			// When
			result := tt.shape.Area()

			// Then
			assert.Equal(t, tt.expected, result, "%#v", tt.shape)
		})
	}
}
