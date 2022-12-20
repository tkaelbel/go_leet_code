package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClimbStairs(t *testing.T) {
	tables := []struct {
		n        int
		expected int
	}{
		{
			2, 2,
		},
		{
			3, 3,
		},
		{
			4, 5,
		},
	}

	for _, table := range tables {
		k := climbStairs(table.n)
		assert.Equal(t, table.expected, k, "they should be equal")
	}
}
