package helper

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadString(t *testing.T) {
	t.Run("should success when string entered correctly", func(t *testing.T) {
		test := "test subject"
		buf := strings.NewReader(test)
		sc := bufio.NewScanner(buf)

		want, err := ReadString(sc, 1)

		assert.NoError(t, err)
		assert.Equal(t, test, want[0])
	})

	t.Run("should return error when number of times is invalid", func(t *testing.T) {
		test := "test subject"
		buf := strings.NewReader(test)
		sc := bufio.NewScanner(buf)

		_, err := ReadString(sc, -100)

		assert.Error(t, err)
		assert.Equal(t, ErrInvalidTimes, err)
	})

	t.Run("should return error when the scanner is nil", func(t *testing.T) {
		_, err := ReadString(nil, 100)

		assert.Error(t, err)
		assert.Equal(t, ErrNilScanner, err)
	})
}

func TestReadInt(t *testing.T) {
	t.Run("should success when string entered correctly", func(t *testing.T) {
		test := "10"
		wanted := 10
		buf := strings.NewReader(test)
		sc := bufio.NewScanner(buf)

		want, err := ReadInteger(sc, 1)

		assert.NoError(t, err)
		assert.Equal(t, wanted, want[0])
	})

	t.Run("should fail when number is invalid", func(t *testing.T) {
		test := "abcd"
		buf := strings.NewReader(test)
		sc := bufio.NewScanner(buf)

		_, err := ReadInteger(sc, 1)

		assert.Error(t, err)
		assert.Equal(t, ErrIntegerInvalid, err)
	})

	t.Run("should fail when the number of times is invalid", func(t *testing.T) {
		test := "10"
		buf := strings.NewReader(test)
		sc := bufio.NewScanner(buf)

		_, err := ReadInteger(sc, -100)

		assert.Error(t, err)
		assert.Equal(t, ErrInvalidTimes, err)
	})
}

func TestReadIntArray(t *testing.T) {
	t.Run("should success when array entered correctly", func(t *testing.T) {
		test := "10 20 30\n50 60 70\n"
		wanted := [][]int{
			{10, 20, 30},
			{50, 60, 70},
		}
		buf := strings.NewReader(test)
		sc := bufio.NewScanner(buf)

		want, err := ReadIntArray(sc, 2)

		assert.NoError(t, err)
		assert.Equal(t, wanted, want)
	})

	t.Run("should fail when number is invalid", func(t *testing.T) {
		test := "10 20 cd\n50 60 70\n"
		buf := strings.NewReader(test)
		sc := bufio.NewScanner(buf)

		_, err := ReadIntArray(sc, 2)

		assert.Error(t, ErrIntegerInvalid, err)
	})

	t.Run("should fail when the number of times is invalid", func(t *testing.T) {
		test := "10 20 30\n50 60 70\n"
		buf := strings.NewReader(test)
		sc := bufio.NewScanner(buf)

		_, err := ReadIntArray(sc, -100)

		assert.Equal(t, ErrInvalidTimes, err)
	})
}

func TestReadStringArray(t *testing.T) {
	t.Run("should success when array entered correctly", func(t *testing.T) {
		test := "a b c\nc d e\n"
		wanted := [][]string{
			{"a", "b", "c"},
			{"c", "d", "e"},
		}
		buf := strings.NewReader(test)
		sc := bufio.NewScanner(buf)

		want, err := ReadStringArray(sc, 2)

		assert.NoError(t, err)
		assert.Equal(t, wanted, want)
	})

	t.Run("should fail when the number of times is invalid", func(t *testing.T) {
		test := "a b c\nc d e\n"
		buf := strings.NewReader(test)
		sc := bufio.NewScanner(buf)

		_, err := ReadStringArray(sc, -100)

		assert.Equal(t, ErrInvalidTimes, err)
	})

	t.Run("should fail when the scanner is nil", func(t *testing.T) {
		_, err := ReadStringArray(nil, 100)

		assert.Equal(t, ErrNilScanner, err)
	})
}
