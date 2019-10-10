package linkedlist

import "github.com/leetcodeProblem/data"

// Given a linked list, determine if it has a cycle in it.
//To represent a cycle in the given linked list, we use an integer pos which
// represents the position (0-indexed) in the linked list where tail connects to.
// If pos is -1, then there is no cycle in the linked list.

func main() {

}

func hasCycle(head *data.ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow := head
	fast := head.Next
	for {
		if slow == fast {
			return true
		}
		if fast.Next == nil || fast.Next.Next == nil {
			return false
		}
		fast = fast.Next.Next
		slow = slow.Next
	}
}
