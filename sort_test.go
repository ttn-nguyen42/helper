package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuickSort(t *testing.T) {
	t.Run("sorts an array of integers", func(t *testing.T) {
		arr := []interface{}{10, 7, 8, 9, 1, 5}
		expected := []interface{}{1, 5, 7, 8, 9, 10}
		larger := func(a, b interface{}) bool {
			return a.(int) > b.(int)
		}
		QuickSort(arr, larger)
		assert.Equal(t, expected, arr)
	})

	t.Run("sorts an array of strings", func(t *testing.T) {
		arr := []interface{}{"cat", "dog", "bird", "elephant"}
		expected := []interface{}{"bird", "cat", "dog", "elephant"}
		larger := func(a, b interface{}) bool {
			return a.(string) > b.(string)
		}
		QuickSort(arr, larger)
		assert.Equal(t, expected, arr)
	})

	t.Run("sorts an empty array", func(t *testing.T) {
		arr := []interface{}{}
		expected := []interface{}{}
		larger := func(a, b interface{}) bool {
			return false
		}
		QuickSort(arr, larger)
		assert.Equal(t, expected, arr)
	})

	t.Run("sorts an array of one element", func(t *testing.T) {
		arr := []interface{}{10}
		expected := []interface{}{10}
		larger := func(a, b interface{}) bool {
			return false
		}
		QuickSort(arr, larger)
		assert.Equal(t, expected, arr)
	})
}
