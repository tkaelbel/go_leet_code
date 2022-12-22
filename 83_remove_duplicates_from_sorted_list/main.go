package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	var dummyHead *ListNode
	if head == nil {
		return dummyHead
	}

	dummyHead = &ListNode{
		Val: head.Val,
	}

	cur := dummyHead
	start := head.Next
	for start != nil {
		if cur.Val != start.Val {
			cur.Next = &ListNode{Val: start.Val}
			cur = cur.Next
		}
		start = start.Next
	}

	return dummyHead
}
