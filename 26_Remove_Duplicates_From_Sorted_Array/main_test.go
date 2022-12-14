package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicates(t *testing.T) {
	tables := []struct {
		values         []int
		expectedValues []int
		expectedK      int
	}{
		{
			[]int{1, 1, 2}, []int{1, 2, 2}, 2,
		},
		{
			[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, []int{0, 1, 2, 3, 4, 2, 2, 3, 3, 4}, 5,
		},
		{
			[]int{-1, 0, 0, 0, 0, 3, 3}, []int{-1, 0, 3, 0, 0, 3, 3}, 3,
		},
		{
			[]int{0, 0, 0, 0, 0}, []int{0, 0, 0, 0, 0}, 1,
		},
		{
			[]int{-3, -1, 0, 0}, []int{-3, -1, 0, 0}, 3,
		},
	}

	for _, table := range tables {
		k := removeDuplicates(table.values)
		assert.Equal(t, table.expectedValues, table.values, "they should be equal")
		assert.Equal(t, table.expectedK, k, "they should be equal")
	}
}
