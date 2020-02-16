package linkedlist

import (
	"fmt"
	"github.com/leetcodeProblem/data"
)

// Merge two sorted linked lists and return it as a new list.
// The new list should be made by splicing together the nodes of the first two lists.
//
//	Example:
//		Input: 1->2->4, 1->3->4
//		Output: 1->1->2->3->4->4

func mergeTwoLists(l1 *data.ListNode, l2 *data.ListNode) *data.ListNode {
	dummy := new(data.ListNode)
	merge(dummy, l1, l2)
	return dummy.Next
}

func merge(head, l1, l2 *data.ListNode) {
	if l1 == nil {
		head.Next = l2
		return
	}
	if l2 == nil {
		head.Next = l1
		return
	}
	if l1.Val <= l2.Val {
		head.Next = l1
		merge(head.Next, l1.Next, l2)
	} else {
		head.Next = l2
		merge(head.Next, l1, l2.Next)
	}
}

func mergeTwoListsNonRecursive(l1 *data.ListNode, l2 *data.ListNode) *data.ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	dummy := new(data.ListNode)
	returnNode := dummy
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			returnNode.Next = l1
			l1 = l1.Next
		} else {
			returnNode.Next = l2
			l2 = l2.Next
		}
		returnNode = returnNode.Next
	}
	if l1 == nil {
		returnNode.Next = l2
	}
	if l2 == nil {
		returnNode.Next = l1
	}
	return dummy.Next
}

func testProblem21() {
	a := data.NewListNode(1)
	b := data.NewListNode(2)
	c := data.NewListNode(4)
	d := data.NewListNode(1)
	e := data.NewListNode(3)
	f := data.NewListNode(4)
	a.Next = b
	b.Next = c
	d.Next = e
	e.Next = f

	res := mergeTwoListsNonRecursive(a, d)
	fmt.Println(res.ToString())
}
