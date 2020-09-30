package stuff

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddEmptyReturns0(t *testing.T) {
	// Given
	// When
	result := Add("")

	// Then
	assert.Equal(t, 0, result)
}
