package linkedlist

import (
	"fmt"
	"github.com/leetcodeProblem/data"
)

// Given a singly linked list L: L0→L1→…→Ln-1→Ln,
// reorder it to: L0→Ln→L1→Ln-1→L2→Ln-2→…
// You may not modify the values in the list's nodes, only nodes itself may be changed.
//	Example 1:
//		Given 1->2->3->4, reorder it to 1->4->2->3.
//	Example 2:
//		Given 1->2->3->4->5, reorder it to 1->5->2->4->3.

func reorderList(head *data.ListNode) {
	if head == nil {
		return
	}
	// 使用一个列表保存所有的节点，这样就可以知道每个节点的next是谁
	nodeList, root := make([]*data.ListNode, 0), head
	for root != nil {
		nodeList = append(nodeList, root)
		root = root.Next
	}
	l, i := len(nodeList), 0
	// 分奇数偶数情况，偶数的话中间两个不用移动，奇数则中间一个不移动
	for (l%2 == 0 && i < (l-1)/2) || (l%2 == 1 && i < l/2) {
		nodeList[i].Next = nodeList[l-i-1]
		nodeList[l-i-1].Next = nodeList[i+1]
		// 最后一个已经移动走了，把倒数第二个的Next置为空
		nodeList[l-i-2].Next = nil
		i++
	}
}

func testProblem143() {
	a := data.NewListNode(1)
	b := data.NewListNode(2)
	c := data.NewListNode(3)
	d := data.NewListNode(4)
	a.Next = b
	b.Next = c
	c.Next = d
	reorderList(a)
	fmt.Println(a.ToString())
}
