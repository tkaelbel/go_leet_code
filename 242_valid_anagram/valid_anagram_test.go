package validanagram

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsAnagram(t *testing.T) {
	tables := []struct {
		s        string
		t        string
		expected bool
	}{
		{
			"anagram", "nagaram", true,
		},
		{
			"rat", "car", false,
		},
	}

	for _, table := range tables {
		k := isAnagramMap(table.s, table.t)
		assert.Equal(t, table.expected, k, "they should be equal")
	}
}

func TestIsAnagramMap2(t *testing.T) {
	tables := []struct {
		s        string
		t        string
		expected bool
	}{
		{
			"anagram", "nagaram", true,
		},
		{
			"rat", "car", false,
		},
	}

	for _, table := range tables {
		k := isAnagramMap2(table.s, table.t)
		assert.Equal(t, table.expected, k, "they should be equal")
	}
}

func TestIsAnagramSort(t *testing.T) {
	tables := []struct {
		s        string
		t        string
		expected bool
	}{
		{
			"anagram", "nagaram", true,
		},
		{
			"rat", "car", false,
		},
	}

	for _, table := range tables {
		k := isAnagramSort(table.s, table.t)
		assert.Equal(t, table.expected, k, "they should be equal")
	}
}

// func TestContainsDuplicateMap(t *testing.T) {
// 	tables := []struct {
// 		nums     []int
// 		expected bool
// 	}{
// 		{
// 			[]int{1, 2, 3, 1}, true,
// 		},
// 		{
// 			[]int{1, 2, 3, 4}, false,
// 		},
// 		{
// 			[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true,
// 		},
// 	}

// 	for _, table := range tables {
// 		k := containsDuplicateMap(table.nums)
// 		assert.Equal(t, table.expected, k, "they should be equal")
// 	}
// }
