package dynamic_programming

import (
	"github.com/leetcodeProblem/data"
	"github.com/leetcodeProblem/utils"
)

// The thief has found himself a new place for his thievery again. There is only
// one entrance to this area, called the "root." Besides the root, each house has
// one and only one parent house. After a tour, the smart thief realized that
// "all houses in this place forms a binary tree". It will automatically contact
// the police if two directly-linked houses were broken into on the same night.
// Determine the maximum amount of money the thief can rob tonight without alerting the police.
//	Example 1:
//		Input: [3,2,3,null,3,null,1]
//
//		     3
//		    / \
//		   2   3
//		    \   \
//		     3   1
//
//		Output: 7
//		Explanation: Maximum amount of money the thief can rob = 3 + 3 + 1 = 7.
//	Example 2:
//		Input: [3,4,5,1,3,null,1]
//
//		     3
//		    / \
//		   4   5
//		  / \   \
//		 1   3   1
//
//		Output: 9
//		Explanation: Maximum amount of money the thief can rob = 4 + 5 = 9.

var dp map[*data.TreeNode]int

func rob337(root *data.TreeNode) int {
	if root == nil {
		return 0
	}
	dp = make(map[*data.TreeNode]int)
	dfs(root)
	return dp[root]
}

func dfs(node *data.TreeNode) {
	if node == nil {
		return
	}
	nextNextLeval := node.Val
	if node.Left != nil {
		dfs(node.Left)
		nextNextLeval = nextNextLeval + dp[node.Left.Left] + dp[node.Left.Right]
	}
	if node.Right != nil {
		dfs(node.Right)
		nextNextLeval = nextNextLeval + dp[node.Right.Left] + dp[node.Right.Right]
	}
	dp[node] = utils.Max(dp[node.Left]+dp[node.Right], nextNextLeval)
}
