package linkedlist

import "leetcodeProblem/data"

//Given the head of a singly linked list, return the middle node of the linked
//list.
//
// If there are two middle nodes, return the second middle node.
//
//
// Example 1:
//
//
//Input: head = [1,2,3,4,5]
//Output: [3,4,5]
//Explanation: The middle node of the list is node 3.
//
//
// Example 2:
//
//
//Input: head = [1,2,3,4,5,6]
//Output: [4,5,6]
//Explanation: Since the list has two middle nodes with values 3 and 4, we
//return the second one.
//
//
//
// Constraints:
//
//
// The number of nodes in the list is in the range [1, 100].
// 1 <= Node.val <= 100
//
// Related Topics Linked List Two Pointers 👍 3958 👎 103

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *data.ListNode) *data.ListNode {
	l, cur := 0, head
	for cur != nil {
		l++
		cur = cur.Next
	}
	mid := l >> 1
	i, cur := 0, head
	for i < mid {
		cur = cur.Next
		i++
	}
	return cur
}

//leetcode submit region end(Prohibit modification and deletion)
