package validpalindrome

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	tables := []struct {
		input    string
		expected bool
	}{
		{
			"A man, a plan, a canal: Panama", true,
		},
		{
			"race a car", false,
		},
		{
			" ", true,
		},
	}

	for _, table := range tables {
		x := isPalindrome(table.input)
		assert.Equal(t, table.expected, x, "they should be equal")
	}
}

func TestIsPalindrome1(t *testing.T) {
	tables := []struct {
		input    string
		expected bool
	}{
		{
			"A man, a plan, a canal: Panama", true,
		},
		{
			"race a car", false,
		},
		{
			" ", true,
		},
	}

	for _, table := range tables {
		x := isPalindrome1(table.input)
		assert.Equal(t, table.expected, x, "they should be equal")
	}
}
