package recursive

import (
	"github.com/leetcodeProblem/data"
	"math"
)

// Given a binary tree, find its maximum depth.
// The maximum depth is the number of nodes along the longest path
// from the root node down to the farthest leaf node.
//	Note: A leaf is a node with no children.

// DFS method
func maxDepth(root *data.TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	return int(math.Max(float64(maxDepth(root.Left))+1, float64(maxDepth(root.Right))+1))
}
