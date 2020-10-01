package integers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}

func TestAdder(t *testing.T) {
	// Given
	// When
	result := Add(2, 2)

	// Then
	assert.Equal(t, 4, result)
}
