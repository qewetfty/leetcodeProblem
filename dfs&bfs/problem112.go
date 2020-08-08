package dfs_bfs

import (
	"fmt"
	"github.com/leetcodeProblem/data"
)

// Given a binary tree and a sum, determine if the tree has a root-to-leaf path such that adding up all the values along
// the path equals the given sum.
//	Note: A leaf is a node with no children.
//	Example:
//		Given the below binary tree and sum = 22,
//
//		      5
//		     / \
//		    4   8
//		   /   / \
//		  11  13  4
//		 /  \      \
//		7    2      1
//		return true, as there exist a root-to-leaf path 5->4->11->2 which sum is 22.

var hasPath = false

func hasPathSum(root *data.TreeNode, sum int) bool {
	hasPath = false
	if root == nil {
		return false
	}
	dfsProblem112(root, 0, sum)
	return hasPath
}

func dfsProblem112(node *data.TreeNode, curSum, sum int) {
	// 判断是否是叶子节点
	nowSum := curSum + node.Val
	if node.Left == nil && node.Right == nil {
		if nowSum == sum {
			hasPath = true
		}
		return
	}
	if node.Left != nil {
		dfsProblem112(node.Left, nowSum, sum)
	}
	if node.Right != nil {
		dfsProblem112(node.Right, nowSum, sum)
	}
}

func testProblem112() {
	a := data.TreeNode{
		Val: 1,
	}
	fmt.Println(hasPathSum(&a, 0))
}
