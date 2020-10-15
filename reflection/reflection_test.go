package reflection

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {
	// Given
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"Struct with one string field",
			struct {
				Name string
			}{"Chris"},
			[]string{"Chris"},
		},
		{
			"Struct with two string fields",
			struct {
				Name string
				City string
			}{"Chris", "London"},
			[]string{"Chris", "London"},
		},
		{
			"Struct with non-string field",
			struct {
				Name string
				Age  int
			}{"Chris", 33},
			[]string{"Chris"},
		},
		{
			"Nested fields",
			Person{"Chris", Profile{33, "London"}},
			[]string{"Chris", "London"},
		},
		{
			"Pointers to things",
			&Person{"Chris", Profile{33, "London"}},
			[]string{"Chris", "London"},
		},
		{
			"Slices",
			[]Profile{
				{33, "London"},
				{34, "Reykjavík"},
			},
			[]string{"London", "Reykjavík"},
		},
		{
			"Arrays",
			[2]Profile{
				{33, "London"},
				{34, "Reykjavík"},
			},
			[]string{"London", "Reykjavík"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string

			// When
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			// Then
			assert.Equal(t, test.ExpectedCalls, got)
		})
	}

	t.Run("with maps", func(t *testing.T) {
		// Given
		aMap := map[string]string{
			"Foo": "Bar",
			"Baz": "Boz",
		}
		var got []string
		want := []string{"Bar", "Boz"}

		// When
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		// Then
		assert.ElementsMatch(t, want, got)
	})

	t.Run("with channels", func(t *testing.T) {
		// Given
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{33, "Berlin"}
			aChannel <- Profile{34, "Katowice"}
			close(aChannel)
		}()

		var got []string
		want := []string{"Berlin", "Katowice"}

		// When
		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		// Then
		assert.ElementsMatch(t, want, got)
	})

	t.Run("with function", func(t *testing.T) {
		// Given
		aFunction := func() (Profile, Profile) {
			return Profile{33, "Berlin"}, Profile{34, "Katowice"}
		}

		var got []string
		want := []string{"Berlin", "Katowice"}

		// When
		walk(aFunction, func(input string) {
			got = append(got, input)
		})

		// Then
		assert.ElementsMatch(t, want, got)
	})
}

// Left off at:  https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/reflection#write-the-test-first-9
