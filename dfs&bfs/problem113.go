package dfs_bfs

import (
	"fmt"
	"github.com/leetcodeProblem/data"
)

// Given a binary tree and a sum, find all root-to-leaf paths where each path's sum equals the given sum.
//	Note: A leaf is a node with no children.
//	Example:
//		Given the below binary tree and sum = 22,
//
//		      5
//		     / \
//		    4   8
//		   /   / \
//		  11  13  4
//		 /  \    / \
//		7    2  5   1
//	Return:
//		[
//		   [5,4,11,2],
//		   [5,8,4,5]
//		]

var res [][]int

func pathSum(root *data.TreeNode, sum int) [][]int {
	res = make([][]int, 0)
	if root == nil {
		return res
	}
	curNumList := make([]int, 0)
	dfsProblem113(root, 0, sum, curNumList)
	return res
}

func dfsProblem113(node *data.TreeNode, curSum, sum int, curNumList []int) {
	nowSum := curSum + node.Val
	curNumList = append(curNumList, node.Val)
	if node.Left == nil && node.Right == nil {
		if nowSum == sum {
			newResList := make([]int, 0)
			newResList = append(newResList, curNumList...)
			res = append(res, newResList)
		}
		curNumList = curNumList[:len(curNumList)-1]
		return
	}
	if node.Left != nil {
		dfsProblem113(node.Left, nowSum, sum, curNumList)
	}
	if node.Right != nil {
		dfsProblem113(node.Right, nowSum, sum, curNumList)
	}
	curNumList = curNumList[:len(curNumList)-1]
}

func testProblem113() {
	a := data.TreeNode{
		Val: 1,
	}
	fmt.Println(pathSum(&a, 0))
}
