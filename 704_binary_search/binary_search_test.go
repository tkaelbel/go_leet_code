package binarysearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	tables := []struct {
		input    []int
		target   int
		expected int
	}{
		{
			[]int{1}, 1, 0,
		},
		{
			[]int{1, 2}, 2, 1,
		},
		{
			[]int{-1, 0, 3, 5, 9, 12}, 9, 4,
		},
		{
			[]int{-1, 0, 3, 5, 9, 12}, 2, -1,
		},
	}

	for _, table := range tables {
		x := search(table.input, table.target)
		assert.Equal(t, table.expected, x, "they should be equal")
	}

}
