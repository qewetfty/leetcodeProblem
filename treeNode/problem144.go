package treeNode

import "github.com/leetcodeProblem/data"

// Given a binary tree, return the preorder traversal of its nodes' values.

func preorderTraversal(root *data.TreeNode) []int {
	res := make([]int, 0)
	preOrder(root, &res)
	return res
}

func preOrder(root *data.TreeNode, res *[]int) {
	if root == nil {
		return
	}
	*res = append(*res, root.Val)
	preOrder(root.Left, res)
	preOrder(root.Right, res)
}
