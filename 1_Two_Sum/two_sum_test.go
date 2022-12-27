package twosum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	tables := []struct {
		nums     []int
		t        int
		expected []int
	}{
		{
			[]int{2, 7, 11, 15}, 9, []int{0, 1},
		},
		{
			[]int{3, 2, 4}, 6, []int{1, 2},
		},
		{
			[]int{3, 3}, 6, []int{0, 1},
		},
	}

	for _, table := range tables {
		k := twoSum(table.nums, table.t)
		assert.Equal(t, table.expected, k, "they should be equal")
	}
}

func TestTwoSum2(t *testing.T) {
	tables := []struct {
		nums     []int
		t        int
		expected []int
	}{
		{
			[]int{2, 7, 11, 15}, 9, []int{0, 1},
		},
		{
			[]int{3, 2, 4}, 6, []int{1, 2},
		},
		{
			[]int{3, 3}, 6, []int{0, 1},
		},
	}

	for _, table := range tables {
		k := twoSum2(table.nums, table.t)
		assert.Equal(t, table.expected, k, "they should be equal")
	}
}
