package dfs_bfs

import (
	"fmt"
	"github.com/leetcodeProblem/data"
)

// Given the head of a singly linked list where elements are sorted in ascending
// order, convert it to a height balanced BST.
// For this problem, a height-balanced binary tree is defined as a binary tree in
// which the depth of the two subtrees of every node never differ by more than 1.
//	Example 1:
//		Input: head = [-10,-3,0,5,9]
//		Output: [0,-3,9,-10,null,5]
//		Explanation: One possible answer is [0,-3,9,-10,null,5], which represents the shown height balanced BST.
//	Example 2:
//		Input: head = []
//		Output: []
//	Example 3:
//		Input: head = [0]
//		Output: [0]
//	Example 4:
//		Input: head = [1,3]
//		Output: [3,1]
//	Constraints:
//		The number of nodes in head is in the range [0, 2 * 104].
//		-10^5 <= Node.val <= 10^5

// 数组解法
func sortedListToBST(head *data.ListNode) *data.TreeNode {
	valueList, node := make([]int, 0), head
	for node != nil {
		valueList = append(valueList, node.Val)
		node = node.Next
	}
	return dfs109(valueList)
}

func dfs109(valueList []int) *data.TreeNode {
	l := len(valueList)
	if l == 0 {
		return nil
	}
	if l == 1 {
		return &data.TreeNode{Val: valueList[0]}
	}
	mid := l >> 1
	curNode := &data.TreeNode{Val: valueList[mid]}
	curNode.Left = dfs109(valueList[:mid])
	curNode.Right = dfs109(valueList[mid+1:])
	return curNode
}

var globalHead *data.ListNode

func sortedListToBST2(head *data.ListNode) *data.TreeNode {
	globalHead = head
	l := data.GetNodeLength(head)
	return buildTree(0, l-1)
}

func buildTree(left, right int) *data.TreeNode {
	if left > right {
		return nil
	}
	mid := (left + right + 1) >> 1
	root := new(data.TreeNode)
	root.Left = buildTree(left, mid-1)
	root.Val = globalHead.Val
	globalHead = globalHead.Next
	root.Right = buildTree(mid+1, right)
	return root
}

func testProblem109() {
	node5 := &data.ListNode{Val: 9}
	node4 := &data.ListNode{Val: 5, Next: node5}
	node3 := &data.ListNode{Val: 0, Next: node4}
	node2 := &data.ListNode{Val: -3, Next: node3}
	node1 := &data.ListNode{Val: -10, Next: node2}
	fmt.Println(sortedListToBST(node1))
}
