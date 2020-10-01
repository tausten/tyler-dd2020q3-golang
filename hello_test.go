package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const expectedDefaultHelloWorldResponse = "Hello, world"

func TestHelloWorld(t *testing.T) {
	// Given
	// When
	result := HelloWorld()

	// Then
	assert.Equal(t, expectedDefaultHelloWorldResponse, result)
}

func TestHello(t *testing.T) {
	// Given
	const name = "Chris"
	const expectedResult = "Hello, " + name

	// When
	result := Hello(name)

	// Then
	assert.Equal(t, expectedResult, result)
}

func TestHelloDefaultsToWorldForEmptyString(t *testing.T) {
	// Given
	// When
	result := Hello("")
	// Then
	assert.Equal(t, expectedDefaultHelloWorldResponse, result)
}
