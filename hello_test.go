package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
	// Given
	// When
	result := Hello()

	// Then
	assert.Equal(t, "Hello, world", result)
}
