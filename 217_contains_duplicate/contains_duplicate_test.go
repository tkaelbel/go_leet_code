package containsduplicate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainsDuplicate(t *testing.T) {
	tables := []struct {
		nums     []int
		expected bool
	}{
		{
			[]int{1, 2, 3, 1}, true,
		},
		{
			[]int{1, 2, 3, 4}, false,
		},
		{
			[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true,
		},
	}

	for _, table := range tables {
		k := containsDuplicate(table.nums)
		assert.Equal(t, table.expected, k, "they should be equal")
	}
}

func TestContainsDuplicateMap(t *testing.T) {
	tables := []struct {
		nums     []int
		expected bool
	}{
		{
			[]int{1, 2, 3, 1}, true,
		},
		{
			[]int{1, 2, 3, 4}, false,
		},
		{
			[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true,
		},
	}

	for _, table := range tables {
		k := containsDuplicateMap(table.nums)
		assert.Equal(t, table.expected, k, "they should be equal")
	}
}
