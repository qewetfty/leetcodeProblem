package linkedlist

import "github.com/leetcodeProblem/data"

// Given the head of a singly linked list and two integers left and right where
// left <= right, reverse the nodes of the list from position left to position
// right, and return the reversed list.
//	Example 1:
//		Input: head = [1,2,3,4,5], left = 2, right = 4
//		Output: [1,4,3,2,5]
//	Example 2:
//		Input: head = [5], left = 1, right = 1
//		Output: [5]
//	Constraints:
//		The number of nodes in the list is n.
//		1 <= n <= 500
//		-500 <= Node.val <= 500
//		1 <= left <= right <= n
//	Follow up: Could you do it in one pass?

func reverseBetween(head *data.ListNode, m int, n int) *data.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &data.ListNode{Next: head}
	startPrev, startNode := dummy, head
	index := 1
	for index < m {
		startPrev, startNode = startPrev.Next, startNode.Next
		index++
	}
	startPrev.Next = nil
	var endPrev *data.ListNode
	endNode := startNode
	for index <= n {
		endPrev = endNode
		endNode = endNode.Next
		index++
	}
	endPrev.Next = nil
	startPrev.Next = reverse92(startNode)
	startNode.Next = endNode
	return dummy.Next
}

func reverse92(node *data.ListNode) *data.ListNode {
	var prev *data.ListNode
	cur := node
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev, cur = cur, next
	}
	return prev
}
