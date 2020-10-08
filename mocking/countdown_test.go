package mocking

import (
	"bytes"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type CountdownOperationsSpy struct {
	Calls []string
}

const sleep = "sleep"
const write = "write"

func (s *CountdownOperationsSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *CountdownOperationsSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func TestCountdown(t *testing.T) {

	t.Run("output content", func(t *testing.T) {
		// Given
		spy := &CountdownOperationsSpy{}
		buffer := &bytes.Buffer{}

		const expected = `3
2
1
Go!`

		// When
		Countdown(spy, buffer)

		// Then
		assert.Equal(t, expected, buffer.String())
	})

	t.Run("sleep before each print", func(t *testing.T) {
		// Given
		spy := &CountdownOperationsSpy{}

		// When
		Countdown(spy, spy)

		// Then
		assert.Equal(t, []string{sleep, write, sleep, write, sleep, write, sleep, write}, spy.Calls)
	})
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func TestConfigurableSleeper(t *testing.T) {
	// Given
	sleepTime := 5 * time.Second
	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}

	// When
	sleeper.Sleep()

	// Then
	assert.Equal(t, sleepTime, spyTime.durationSlept)
}
