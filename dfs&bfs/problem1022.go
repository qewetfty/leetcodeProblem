package dfs_bfs

import (
	"container/list"
	"github.com/leetcodeProblem/data"
)

// Given a binary tree, each node has value 0 or 1.  Each root-to-leaf path represents a binary number starting with
// the most significant bit.  For example, if the path is 0 -> 1 -> 1 -> 0 -> 1, then this could represent 01101 in
// binary, which is 13.
// For all leaves in the tree, consider the numbers represented by the path from the root to that leaf.
// Return the sum of these numbers.
//	Example 1:
//		Input: [1,0,1,0,1,0,1]
//		Output: 22
//		Explanation: (100) + (101) + (110) + (111) = 4 + 5 + 6 + 7 = 22
//	Note:
//		The number of nodes in the tree is between 1 and 1000.
//		node.val is 0 or 1.
//		The answer will not exceed 2^31 - 1.

// bfs method
func sumRootToLeafBfs(root *data.TreeNode) int {
	if root == nil {
		return 0
	}
	deque, sum := new(list.List), 0
	deque.PushFront(rootToLeaf{node: root, curNumber: 0})
	for deque.Len() > 0 {
		l := deque.Len()
		for i := 0; i < l; i++ {
			curRootToLeaf := deque.Remove(deque.Back()).(rootToLeaf)
			node, num := curRootToLeaf.node, curRootToLeaf.curNumber
			newNum := num*2 + node.Val
			if node.Left == nil && node.Right == nil {
				sum = sum + newNum
				continue
			}
			if node.Left != nil {
				deque.PushFront(rootToLeaf{node: node.Left, curNumber: newNum})
			}
			if node.Right != nil {
				deque.PushFront(rootToLeaf{node: node.Right, curNumber: newNum})
			}
		}
	}
	return sum
}

// dfs1022 method
func sumRootToLeafDfs(root *data.TreeNode) int {
	if root == nil {
		return 0
	}
	return dfs1022(root, 0)
}

func dfs1022(node *data.TreeNode, value int) int {
	currentValue := (value << 1) + node.Val
	if node.Left == nil && node.Right == nil {
		return currentValue
	}
	sum := 0
	if node.Left != nil {
		sum = sum + dfs1022(node.Left, currentValue)
	}
	if node.Right != nil {
		sum = sum + dfs1022(node.Right, currentValue)
	}
	return sum
}

type rootToLeaf struct {
	node      *data.TreeNode
	curNumber int
}
