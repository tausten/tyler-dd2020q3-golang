package sync

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCounter(t *testing.T) {

	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		// Given
		counter := Counter{}

		// When
		counter.Inc()
		counter.Inc()
		counter.Inc()

		// Then
		assert.Equal(t, 3, counter.Value())
	})
}
