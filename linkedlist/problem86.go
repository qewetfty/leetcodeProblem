package linkedlist

import "github.com/leetcodeProblem/data"

// Given the head of a linked list and a value x, partition it such that all
// nodes less than x come before nodes greater than or equal to x.
// You should preserve the original relative order of the nodes in each of the two partitions.
//	Example 1:
//		Input: head = [1,4,3,2,5,2], x = 3
//		Output: [1,2,2,4,3,5]
//	Example 2:
//		Input: head = [2,1], x = 2
//		Output: [1,2]
//	Constraints:
//		The number of nodes in the list is in the range [0, 200].
//		-100 <= Node.val <= 100
//		-200 <= x <= 200

// 使用列表存储被删除的节点
func partition(head *data.ListNode, x int) *data.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &data.ListNode{Next: head}
	prev, cur := dummy, head
	moveList := make([]*data.ListNode, 0)
	for cur != nil {
		if cur.Val >= x {
			next := cur.Next
			prev.Next = next
			cur.Next = nil
			moveList = append(moveList, cur)
			cur = next
		} else {
			prev, cur = cur, cur.Next
		}
	}
	for _, node := range moveList {
		prev.Next = node
		prev = prev.Next
	}
	return dummy.Next
}
