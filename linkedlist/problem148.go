package linkedlist

import "github.com/leetcodeProblem/data"

// Given the head of a linked list, return the list after sorting it in ascending order.
// Follow up: Can you sort the linked list in O(n logn) time and O(1) memory (i.e. constant space)?
//	Example 1:
//		Input: head = [4,2,1,3]
//		Output: [1,2,3,4]
//	Example 2:
//		Input: head = [-1,5,3,4,0]
//		Output: [-1,0,3,4,5]
//	Example 3:
//		Input: head = []
//		Output: []
//	Constraints:
//		The number of nodes in the list is in the range [0, 5 * 104].
//		-105 <= Node.val <= 105

// O(nlog(n))时间复杂度，O(1)空间复杂度，不使用递归
func sortList(head *data.ListNode) *data.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// 获取链表长度
	h, l := head, 0
	for h != nil {
		h, l = h.Next, l+1
	}
	// 根据interval来计算哪两个位置要合并，dummy记录head之前和head的数据，方便之后获取
	interval, dummy := 1, &data.ListNode{Next: head}
	for interval < l {
		pre, h := dummy, dummy.Next
		for h != nil {
			// 获取第一个要合并的head
			h1, i := h, interval
			for i > 0 && h != nil {
				h, i = h.Next, i-1
			}
			// 如果第一个合并的head小于interval，直接跳过
			if i > 0 || h == nil {
				break
			}
			// 获取第二个要合并的head
			h2, i := h, interval
			for i > 0 && h != nil {
				h, i = h.Next, i-1
			}
			// 获取两个要合并链表的长度
			c1, c2 := interval, interval-i
			// 开始合并
			for c1 > 0 && c2 > 0 {
				if h1.Val < h2.Val {
					pre.Next, h1, c1 = h1, h1.Next, c1-1
				} else {
					pre.Next, h2, c2 = h2, h2.Next, c2-1
				}
				pre = pre.Next
			}
			if c1 != 0 {
				pre.Next = h1
			} else {
				pre.Next = h2
			}
			// 将pre的下一个节点指向h，维持每个interval之间相互指的顺序
			for c1 > 0 || c2 > 0 {
				pre, c1, c2 = pre.Next, c1-1, c2-1
			}
			pre.Next = h
		}
		interval = interval * 2
	}
	return dummy.Next
}

// O(nlog(n))时间复杂度，O(log(n))的空间复杂度，因为使用了递归方法
func sortList2(head *data.ListNode) *data.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	mid := getMid(head)
	left, right := sortList2(head), sortList2(mid)
	return merge148(left, right)
}

func merge148(l1, l2 *data.ListNode) *data.ListNode {
	dummy := &data.ListNode{}
	node := dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			node.Next = l1
			l1 = l1.Next
		} else {
			node.Next = l2
			l2 = l2.Next
		}
		node = node.Next
	}
	if l1 != nil {
		node.Next = l1
	} else {
		node.Next = l2
	}
	return dummy.Next
}

func getMid(head *data.ListNode) *data.ListNode {
	var midPre *data.ListNode
	midPre = nil
	for head != nil && head.Next != nil {
		if midPre == nil {
			midPre = head
		} else {
			midPre = midPre.Next
		}
		head = head.Next.Next
	}
	mid := midPre.Next
	midPre.Next = nil
	return mid
}
