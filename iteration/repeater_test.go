package iteration

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ExampleRepeat() {
	defaultRepeat := Repeat("B")
	fmt.Println(defaultRepeat)

	explicitRepeat := Repeat("A", 2)
	fmt.Println(explicitRepeat)
	//Output: BBBBB
	//AA
}

func TestRepeat(t *testing.T) {
	// Given
	// When
	result := Repeat("a")

	// Then
	assert.Equal(t, "aaaaa", result)
}

func TestRepeatNumTimes(t *testing.T) {
	// Given
	const times = 3

	// When
	result := Repeat("a", times)

	// Then
	assert.Equal(t, "aaa", result)
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
