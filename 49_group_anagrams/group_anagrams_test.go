package groupanagrams

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroupAnagrams(t *testing.T) {
	tables := []struct {
		strs     []string
		expected [][]string
	}{
		{
			[]string{"eat", "tea", "tan", "ate", "nat", "bat"}, [][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}},
		},
		{
			[]string{""}, [][]string{{""}},
		},
		{
			[]string{"a"}, [][]string{{"a"}},
		},
	}

	for _, table := range tables {
		k := groupAnagrams(table.strs)
		assert.Equal(t, table.expected, k, "they should be equal")
	}
}

func TestGroupAnagramsSort(t *testing.T) {
	tables := []struct {
		strs     []string
		expected [][]string
	}{
		{
			[]string{"eat", "tea", "tan", "ate", "nat", "bat"}, [][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}},
		},
		{
			[]string{""}, [][]string{{""}},
		},
		{
			[]string{"a"}, [][]string{{"a"}},
		},
	}

	for _, table := range tables {
		k := groupAnagramsSort(table.strs)
		assert.Equal(t, table.expected, k, "they should be equal")
	}
}
