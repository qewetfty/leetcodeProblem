package dfs_bfs

import (
	"container/list"
	"github.com/leetcodeProblem/data"
	"math"
)

// You need to find the largest value in each row of a binary tree.
//	Example:
//	Input:
//
//          1
//         / \
//        3   2
//       / \   \
//      5   3   9
//
//	Output: [1, 3, 9]

// BFS method
func largestValues(root *data.TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	deque := new(list.List)
	deque.PushBack(root)
	for deque.Len() > 0 {
		maxNumber := math.MinInt32
		length := deque.Len()
		for i := 0; i < length; i++ {
			treeNode := deque.Remove(deque.Front()).(*data.TreeNode)
			maxNumber = int(math.Max(float64(maxNumber), float64(treeNode.Val)))
			if treeNode.Left != nil {
				deque.PushBack(treeNode.Left)
			}
			if treeNode.Right != nil {
				deque.PushBack(treeNode.Right)
			}
		}
		res = append(res, maxNumber)
	}
	return res
}

// DFS method
func largestValues2(root *data.TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	dfs515(root, 1, &res)
	return res
}

func dfs515(node *data.TreeNode, level int, res *[]int) {
	if len(*res) < level {
		*res = append(*res, math.MinInt32)
	}
	(*res)[level-1] = int(math.Max(float64((*res)[level-1]), float64(node.Val)))
	if node.Left != nil {
		dfs515(node.Left, level+1, res)
	}
	if node.Right != nil {
		dfs515(node.Right, level+1, res)
	}
}
