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

func TestAdd(t *testing.T) {
	// Given
	const word = "test"
	const definition = "this is just a test"

	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}

		// When
		addErr := dictionary.Add(word, definition)
		result, err := dictionary.Search(word)

		// Then
		assert.NoError(t, addErr)
		assert.NoError(t, err)
		assert.Equal(t, definition, result)
	})

	t.Run("existing word", func(t *testing.T) {
		dictionary := Dictionary{word: definition}

		// When
		err := dictionary.Add(word, "some new definition")

		// Then
		assert.Equal(t, ErrWordExists, errors.Cause(err))
	})
}

func TestUpdate(t *testing.T) {
	// Given
	const word = "test"
	const initialDefinition = "this is just a test"
	const newDefinition = "new definition"

	t.Run("existing word", func(t *testing.T) {
		dictionary := Dictionary{word: initialDefinition}

		// When
		dictionary.Update(word, newDefinition)
		result, err := dictionary.Search(word)

		// Then
		assert.NoError(t, err)
		assert.Equal(t, newDefinition, result)
	})

	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}

		// When
		err := dictionary.Update(word, newDefinition)

		// Then
		assert.Equal(t, ErrWordDoesNotExist, errors.Cause(err))
	})
}
