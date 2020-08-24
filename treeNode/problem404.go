package treeNode

import "github.com/leetcodeProblem/data"

// Find the sum of all left leaves in a given binary tree.
//	Example:
//	  	  3
//	  	 / \
//	  	9  20
//	  	  /  \
//	  	 15   7
//	There are two left leaves in the binary tree, with values 9 and 15 respectively. Return 24.

var sum int

func sumOfLeftLeaves(root *data.TreeNode) int {
	sum = 0
	if root == nil {
		return sum
	}
	dfs404(root, false)
	return sum
}

func dfs404(node *data.TreeNode, isLeft bool) {
	if node == nil {
		return
	}
	if isLeft && node.Left == nil && node.Right == nil {
		sum = sum + node.Val
		return
	}
	dfs404(node.Left, true)
	dfs404(node.Right, false)
}
