package dependencyinjection

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGreet(t *testing.T) {
	// Given
	buffer := bytes.Buffer{}

	// When
	Greet(&buffer, "Chris")

	// Then
	assert.Equal(t, "Hello, Chris", buffer.String())
}
