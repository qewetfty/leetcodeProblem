package linkedlist

import "github.com/leetcodeProblem/data"

// You are given an array of k linked-lists lists, each linked-list is sorted in ascending order.
// Merge all the linked-lists into one sorted linked-list and return it.
//	Example 1:
//		Input: lists = [[1,4,5],[1,3,4],[2,6]]
//		Output: [1,1,2,3,4,4,5,6]
//		Explanation: The linked-lists are:
//			[
//			  1->4->5,
//			  1->3->4,
//			  2->6
//			]
//			merging them into one sorted list:
//			1->1->2->3->4->4->5->6
//	Example 2:
//		Input: lists = []
//		Output: []
//	Example 3:
//		Input: lists = [[]]
//		Output: []
//	Constraints:
//		k == lists.length
//		0 <= k <= 10^4
//		0 <= lists[i].length <= 500
//		-10^4 <= lists[i][j] <= 10^4
//		lists[i] is sorted in ascending order.
//		The sum of lists[i].length won't exceed 10^4.

// 分治法，思路类似于归并排序的方法，对链表进行两两合并
func mergeKLists(lists []*data.ListNode) *data.ListNode {
	l := len(lists)
	if l == 0 {
		return nil
	}
	if l == 1 {
		return lists[0]
	}
	// 两个链表进行合并，然后将新的链表头放在第一个链表在列表中的位置。最后合并完之后，lists[0]就是我们想要的结果
	interval := 1
	for interval < l {
		for i := 0; i < l; i = i + 2*interval {
			node1 := lists[i]
			var node2 *data.ListNode
			if i+interval < l {
				node2 = lists[i+interval]
			}
			lists[i] = merge2List(node1, node2)
		}
		interval *= 2
	}
	return lists[0]
}

// 合并两个链表的方法
func merge2List(node1, node2 *data.ListNode) *data.ListNode {
	if node1 == nil {
		return node2
	}
	if node2 == nil {
		return node1
	}
	dummy := &data.ListNode{Val: 0}
	tail, a, b := dummy, node1, node2
	for a != nil && b != nil {
		if a.Val <= b.Val {
			tail.Next = a
			a = a.Next
		} else {
			tail.Next = b
			b = b.Next
		}
		tail = tail.Next
	}
	if a == nil {
		tail.Next = b
	} else {
		tail.Next = a
	}
	return dummy.Next
}
