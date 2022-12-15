package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlusOne(t *testing.T) {
	tables := []struct {
		input    []int
		expected []int
	}{
		{
			[]int{1, 2, 3}, []int{1, 2, 4},
		},
		{
			[]int{4, 3, 2, 1}, []int{4, 3, 2, 2},
		},
		{
			[]int{9}, []int{1, 0},
		},
		{
			[]int{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3, 2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 0, 0, 6}, []int{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3, 2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 0, 0, 7},
		},
	}

	for _, table := range tables {
		k := plusOne(table.input)
		assert.Equal(t, table.expected, k, "they should be equal")
	}
}
