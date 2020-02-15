package linkedlist

import "github.com/leetcodeProblem/data"

func swapPairs(head *data.ListNode) *data.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	result := new(data.ListNode)
	next := head.Next
	prev := result
	for head != nil && next != nil {
		nextNext := next.Next
		prev.Next = next
		next.Next = head
		head.Next = nextNext
		prev = head
		head = nextNext
		if nextNext != nil {
			next = nextNext.Next
		}
	}
	return result.Next
}
