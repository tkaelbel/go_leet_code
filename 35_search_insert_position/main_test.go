package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchInsert(t *testing.T) {
	tables := []struct {
		values         []int
		target         int
		expectedOutput int
	}{
		{
			[]int{1, 3, 5, 6}, 5, 2,
		},
		{
			[]int{1, 3, 5, 6}, 2, 1,
		},
		{
			[]int{1, 3, 5, 6}, 7, 4,
		},
	}

	for _, table := range tables {
		k := searchInsert(table.values, table.target)
		assert.Equal(t, table.expectedOutput, k, "they should be equal")
	}
}
