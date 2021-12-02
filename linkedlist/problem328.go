package linkedlist

import "leetcodeProblem/data"

//Given the head of a singly linked list, group all the nodes with odd indices
//together followed by the nodes with even indices, and return the reordered list.
//
//
// The first node is considered odd, and the second node is even, and so on.
//
// Note that the relative order inside both the even and odd groups should
//remain as it was in the input.
//
// You must solve the problem in O(1) extra space complexity and O(n) time
//complexity.
//
//
// Example 1:
//
//
//Input: head = [1,2,3,4,5]
//Output: [1,3,5,2,4]
//
//
// Example 2:
//
//
//Input: head = [2,1,3,5,6,4,7]
//Output: [2,3,6,7,1,5,4]
//
//
//
// Constraints:
//
//
// n == number of nodes in the linked list
// 0 <= n <= 10⁴
// -10⁶ <= Node.val <= 10⁶
//
// Related Topics Linked List 👍 4105 👎 359

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *data.ListNode) *data.ListNode {
	evenDummy, oddDummy := &data.ListNode{}, &data.ListNode{}
	evenPre, oddPre, cur := evenDummy, oddDummy, head
	index := 1
	for cur != nil {
		next := cur.Next
		cur.Next = nil
		// 奇数处理
		if index%2 == 1 {
			evenPre.Next = cur
			evenPre = cur
		} else {
			oddPre.Next = cur
			oddPre = cur
		}
		cur = next
		index++
	}
	evenPre.Next = oddDummy.Next
	return evenDummy.Next
}
