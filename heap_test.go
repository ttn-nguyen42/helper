package helper

import (
	"container/heap"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxHeap(t *testing.T) {
	t.Run("should have correct order", func(t *testing.T) {
		elems := []int{10, 5, 8, 13, 1, 9}

		h := &MaxHeap{}

		for _, e := range elems {
			heap.Push(h, e)
		}

		test := []int{}
		for h.Len() > 0 {
			e := heap.Pop(h).(int)
			test = append(test, e)
		}

		assert.Equal(t, []int{13, 10, 9, 8, 5, 1}, test)
	})

}
