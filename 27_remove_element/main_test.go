package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveElement(t *testing.T) {
	tables := []struct {
		values         []int
		removeNumber   int
		expectedValues []int
		expectedK      int
	}{
		{
			[]int{3, 2, 2, 3}, 3, []int{2, 2, 3, 3}, 2,
		},
		{
			[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2, []int{0, 0, 1, 3, 4, 2, 3, 4}, 5,
		},
	}

	for _, table := range tables {
		k := removeElement(table.values, table.removeNumber)
		assert.Equal(t, table.expectedValues, table.values, "they should be equal")
		assert.Equal(t, table.expectedK, k, "they should be equal")
	}
}
