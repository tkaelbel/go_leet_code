package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteDuplicates(t *testing.T) {
	tables := []struct {
		head     *ListNode
		expected *ListNode
	}{
		{
			&ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2}}}, &ListNode{Val: 1, Next: &ListNode{Val: 2}},
		},
		{
			&ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 3}}}}}, &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}},
		},
		{
			&ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0}}}}}, &ListNode{Val: 0},
		},
		{
			&ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 3}}}}}, &ListNode{Val: 0, Next: &ListNode{Val: 3}},
		},
	}

	for _, table := range tables {
		k := deleteDuplicates(table.head)
		assert.Equal(t, table.expected, k, "they should be equal")
	}
}
