package maps

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {

	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		// Given
		// When
		result, err := dictionary.Search("test")

		// Then
		assert.NoError(t, err)
		assert.Equal(t, "this is just a test", result)
	})

	t.Run("unknown word", func(t *testing.T) {
		// Given
		// When
		_, err := dictionary.Search("unknown")

		// Then
		assert.Equal(t, ErrNotFound, errors.Cause(err))
	})
}
