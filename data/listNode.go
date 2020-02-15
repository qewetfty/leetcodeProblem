package data

import "strconv"

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(val int) *ListNode {
	l := new(ListNode)
	l.Val = val
	return l
}

func (l *ListNode) ToString() string {
	head := l
	resultStr := ""
	if head == nil {
		return resultStr
	}
	resultStr = strconv.Itoa(head.Val)
	for head.Next != nil {
		head = head.Next
		resultStr = resultStr + "->" + strconv.Itoa(head.Val)
	}
	return resultStr
}
