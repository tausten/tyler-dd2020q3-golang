package reflection

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func walk(x interface{}, fn func(input string)) {
	val := reflect.ValueOf(x)
	field := val.Field(0)

	fn(field.String())
}

func TestWalk(t *testing.T) {
	// Given
	const expected = "Chris"
	var got []string

	x := struct {
		Name string
	}{expected}

	// When
	walk(x, func(input string) {
		got = append(got, input)
	})

	// Then
	assert.Equal(t, 1, len(got))
	assert.Equal(t, expected, got[0])
}
