package proptests

import (
	"fmt"
	"log"
	"testing"
	"testing/quick"

	"github.com/stretchr/testify/assert"
)

var cases = []struct {
	Arabic uint16
	Roman  string
}{
	{Arabic: 1, Roman: "I"},
	{Arabic: 2, Roman: "II"},
	{Arabic: 3, Roman: "III"},
	{Arabic: 4, Roman: "IV"},
	{Arabic: 5, Roman: "V"},
	{Arabic: 6, Roman: "VI"},
	{Arabic: 7, Roman: "VII"},
	{Arabic: 8, Roman: "VIII"},
	{Arabic: 9, Roman: "IX"},
	{Arabic: 10, Roman: "X"},
	{Arabic: 11, Roman: "XI"},
	{Arabic: 14, Roman: "XIV"},
	{Arabic: 15, Roman: "XV"},
	{Arabic: 18, Roman: "XVIII"},
	{Arabic: 19, Roman: "XIX"},
	{Arabic: 20, Roman: "XX"},
	{Arabic: 39, Roman: "XXXIX"},
	{Arabic: 40, Roman: "XL"},
	{Arabic: 47, Roman: "XLVII"},
	{Arabic: 49, Roman: "XLIX"},
	{Arabic: 50, Roman: "L"},
	{Arabic: 100, Roman: "C"},
	{Arabic: 90, Roman: "XC"},
	{Arabic: 400, Roman: "CD"},
	{Arabic: 500, Roman: "D"},
	{Arabic: 900, Roman: "CM"},
	{Arabic: 1000, Roman: "M"},
	{Arabic: 1984, Roman: "MCMLXXXIV"},
	{Arabic: 3999, Roman: "MMMCMXCIX"},
	{Arabic: 2014, Roman: "MMXIV"},
	{Arabic: 1006, Roman: "MVI"},
	{Arabic: 798, Roman: "DCCXCVIII"},
}

func TestRomanNumerals(t *testing.T) {

	for _, test := range cases {
		t.Run(fmt.Sprintf("%d => %q", test.Arabic, test.Roman), func(t *testing.T) {
			// Given
			// When
			got := ConvertToRoman(test.Arabic)

			// Then
			assert.Equal(t, test.Roman, got)
		})
	}
}

func TestConvertingToArabic(t *testing.T) {

	for _, test := range cases {
		t.Run(fmt.Sprintf("%q => %d", test.Roman, test.Arabic), func(t *testing.T) {
			// Given
			// When
			got := ConvertToArabic(test.Roman)

			// Then
			assert.Equal(t, test.Arabic, got)
		})
	}
}

func TestPropertiesOfConversion(t *testing.T) {
	// Given
	assertion := func(arabic uint16) bool {
		if arabic < 0 || arabic > 3999 {
			log.Println(arabic)
			return true
		}
		roman := ConvertToRoman(arabic)
		fromRoman := ConvertToArabic(roman)
		return fromRoman == arabic
	}

	// When
	err := quick.Check(assertion, nil)

	// Then
	assert.NoError(t, err)
}
