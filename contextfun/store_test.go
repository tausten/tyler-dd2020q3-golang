package contextfun

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type SpyStore struct {
	response string
	t        *testing.T
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)

	go func() {
		var result string
		for _, c := range s.response {
			select {
			case <-ctx.Done():
				s.t.Log("spy store got cancelled")
			default:
				time.Sleep(10 * time.Millisecond)
				result += string(c)
			}
		}
		data <- result
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-data:
		return res, nil
	}
}

func NewSpyStore(t *testing.T, data string) *SpyStore {
	return &SpyStore{t: t, response: data}
}

type SpyResponseWriter struct {
	written bool
}

func (s *SpyResponseWriter) Header() http.Header {
	s.written = true
	return nil
}

func (s *SpyResponseWriter) Write([]byte) (int, error) {
	s.written = true
	return 0, errors.New("not implemented")
}

func (s *SpyResponseWriter) WriteHeader(statusCode int) {
	s.written = true
}

func TestHandler(t *testing.T) {
	data := "hello, world"

	t.Run("happy path", func(t *testing.T) {
		// Given
		store := NewSpyStore(t, data)
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		// When
		svr.ServeHTTP(response, request)

		// Then
		assert.Equal(t, data, response.Body.String())
		//		assert.False(t, store.cancelled, "it should not have cancelled the store")
	})

	t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
		// Given
		store := NewSpyStore(t, data)
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)

		request = request.WithContext(cancellingCtx)
		response := &SpyResponseWriter{}

		// When
		svr.ServeHTTP(response, request)

		// Then
		assert.False(t, response.written, "response was written eventhough cancelled")
	})
}
