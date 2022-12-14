package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthOfLastWord(t *testing.T) {
	tables := []struct {
		s              string
		expectedLength int
	}{
		{
			"Hello World", 5,
		},
		{
			"   fly me   to   the moon  ", 4,
		},
		{
			"luffy is still joyboy", 6,
		},
		{
			"Today is a nice day", 3,
		},
		{
			"a", 1,
		},
		{
			"a ", 1,
		},
		{
			"day", 3,
		},
	}

	for _, table := range tables {
		k := lengthOfLastWord(table.s)
		assert.Equal(t, table.expectedLength, k, "they should be equal")
	}
}
