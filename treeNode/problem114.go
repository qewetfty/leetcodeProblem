package treeNode

import "github.com/leetcodeProblem/data"

// Given a binary tree, flatten it to a linked list in-place.

func flatten(root *data.TreeNode) {
	if root == nil {
		return
	}
	if root.Left == nil {
		flatten(root.Right)
		return
	}
	if root.Right == nil {
		root.Right = root.Left
		root.Left = nil
		flatten(root.Right)
		return
	}
	flatten(root.Left)
	flatten(root.Right)
	node := root.Left
	for node.Right != nil {
		node = node.Right
	}
	node.Right = root.Right
	root.Right = root.Left
	root.Left = nil
}
