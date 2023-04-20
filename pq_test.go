package helper

import (
	"container/heap"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPQ(t *testing.T) {
	// Define a compare function that sorts elements in decreasing order.
	compare := func(list []interface{}, i int, j int) bool {
		return list[i].(int) > list[j].(int)
	}

	t.Run("test priority queue initialization", func(t *testing.T) {
		pq := &PQ{
			list:    []interface{}{3, 1, 4, 1},
			Compare: compare,
		}

		heap.Init(pq)

		// Assert that the priority queue is initialized with the correct elements.
		expected := []int{4, 3, 1, 1}
		for _, exp := range expected {
			val := heap.Pop(pq).(int)
			assert.Equal(t, exp, val)
		}
	})

	t.Run("test adding and popping elements from priority queue", func(t *testing.T) {
		pq := &PQ{
			list:    []interface{}{3, 1, 4, 1},
			Compare: compare,
		}

		heap.Init(pq)

		heap.Push(pq, 5)
		heap.Push(pq, 2)

		// Assert that the priority queue has the correct size.
		assert.Equal(t, 6, pq.Len())

		// Assert that the elements are in the correct order.
		expected := []int{5, 4, 3, 2, 1, 1}
		for _, exp := range expected {
			val := heap.Pop(pq).(int)
			assert.Equal(t, exp, val)
		}

		// Assert that the priority queue is empty.
		assert.Equal(t, 0, pq.Len())
	})
}
