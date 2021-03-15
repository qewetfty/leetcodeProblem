package linkedlist

import "github.com/leetcodeProblem/data"

// You are given the head of a linked list, and an integer k.
// Return the head of the linked list after swapping the values of the kth node from the beginning and the kth node from the end (the list is 1-indexed).
//	Example 1:
//		Input: head = [1,2,3,4,5], k = 2
//		Output: [1,4,3,2,5]
//	Example 2:
//		Input: head = [7,9,6,6,7,8,3,0,9,5], k = 5
//		Output: [7,9,6,6,8,7,3,0,9,5]
//	Example 3:
//		Input: head = [1], k = 1
//		Output: [1]
//	Example 4:
//		Input: head = [1,2], k = 1
//		Output: [2,1]
//	Example 5:
//		Input: head = [1,2,3], k = 2
//		Output: [1,2,3]
//	Constraints:
//		The number of nodes in the list is n.
//		1 <= k <= n <= 105
//		0 <= Node.val <= 100

// 直接用array存储，替换数字法
func swapNodes(head *data.ListNode, k int) *data.ListNode {
	nodeList, node := make([]*data.ListNode, 0), head
	for node != nil {
		nodeList = append(nodeList, node)
		node = node.Next
	}
	l := len(nodeList)
	if k > l {
		return head
	}
	nodeList[k-1].Val, nodeList[l-k].Val = nodeList[l-k].Val, nodeList[k-1].Val
	return head
}

// 获取长度强行转换法
func swapNodes2(head *data.ListNode, k int) *data.ListNode {
	l, node := 0, head
	for node != nil {
		l++
		node = node.Next
	}
	if k > l {
		return head
	}
	dummy := &data.ListNode{Next: head}
	node, count := dummy, 0
	var kPrev, kCur, kLastPrev, kLastCur *data.ListNode
	for node != nil {
		if count == k-1 {
			kPrev = node
		}
		if count == l-k {
			kLastPrev = node
		}
		node = node.Next
		count++
		if count == k {
			kCur = node
		}
		if count == l-k+1 {
			kLastCur = node
		}
	}
	kPrev.Next, kLastPrev.Next = kLastCur, kCur
	kCur.Next, kLastCur.Next = kLastCur.Next, kCur.Next
	return dummy.Next
}
