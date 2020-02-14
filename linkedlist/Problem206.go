package linkedlist

import "github.com/leetcodeProblem/data"

// Reverse a singly linked list.
//
// Example:
//
// Input: 1->2->3->4->5->NULL
// Output: 5->4->3->2->1->NULL
// Follow up:
//
// A linked list can be reversed either iteratively or recursively. Could you implement both?

func reverseListByIterativeSolution(head *data.ListNode) *data.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var prev *data.ListNode = nil
	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev
}

func reverseListByRecursiveSolution(head *data.ListNode) *data.ListNode {
	return reverseLinkedList(head, nil)
}

func reverseLinkedList(head *data.ListNode, newHead *data.ListNode) *data.ListNode {
	if head == nil {
		return newHead
	}
	next := head.Next
	head.Next = newHead
	return reverseLinkedList(next, head)
}
