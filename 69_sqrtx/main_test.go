package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSqrt(t *testing.T) {
	tables := []struct {
		x        int
		expected int
	}{
		{
			4, 2,
		},
		{
			8, 2,
		},
		{
			16, 4,
		},
		{
			9, 3,
		},
	}

	for _, table := range tables {
		k := mySqrt(table.x)
		assert.Equal(t, table.expected, k, "they should be equal")
	}
}
