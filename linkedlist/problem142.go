package linkedlist

import (
	"github.com/leetcodeProblem/data"
)

// Given a linked list, return the node where the cycle begins. If there is no cycle, return null.

// To represent a cycle in the given linked list,
// we use an integer pos which represents the position (0-indexed) in the linked list where tail connects to.
// If pos is -1, then there is no cycle in the linked list.

// Note: Do not modify the linked list.

// Three steps to get the result:
// 1.Check if there is a cycle.
// 2.Find the loop number n.
// 3.2 pointers. 1st move n steps then 2nd move. Where the meet is the answer
func detectCycle(head *data.ListNode) *data.ListNode {
	// Check if there is a cycle
	if head == nil || head.Next == nil {
		return nil
	}
	slow := head
	fast := head.Next
	for {
		if slow == fast {
			break
		}
		if fast.Next == nil || fast.Next.Next == nil {
			return nil
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	// find the loop number n
	num := 1
	fast = fast.Next
	for slow != fast {
		fast = fast.Next
		num++
	}
	// 2 pointers move
	i := 0
	first := head
	second := head
	for i != num {
		first = first.Next
		i++
	}
	for first != second {
		first = first.Next
		second = second.Next
	}
	return second
}

func detectCycle2(head *data.ListNode) *data.ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	slow := head
	fast := head
	for slow != nil && fast != nil {
		if fast.Next == nil || fast.Next.Next == nil {
			return nil
		}
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			break
		}
	}

	pointer := head
	for pointer != fast {
		pointer = pointer.Next
		fast = fast.Next
	}
	return pointer
}
