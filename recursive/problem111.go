package recursive

import (
	"github.com/leetcodeProblem/data"
	"math"
)

// Given a binary tree, find its minimum depth.
// The minimum depth is the number of nodes along the shortest path
// from the root node down to the nearest leaf node.
//	Note: A leaf is a node with no children.

func minDepth(root *data.TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	left := minDepth(root.Left)
	right := minDepth(root.Right)
	if root.Left == nil || root.Right == nil {
		return left + right + 1
	}
	return int(math.Min(float64(left), float64(right))) + 1
}
