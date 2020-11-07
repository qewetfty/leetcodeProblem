package linkedlist

import (
	"fmt"
	"github.com/leetcodeProblem/data"
)

// You are given two non-empty linked lists representing two non-negative
// integers. The most significant digit comes first and each of their nodes
// contain a single digit. Add the two numbers and return it as a linked list.
// You may assume the two numbers do not contain any leading zero, except the number 0 itself.
//	Follow up:
//		What if you cannot modify the input lists? In other words, reversing the lists is not allowed.
//	Example:
//		Input: (7 -> 2 -> 4 -> 3) + (5 -> 6 -> 4)
//		Output: 7 -> 8 -> 0 -> 7

// 使用栈，从后向前进行加法计算
func addTwoNumbers445(l1 *data.ListNode, l2 *data.ListNode) *data.ListNode {
	stack1, stack2 := make([]*data.ListNode, 0), make([]*data.ListNode, 0)
	for l1 != nil {
		stack1 = append(stack1, l1)
		l1 = l1.Next
	}
	for l2 != nil {
		stack2 = append(stack2, l2)
		l2 = l2.Next
	}
	carry := 0
	var res *data.ListNode
	for len(stack1) > 0 || len(stack2) > 0 || carry > 0 {
		num1, num2 := 0, 0
		if len(stack1) > 0 {
			num1 = stack1[len(stack1)-1].Val
			stack1 = stack1[:len(stack1)-1]
		}
		if len(stack2) > 0 {
			num2 = stack2[len(stack2)-1].Val
			stack2 = stack2[:len(stack2)-1]
		}
		val := num1 + num2 + carry
		carry, val = val/10, val%10
		curNode := &data.ListNode{Val: val, Next: res}
		res = curNode
	}
	return res
}

// 手动补齐0，按位相加
func addTwoNumbers2(l1 *data.ListNode, l2 *data.ListNode) *data.ListNode {
	// 计算两个链表长度
	node1, node1Len, node2, node2Len := l1, 0, l2, 0
	for node1 != nil {
		node1, node1Len = node1.Next, node1Len+1
	}
	for node2 != nil {
		node2, node2Len = node2.Next, node2Len+1
	}
	// 给短的链表前面加0，强制让l1为长的，l2为短的
	if node1Len < node2Len {
		l1, l2, node1Len, node2Len = l2, l1, node2Len, node1Len
	}
	diff := node1Len - node2Len
	dummy1, dummy2 := &data.ListNode{Next: l1}, new(data.ListNode)
	node2 = dummy2
	for diff > 0 {
		node2.Next = &data.ListNode{Val: 0}
		node2, diff = node2.Next, diff-1
	}
	node2.Next = l2
	// 从低位向高位累加
	addValue := add(dummy1.Next, dummy2.Next)
	if addValue > 0 {
		dummy1.Val = addValue
		return dummy1
	} else {
		return dummy1.Next
	}
}

func add(l1, l2 *data.ListNode) int {
	if l1.Next == nil && l2.Next == nil {
		l1.Val = l1.Val + l2.Val
		if l1.Val >= 10 {
			l1.Val -= 10
			return 1
		} else {
			return 0
		}
	}
	addNumber := add(l1.Next, l2.Next)
	l1.Val += addNumber + l2.Val
	if l1.Val >= 10 {
		l1.Val -= 10
		return 1
	} else {
		return 0
	}
}

func testProblem445() {
	listnode4 := &data.ListNode{Val: 3}
	listnode3 := &data.ListNode{Val: 4, Next: listnode4}
	listnode2 := &data.ListNode{Val: 2, Next: listnode3}
	listnode1 := &data.ListNode{Val: 7, Next: listnode2}

	listnode7 := &data.ListNode{Val: 4}
	listnode6 := &data.ListNode{Val: 6, Next: listnode7}
	listnode5 := &data.ListNode{Val: 5, Next: listnode6}

	numbers := addTwoNumbers445(listnode1, listnode5)
	fmt.Println(numbers.ToString())
}
