package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSet(t *testing.T) {
	t.Run("test adding and removing elements from a set", func(t *testing.T) {
		s := Set{}
		s.Add(1)
		s.Add(2)
		s.Add(3)

		assert.True(t, s.Contains(1))
		assert.True(t, s.Contains(2))
		assert.True(t, s.Contains(3))

		s.Remove(2)
		assert.False(t, s.Contains(2))
		assert.Equal(t, 2, s.Size())

		s.Add(4)
		assert.True(t, s.Contains(4))
		assert.Equal(t, 3, s.Size())
	})

	t.Run("test getting all elements from a set", func(t *testing.T) {
		s := Set{}
		s.Add(1)
		s.Add(2)
		s.Add(3)

		all := s.All()
		assert.ElementsMatch(t, []interface{}{1, 2, 3}, all)
	})
}
