package selecting

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRacer(t *testing.T) {
	// Given
	slowURL := "http://www.facebook.com"
	fastURL := "http://www.quii.co.uk"
	want := fastURL

	// When
	got := Racer(slowURL, fastURL)

	// Then
	assert.Equal(t, want, got)
}
