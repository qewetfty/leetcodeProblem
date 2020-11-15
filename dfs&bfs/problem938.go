package dfs_bfs

import "github.com/leetcodeProblem/data"

// Given the root node of a binary search tree, return the sum of values of all
// nodes with a value in the range [low, high].
//	Example 1:
//		Input: root = [10,5,15,3,7,null,18], low = 7, high = 15
//		Output: 32
//	Example 2:
//		Input: root = [10,5,15,3,7,13,18,1,null,6], low = 6, high = 10
//		Output: 23
//	Constraints:
//		The number of nodes in the tree is in the range [1, 2 * 104].
//		1 <= Node.val <= 105
//		1 <= low <= high <= 105
//		All Node.val are unique.

var result938 int

func rangeSumBST(root *data.TreeNode, low int, high int) int {
	result938 = 0
	dfs938(root, low, high)
	return result938
}

func dfs938(node *data.TreeNode, low, high int) {
	if node == nil {
		return
	}
	if node.Val >= low && node.Val <= high {
		result938 += node.Val
	}
	dfs938(node.Left, low, high)
	dfs938(node.Right, low, high)
}
