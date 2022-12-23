package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteDuplicates(t *testing.T) {
	tables := []struct {
		nums1    []int
		m        int
		nums2    []int
		n        int
		expected []int
	}{
		{
			[]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3, []int{1, 2, 2, 3, 5, 6},
		},
		{
			[]int{1}, 1, []int{}, 0, []int{1},
		},
		{
			[]int{0}, 0, []int{1}, 1, []int{1},
		},
		{
			[]int{0, 0, 0, 0, 0}, 0, []int{1, 2, 3, 4, 5}, 5, []int{1, 2, 3, 4, 5},
		},
		{
			[]int{4, 0, 0, 0, 0, 0}, 1, []int{1, 2, 3, 5, 6}, 5, []int{1, 2, 3, 4, 5, 6},
		},
		{
			[]int{1, 2, 4, 5, 6, 0}, 5, []int{3}, 1, []int{1, 2, 3, 4, 5, 6},
		},
		{
			[]int{-1, 0, 0, 3, 3, 3, 0, 0, 0}, 6, []int{1, 2, 2}, 3, []int{-1, 0, 0, 1, 2, 2, 3, 3, 3},
		},
	}

	for _, table := range tables {
		merge(table.nums1, table.m, table.nums2, table.n)
		assert.Equal(t, table.expected, table.nums1, "they should be equal")
	}
}
