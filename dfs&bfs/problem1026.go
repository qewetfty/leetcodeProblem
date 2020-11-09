package dfs_bfs

import (
	"github.com/leetcodeProblem/data"
	"github.com/leetcodeProblem/utils"
	"math"
)

// Given the root of a binary tree, find the maximum value V for which there
// exist different nodes A and B where V = |A.val - B.val| and A is an ancestor
// of B.
// A node A is an ancestor of B if either: any child of A is equal to B, or any child of A is an ancestor of B.
//	Example 1:
//		Input: root = [8,3,10,1,6,null,14,null,null,4,7,13]
//		Output: 7
//		Explanation: We have various ancestor-node differences, some of which are given below :
//			|8 - 3| = 5
//			|3 - 7| = 4
//			|8 - 1| = 7
//			|10 - 13| = 3
//			Among all possible differences, the maximum value of 7 is obtained by |8 - 1| = 7.
//	Example 2:
//		Input: root = [1,null,2,null,0,3]
//		Output: 3
//	Constraints:
//		The number of nodes in the tree is in the range [2, 5000].
//		0 <= Node.val <= 105

var res1026 int

// 自上而下记录路径上的最大最小值，当到达叶子节点时候计算最大最小值的差值
func maxAncestorDiff(root *data.TreeNode) int {
	res1026 = 0
	helper1026(root, root.Val, root.Val)
	return res1026
}

func helper1026(node *data.TreeNode, max, min int) {
	if node == nil {
		return
	}
	max, min = utils.Max(max, node.Val), utils.Min(min, node.Val)
	// 到达叶子节点，记录路径上最大最小值差值，然后更新result
	if node.Left == nil && node.Right == nil {
		res1026 = utils.Max(res1026, int(math.Abs(float64(max-min))))
	}
	helper1026(node.Left, max, min)
	helper1026(node.Right, max, min)
}

// 自下而上记录所有叶子节点的最大最小值，在当前层计算
func maxAncestorDiff2(root *data.TreeNode) int {
	res1026 = 0
	dfs1026(root)
	return res1026
}

func dfs1026(node *data.TreeNode) (int, int) {
	// 叶子节点的最大最小值为叶子节点本身
	if node.Left == nil && node.Right == nil {
		return node.Val, node.Val
	}
	leftMax, leftMin, rightMax, rightMin := node.Val, node.Val, node.Val, node.Val
	// 获取左节点的最大最小值
	if node.Left != nil {
		newLeftMax, newLeftMin := dfs1026(node.Left)
		leftMax, leftMin = utils.Max(leftMax, newLeftMax), utils.Min(leftMin, newLeftMin)
	}
	// 获取右节点的最大最小值
	if node.Right != nil {
		newRightMax, newRightMin := dfs1026(node.Right)
		rightMax, rightMin = utils.Max(rightMax, newRightMax), utils.Min(rightMin, newRightMin)
	}
	// 获取当前节点的最大最小值并计算
	curMax, curMin := utils.Max(leftMax, rightMax), utils.Min(leftMin, rightMin)
	res1026 = utils.Max(utils.Max(int(math.Abs(float64(node.Val-curMax))), int(math.Abs(float64(node.Val-curMin)))), res1026)
	return curMax, curMin
}
