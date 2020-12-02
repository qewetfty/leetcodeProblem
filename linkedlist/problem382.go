package linkedlist

import (
	"github.com/leetcodeProblem/data"
	"math/rand"
)

// Given a singly linked list, return a random node's value from the linked list.
// Each node must have the same probability of being chosen.
//	Follow up:
//		What if the linked list is extremely large and its length is unknown to you?
//		Could you solve this efficiently without using extra space?
//	Example:
//		// Init a singly linked list [1,2,3].
//		ListNode head = new ListNode(1);
//		head.next = new ListNode(2);
//		head.next.next = new ListNode(3);
//		Solution solution = new Solution(head);
//		// getRandom() should return either 1, 2, or 3 randomly. Each element should have equal probability of returning.
//		solution.getRandom();

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type Solution struct {
	node *data.ListNode
}

/** @param head The linked list's head.
  Note that the head is guaranteed to be not null, so it contains at least one node. */
func Constructor382(head *data.ListNode) Solution {
	return Solution{
		node: head,
	}
}

/** Returns a random node's value. */
func (this *Solution) GetRandom() int {
	node := this.node
	if node == nil {
		return -1
	}
	// 蓄水池算法，获取范围内的随机数。如果随机数是0，则替换蓄水池中的数
	res, i := node.Val, 0
	for node != nil {
		i++
		if rand.Intn(i) == 0 {
			res = node.Val
		}
		node = node.Next
	}
	return res
}
