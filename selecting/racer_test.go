package selecting

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRacer(t *testing.T) {

	t.Run("fastest wins", func(t *testing.T) {
		// Given
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL
		want := fastURL

		// When
		got, err := Racer(slowURL, fastURL)

		// Then
		assert.Equal(t, want, got)
		assert.NoError(t, err)
	})

	t.Run("returns error if server doesn't respond within 10s", func(t *testing.T) {
		// Given
		const timeout = 10 * time.Millisecond
		server := makeDelayedServer(timeout + 10*time.Millisecond)
		defer server.Close()

		// When
		_, err := ConfigurableRacer(server.URL, server.URL, timeout)

		// Then
		assert.Error(t, err)
	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
