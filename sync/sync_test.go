package sync

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCounter(t *testing.T) {

	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		// Given
		counter := NewCounter()

		// When
		counter.Inc()
		counter.Inc()
		counter.Inc()

		// Then
		assert.Equal(t, 3, counter.Value())
	})

	t.Run("it runs safely concurrently", func(t *testing.T) {
		// Given
		wantedCount := 1000
		counter := NewCounter()

		var wg sync.WaitGroup
		wg.Add(wantedCount)

		// When
		for i := 0; i < wantedCount; i++ {
			go func(w *sync.WaitGroup) {
				counter.Inc()
				w.Done()
			}(&wg)
		}
		wg.Wait()

		// Then
		assert.Equal(t, wantedCount, counter.Value())
	})
}
