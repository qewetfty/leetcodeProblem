package linkedlist

import (
	"fmt"
	"github.com/leetcodeProblem/data"
)

// Sort a linked list using insertion sort.
// A graphical example of insertion sort. The partial sorted list (black)
// initially contains only the first element in the list. With each iteration one
// element (red) is removed from the input data and inserted in-place into the
// sorted list
// Algorithm of Insertion Sort:
//	Insertion sort iterates, consuming one input element each repetition, and growing a sorted output list.
//	At each iteration, insertion sort removes one element from the input data, finds the location it belongs
//	within the sorted list, and inserts it there.
//	It repeats until no input elements remain.
//	Example 1:
//		Input: 4->2->1->3
//		Output: 1->2->3->4
//	Example 2:
//		Input: -1->5->3->4->0
//		Output: -1->0->3->4->5

// 简便方法，比之前方法较易理解
func insertionSortList(head *data.ListNode) *data.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	pre, cur, dummy := head, head.Next, &data.ListNode{Next: head}
	for cur != nil {
		// 如果已经有序了，就不需要排序了，直接找下一个
		if pre.Val < cur.Val {
			pre, cur = pre.Next, cur.Next
		} else {
			// 如果无序，则需找到要插入的位置
			p := dummy
			// cur节点插入在p和p.Next之间
			for p.Next != cur && p.Next.Val < cur.Val {
				p = p.Next
			}
			// 删除当前节点
			pre.Next = cur.Next
			// 连接当前节点
			p.Next, cur.Next = cur, p.Next
			cur = pre.Next
		}
	}
	return dummy.Next
}

func insertionSortList2(head *data.ListNode) *data.ListNode {
	if head == nil {
		return nil
	}
	dummy := &data.ListNode{Next: head}
	prev, cur, next := dummy, dummy, dummy.Next
	for next != nil {
		prev, cur, next = cur, next, next.Next
		// 从头到尾进行遍历，找到要插入的位置
		nodePrev, node := dummy, dummy.Next
		for ; node != cur; node, nodePrev = node.Next, nodePrev.Next {
			// 如果cur的值比当前node值要小，则找到了要插入的节点node，直接break
			if cur.Val < node.Val {
				break
			}
		}
		// 如果位置是当前节点，则不动
		if node == cur {
			continue
		}
		// 删除当前节点
		prev.Next, cur.Next = next, nil
		// 当前节点连接到node之前
		nodePrev.Next, cur.Next, cur = cur, node, prev
	}
	return dummy.Next
}

func testProblem147() {
	listnode4 := &data.ListNode{Val: 3}
	listnode3 := &data.ListNode{Val: 1, Next: listnode4}
	listnode2 := &data.ListNode{Val: 2, Next: listnode3}
	listnode1 := &data.ListNode{Val: 4, Next: listnode2}
	list := insertionSortList(listnode1)
	fmt.Println(list.ToString())
}
