package linkedlist

// You are given two non-empty linked lists representing two non-negative integers.
// The digits are stored in reverse order and each of their nodes contain a single digit.
// Add the two numbers and return it as a linked list.

// You may assume the two numbers do not contain any leading zero, except the number 0 itself.

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := l1
	for l1 != nil && l2 != nil {
		l1.Val += l2.Val
		if l1.Val > 9 {
			l1.Val -= 10
			if l1.Next != nil {
				l1.Next.Val += 1
			} else {
				newNode := new(ListNode)
				newNode.Val = 1
				l1.Next = newNode
			}
		}
		l2 = l2.Next
		if l1.Next == nil {
			l1.Next = l2
			break
		}
		l1 = l1.Next
	}
	for l1 != nil && l1.Val > 9 {
		l1.Val -= 10
		if l1.Next != nil {
			l1.Next.Val += 1
		} else {
			newNode := new(ListNode)
			newNode.Val = 1
			l1.Next = newNode
		}
		l1 = l1.Next
	}
	return head
}
