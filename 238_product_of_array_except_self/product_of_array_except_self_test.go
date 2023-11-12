package productofarrayexceptself

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductExceptSelf(t *testing.T) {
	tables := []struct {
		input    []int
		expected []int
	}{
		{
			[]int{1, 2, 3, 4}, []int{24, 12, 8, 6},
		},
		{
			[]int{-1, 1, 0, -3, 3}, []int{0, 0, 9, 0, 0},
		},
	}

	for _, table := range tables {
		x := productExceptSelf(table.input)
		assert.Equal(t, table.expected, x, "they should be equal")
	}

}
