package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxInArray(t *testing.T) {
	t.Run("returns the maximum element in an array of integers", func(t *testing.T) {
		arr := []interface{}{10, 7, 8, 9, 1, 5}
		larger := func(a, b interface{}) bool {
			return a.(int) > b.(int)
		}
		max := MaxInArray(arr, larger)
		assert.Equal(t, 10, max)
	})

	t.Run("returns the maximum element in an array of strings", func(t *testing.T) {
		arr := []interface{}{"cat", "dog", "bird", "elephant"}
		larger := func(a, b interface{}) bool {
			return len(a.(string)) > len(b.(string))
		}
		max := MaxInArray(arr, larger)
		assert.Equal(t, "elephant", max)
	})

	t.Run("returns the maximum element in an empty array", func(t *testing.T) {
		arr := []interface{}{}
		larger := func(a, b interface{}) bool {
			return false
		}
		max := MaxInArray(arr, larger)
		assert.Nil(t, max)
	})

	t.Run("returns the maximum element in an array of one element", func(t *testing.T) {
		arr := []interface{}{10}
		larger := func(a, b interface{}) bool {
			return a.(int) > b.(int)
		}
		max := MaxInArray(arr, larger)
		assert.Equal(t, 10, max)
	})
}

func TestMinInArray(t *testing.T) {
	t.Run("returns the minimum element in an array of integers", func(t *testing.T) {
		arr := []interface{}{10, 7, 8, 9, 1, 5}
		larger := func(a, b interface{}) bool {
			return a.(int) > b.(int)
		}
		min := MinInArray(arr, larger)
		assert.Equal(t, 1, min)
	})

	t.Run("returns the minimum element in an array of strings", func(t *testing.T) {
		arr := []interface{}{"cat", "dog", "bird", "elephant"}
		larger := func(a, b interface{}) bool {
			return len(a.(string)) > len(b.(string))
		}
		min := MinInArray(arr, larger)
		assert.Equal(t, "cat", min)
	})

	t.Run("returns the minimum element in an empty array", func(t *testing.T) {
		arr := []interface{}{}
		larger := func(a, b interface{}) bool {
			return false
		}
		min := MinInArray(arr, larger)
		assert.Nil(t, min)
	})

	t.Run("returns the minimum element in an array of one element", func(t *testing.T) {
		arr := []interface{}{10}
		larger := func(a, b interface{}) bool {
			return a.(int) > b.(int)
		}
		min := MinInArray(arr, larger)
		assert.Equal(t, 10, min)
	})
}
