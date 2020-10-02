package mocking

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func TestCountdown(t *testing.T) {
	// Given
	sleeper := &SpySleeper{}
	buffer := &bytes.Buffer{}

	const expected = `3
2
1
Go!`

	// When
	Countdown(sleeper, buffer)

	// Then
	assert.Equal(t, expected, buffer.String())
	assert.Equal(t, 4, sleeper.Calls)
}
