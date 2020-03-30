package recursive

import "github.com/leetcodeProblem/data"

// Invert a binary tree.
func invertTree(root *data.TreeNode) *data.TreeNode {
	invert(root)
	return root
}

func invert(root *data.TreeNode) {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return
	}
	invert(root.Left)
	invert(root.Right)
	left := root.Left
	right := root.Right
	root.Left = right
	root.Right = left
}
