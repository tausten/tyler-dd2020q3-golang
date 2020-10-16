package contextfun

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type SpyStore struct {
	response  string
	cancelled bool
}

func (s *SpyStore) Fetch() string {
	time.Sleep(100 * time.Millisecond)
	return s.response
}

func (s *SpyStore) Cancel() {
	s.cancelled = true
}

func TestHandler(t *testing.T) {
	data := "hello, world"

	t.Run("happy path", func(t *testing.T) {
		// Given
		store := &SpyStore{response: data}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		// When
		svr.ServeHTTP(response, request)

		// Then
		assert.Equal(t, data, response.Body.String())
		assert.False(t, store.cancelled, "it should not have cancelled the store")
	})

	t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
		// Given
		store := &SpyStore{response: data}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)

		request = request.WithContext(cancellingCtx)
		response := httptest.NewRecorder()

		// When
		svr.ServeHTTP(response, request)

		// Then
		assert.True(t, store.cancelled, "store was not told to cancel")
	})
}
