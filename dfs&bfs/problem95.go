package dfs_bfs

import "github.com/leetcodeProblem/data"

// Given an integer n, return all the structurally unique BST's (binary search
// trees), which has exactly n nodes of unique values from 1 to n. Return the
// answer in any order.
//	Example 1:
//		Input: n = 3
//		Output: [[1,null,2,null,3],[1,null,3,2],[2,1,3],[3,1,null,null,2],[3,2,null,1]]
//	Example 2:
//		Input: n = 1
//		Output: [[1]]
//	Constraints:
//		1 <= n <= 8

func generateTrees(n int) []*data.TreeNode {
	if n == 0 {
		return nil
	}
	return helper95(1, n)
}

func helper95(start, end int) []*data.TreeNode {
	if start > end {
		return []*data.TreeNode{nil}
	}
	var allTrees []*data.TreeNode
	for i := start; i <= end; i++ {
		leftTrees := helper95(start, i-1)
		rightTrees := helper95(i+1, end)
		for _, leftTree := range leftTrees {
			for _, rightTree := range rightTrees {
				curTree := &data.TreeNode{Val: i, Left: leftTree, Right: rightTree}
				allTrees = append(allTrees, curTree)
			}
		}
	}
	return allTrees
}
