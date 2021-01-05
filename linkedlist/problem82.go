package linkedlist

import "github.com/leetcodeProblem/data"

// Given the head of a sorted linked list, delete all nodes that have duplicate
// numbers, leaving only distinct numbers from the original list. Return the
// linked list sorted as well.
//	Example 1:
//		Input: head = [1,2,3,3,4,4,5]
//		Output: [1,2,5]
//	Example 2:
//		Input: head = [1,1,1,2,3]
//		Output: [2,3]
//	Constraints:
//		The number of nodes in the list is in the range [0, 300].
//		-100 <= Node.val <= 100
//		The list is guaranteed to be sorted in ascending order.

func deleteDuplicates(head *data.ListNode) *data.ListNode {
	dummy := &data.ListNode{Next: head}
	prev, cur := dummy, head
	for cur != nil {
		if cur.Next != nil && cur.Next.Val == cur.Val {
			delNum := cur.Val
			for cur != nil && cur.Val == delNum {
				next := cur.Next
				prev.Next = next
				cur.Next = nil
				cur = next
			}
		} else {
			cur = cur.Next
			prev = prev.Next
		}
	}
	return dummy.Next
}
