package treeNode

import "github.com/leetcodeProblem/data"

// Given a binary search tree, rearrange the tree in in-order so that the leftmost
// node in the tree is now the root of the tree, and every node has no left child and only 1 right child.

func increasingBST(root *data.TreeNode) *data.TreeNode {
	if root == nil {
		return nil
	}
	nums := make([]int, 0)
	nums = readMid(root, &nums)
	newRoot := new(data.TreeNode)
	newRoot.Val = nums[0]
	head := newRoot
	for i, num := range nums {
		if i == 0 {
			continue
		}
		newTreeNode := new(data.TreeNode)
		newTreeNode.Val = num
		head.Right = newTreeNode
		head = newTreeNode
	}
	return newRoot
}

func readMid(root *data.TreeNode, nums *[]int) []int {
	if root.Left != nil {
		readMid(root.Left, nums)
	}
	*nums = append(*nums, root.Val)
	if root.Right != nil {
		readMid(root.Right, nums)
	}
	return *nums
}
