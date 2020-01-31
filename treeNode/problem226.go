package treeNode

import "github.com/leetcodeProblem/data"

// Invert a binary tree.

func main() {

}

func invertTree(root *data.TreeNode) *data.TreeNode {
	if root == nil {
		return nil
	}
	if root.Left != nil {
		invertTree(root.Left)
	}
	if root.Right != nil {
		invertTree(root.Right)
	}
	left := root.Left
	right := root.Right
	root.Left = right
	root.Right = left
	return root
}
