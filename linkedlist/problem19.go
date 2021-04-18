package linkedlist

import "github.com/leetcodeProblem/data"

// Given the head of a linked list, remove the nth node from the end of the list and return its head.
// Follow up: Could you do this in one pass?
//	Example 1:
//		Input: head = [1,2,3,4,5], n = 2
//		Output: [1,2,3,5]
//	Example 2:
//		Input: head = [1], n = 1
//		Output: []
//	Example 3:
//		Input: head = [1,2], n = 1
//		Output: [1]
//	Constraints:
//		The number of nodes in the list is sz.
//		1 <= sz <= 30
//		0 <= Node.val <= 100
//		1 <= n <= sz

func removeNthFromEnd(head *data.ListNode, n int) *data.ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	prev, after := head, head
	var move *data.ListNode
	i := 0
	for prev != nil {
		prev = prev.Next
		if i >= n {
			move = after
			after = after.Next
		}
		i++
	}
	if move == nil {
		return head.Next
	}
	move.Next = move.Next.Next
	return head
}
