package extratests

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompareEquals(t *testing.T) {
	// Given
	const thing = "a"

	// When
	result := strings.Compare(thing, thing)

	// Then
	assert.Equal(t, 0, result)
}

func TestCompareFirstGreater(t *testing.T) {
	// Given
	// When
	result := strings.Compare("bbb", "aa")

	// Then
	assert.Equal(t, 1, result)
}

func TestCompareFirstLesser(t *testing.T) {
	// Given
	// When
	result := strings.Compare("aa", "b")

	// Then
	assert.Equal(t, -1, result)
}

func TestFieldsFunc(t *testing.T) {
	// Given
	f := func(c rune) bool {
		return '-' == c
	}

	// When
	result := strings.FieldsFunc("this is fun-what do double dashes--do?", f)

	// Then
	assert.Equal(t, "this is fun", result[0])
	assert.Equal(t, "what do double dashes", result[1])
	assert.Equal(t, "do?", result[2])
}
