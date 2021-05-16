package dynamic_programming

import (
	"github.com/leetcodeProblem/data"
	"github.com/leetcodeProblem/utils"
	"math"
)

// Given a binary tree, we install cameras on the nodes of the tree.
// Each camera at a node can monitor its parent, itself, and its immediate children.
// Calculate the minimum number of cameras needed to monitor all nodes of the tree.
//	Example 1:
//		Input: [0,0,null,0,0]
//		Output: 1
//		Explanation: One camera is enough to monitor all nodes if placed as shown.
//	Example 2:
//		Input: [0,0,null,0,null,0,null,null,0]
//		Output: 2
//		Explanation: At least two cameras are needed to monitor all nodes of the tree. The above image shows one of the valid configurations of camera placement.
//	Note:
//		The number of nodes in the given tree will be in the range [1, 1000].
//		Every node has value 0.

// dfs动态规划解法
func minCameraCover(root *data.TreeNode) int {
	_, result, _ := dfs968(root)
	return result
}

// a: node放置摄像头的情况下，覆盖到整棵树需要的摄像头的数目
// b: 覆盖到整棵树需要的摄像头的数目。（不论root是否有摄像头）
// c: 覆盖到node的两个叶子节点时，需要的摄像头的数目（不论root是否有摄像头）。
func dfs968(node *data.TreeNode) (int, int, int) {
	if node == nil {
		return math.MaxInt32 >> 1, 0, 0
	}
	la, lb, lc := dfs968(node.Left)
	ra, rb, rc := dfs968(node.Right)
	a := lc + rc + 1
	b := utils.Min(a, utils.Min(la+rb, lb+ra))
	c := utils.Min(a, lb+rb)
	return a, b, c
}
