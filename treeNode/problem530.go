package treeNode

import (
	"math"
)

// 给定一个所有节点为非负值的二叉搜索树，求树中任意两节点的差的绝对值的最小值。
// 注意: 树中至少有2个节点。

func getMinimumDifference(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := make([]int, 0)
	searchTreeNode(root, &res)
	min := math.MaxInt32
	for i := 0; i < len(res)-1; i++ {
		min = minNumber(min, res[i+1]-res[i])
	}
	return min
}

func searchTreeNode(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	searchTreeNode(root.Left, res)
	*res = append(*res, root.Val)
	searchTreeNode(root.Right, res)
}

func minNumber(a, b int) int {
	if a < b {
		return a
	}
	return b
}
