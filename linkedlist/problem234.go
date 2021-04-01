package linkedlist

import "github.com/leetcodeProblem/data"

// Given the head of a singly linked list, return true if it is a palindrome.
//	Example 1:
//		Input: head = [1,2,2,1]
//		Output: true
//	Example 2:
//		Input: head = [1,2]
//		Output: false
//	Constraints:
//		The number of nodes in the list is in the range [1, 105].
//		0 <= Node.val <= 9

// 反转链表，O(1)空间复杂度
func isPalindrome2(head *data.ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}
	if fast != nil {
		slow = slow.Next
	}
	slow = reverse(slow)
	fast = head
	for slow != nil {
		if slow.Val != fast.Val {
			return false
		}
		slow, fast = slow.Next, fast.Next
	}
	return true
}

func reverse(head *data.ListNode) *data.ListNode {
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

// 使用数组，O(n)空间复杂度
func isPalindrome(head *data.ListNode) bool {
	arr := make([]int, 0)
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	l := len(arr)
	for i := 0; i < l/2; i++ {
		if arr[i] != arr[l-i-1] {
			return false
		}
	}
	return true
}
