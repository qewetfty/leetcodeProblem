package linkedlist

import (
	"fmt"
	"github.com/leetcodeProblem/data"
)

// Given a linked list, reverse the nodes of a linked list k at a time and return its modified list.
//
// k is a positive integer and is less than or equal to the length of the linked list.
// If the number of nodes is not a multiple of k then left-out nodes in the end should remain as it is.
//
// Example:
//	Given this linked list: 1->2->3->4->5
//	For k = 2, you should return: 2->1->4->3->5
//	For k = 3, you should return: 3->2->1->4->5

// Note:
//	Only constant extra memory is allowed.
//	You may not alter the values in the list's nodes, only nodes itself may be changed.

func reverseKGroup(head *data.ListNode, k int) *data.ListNode {
	if head == nil || head.Next == nil || k == 1 {
		return head
	}
	n := 0
	h := head
	for h != nil {
		h = h.Next
		n++
	}
	loopNum := n / k
	dummy := new(data.ListNode)
	dummy.Next = head
	parent := dummy
	var prev, next *data.ListNode
	curr := head
	for i := 0; i < loopNum; i++ {
		j := 0
		currHead := parent.Next
		for j < k {
			next = curr.Next
			curr.Next = prev
			prev = curr
			curr = next
			j++
		}
		currHead.Next = curr
		parent.Next = prev
		parent = currHead
	}
	return dummy.Next
}

func testProblem25() {
	a := data.NewListNode(1)
	b := data.NewListNode(2)
	c := data.NewListNode(3)
	d := data.NewListNode(4)
	e := data.NewListNode(5)
	a.Next = b
	b.Next = c
	c.Next = d
	d.Next = e
	fmt.Println(a.ToString())
	h := reverseKGroup(a, 4)
	fmt.Println(h.ToString())
}
