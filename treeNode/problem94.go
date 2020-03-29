package treeNode

import "github.com/leetcodeProblem/data"

// Given a binary tree, return the inorder traversal of its nodes' values.

func inorderTraversal(root *data.TreeNode) []int {
	res := make([]int, 0)
	inorder(root, &res)
	return res
}

func inorder(root *data.TreeNode, res *[]int) {
	if root == nil {
		return
	}
	inorder(root.Left, res)
	*res = append(*res, root.Val)
	inorder(root.Right, res)
}
