package linkedlist

import (
	"fmt"
	"github.com/leetcodeProblem/data"
)

// Given a linked list, rotate the list to the right by k places, where k is non-negative.
//	Example 1:
//		Input: 1->2->3->4->5->NULL, k = 2
//		Output: 4->5->1->2->3->NULL
//		Explanation:
//			rotate 1 steps to the right: 5->1->2->3->4->NULL
//			rotate 2 steps to the right: 4->5->1->2->3->NULL
//	Example 2:
//		Input: 0->1->2->NULL, k = 4
//		Output: 2->0->1->NULL
//		Explanation:
//			rotate 1 steps to the right: 2->0->1->NULL
//			rotate 2 steps to the right: 1->2->0->NULL
//			rotate 3 steps to the right: 0->1->2->NULL
//			rotate 4 steps to the right: 2->0->1->NULL

func rotateRight(head *data.ListNode, k int) *data.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	l, dummy, node, tail := 0, &data.ListNode{Next: head}, head, head
	// 获取链表长度
	for node != nil {
		tail, l, node = node, l+1, node.Next
	}
	// 获取偏移长度
	movingStep, node, i := k%l, dummy, 0
	if movingStep == 0 {
		return head
	}
	// 获取要偏移的节点
	for i < l-movingStep {
		i, node = i+1, node.Next
	}
	// 移动节点
	dummy.Next, tail.Next, node.Next = node.Next, head, nil
	return dummy.Next
}

func testProblem61() {
	node5 := &data.ListNode{Val: 5}
	node4 := &data.ListNode{Val: 4, Next: node5}
	node3 := &data.ListNode{Val: 3, Next: node4}
	node2 := &data.ListNode{Val: 2, Next: node3}
	node1 := &data.ListNode{Val: 1, Next: node2}
	right := rotateRight(node1, 2)
	fmt.Println(right.ToString())
}
