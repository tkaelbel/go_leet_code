package topkfrequentelements

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKFrequent(t *testing.T) {
	tables := []struct {
		input    []int
		k        int
		expected []int
	}{
		{
			[]int{1, 1, 1, 2, 2, 3}, 2, []int{1, 2},
		},
		{
			[]int{1}, 1, []int{1},
		},
		{
			[]int{4, 1, -1, 2, -1, 2, 3}, 2, []int{-1, 2},
		},
		{
			[]int{-1, -1}, 1, []int{-1},
		},
		{
			[]int{1, 2}, 2, []int{1, 2},
		},
	}

	for _, table := range tables {
		x := topKFrequent(table.input, table.k)
		assert.Equal(t, table.expected, x, "they should be equal")
	}
}
