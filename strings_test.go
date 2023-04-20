package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToRune(t *testing.T) {
	t.Run("converts a string to an array of runes", func(t *testing.T) {
		str := "hello"
		expected := []rune{'h', 'e', 'l', 'l', 'o'}
		runes := ToRune(str)
		assert.Equal(t, expected, runes)
	})

	t.Run("converts an empty string to an empty array of runes", func(t *testing.T) {
		str := ""
		expected := []rune{}
		runes := ToRune(str)
		assert.Equal(t, expected, runes)
	})
}

func TestToFrequency(t *testing.T) {
	t.Run("returns the frequency of characters in a string", func(t *testing.T) {
		str := "hello"
		expected := Frequencies{
			'h': 1,
			'e': 1,
			'l': 2,
			'o': 1,
		}
		freq := ToFrequency(str)
		assert.Equal(t, expected, freq)
	})

	t.Run("returns an empty map for an empty string", func(t *testing.T) {
		str := ""
		expected := Frequencies{}
		freq := ToFrequency(str)
		assert.Equal(t, expected, freq)
	})
}

func TestToSet(t *testing.T) {
	t.Run("converts a frequency map to a set", func(t *testing.T) {
		freq := Frequencies{
			'h': 1,
			'e': 1,
			'l': 2,
			'o': 1,
		}
		expected := Set{}
		expected.Add('h')
		expected.Add('e')
		expected.Add('l')
		expected.Add('o')

		s := freq.ToSet()

		assert.Equal(t, expected, s)
	})

	t.Run("returns an empty set for an empty frequency map", func(t *testing.T) {
		freq := Frequencies{}
		expected := Set{}
		s := freq.ToSet()
		assert.Equal(t, expected, s)
	})
}

func TestToArray(t *testing.T) {
	t.Run("converts a frequency map to an array of frequencies", func(t *testing.T) {
		freq := Frequencies{
			'h': 1,
			'e': 1,
			'l': 2,
			'o': 1,
		}
		expected := []int{1, 1, 2, 1}
		freqs := freq.ToArray()
		assert.ElementsMatch(t, expected, freqs)
	})

	t.Run("returns an empty array for an empty frequency map", func(t *testing.T) {
		freq := Frequencies{}
		expected := []int{}
		freqs := freq.ToArray()
		assert.Equal(t, expected, freqs)
	})
}
