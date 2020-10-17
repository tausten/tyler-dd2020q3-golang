package maths

import (
	"math"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSecondHandAtMidnight(t *testing.T) {
	// Given
	tm := time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC)
	want := Point{X: 150, Y: 150 - 90}

	// When
	got := SecondHand(tm)

	// Then
	assert.Equal(t, want, got)
}

func TestSecondHandAt30Seconds(t *testing.T) {
	// Given
	tm := time.Date(1337, time.January, 1, 0, 0, 30, 0, time.UTC)
	want := Point{X: 150, Y: 150 + 90}

	// When
	got := SecondHand(tm)

	// Then
	assert.Equal(t, want, got)
}

func TestSecondsInRadians(t *testing.T) {
	// Given
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{simpleTime(0, 0, 30), math.Pi},
		{simpleTime(0, 0, 0), 0},
		{simpleTime(0, 0, 45), (math.Pi / 2) * 3},
		{simpleTime(0, 0, 7), (math.Pi / 30) * 7},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			// When
			got := secondsInRadians(c.time)

			// Then
			assert.Equal(t, c.angle, got)
		})
	}
}

func simpleTime(hours, minutes, seconds int) time.Time {
	return time.Date(312, time.October, 28, hours, minutes, seconds, 0, time.UTC)
}

func testName(t time.Time) string {
	return t.Format("15:04:05")
}

func TestSecondHandVector(t *testing.T) {
	// Given
	cases := []struct {
		time  time.Time
		point Point
	}{
		{simpleTime(0, 0, 30), Point{0, -1}},
		{simpleTime(0, 0, 45), Point{-1, 0}},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			// When
			got := secondHandPoint(c.time)

			// Then
			AssertPointsEquivalent(t, c.point, got)
		})
	}
}

func AssertPointsEquivalent(t *testing.T, pWant, pGot Point) {
	const equalityThreshold = 1e-7
	assert.InDelta(t, pWant.X, pGot.X, equalityThreshold, "X-coordinates differ")
	assert.InDelta(t, pWant.Y, pGot.Y, equalityThreshold, "Y-coordinates differ")
}
