package treeNode

import (
	"container/list"
	"github.com/leetcodeProblem/data"
)

// Given a binary tree, return the level order traversal of its nodes' values. (ie, from left to right, level by level).
//	For example:
//		Given binary tree [3,9,20,null,null,15,7],
//		    3
//		   / \
//		  9  20
//		    /  \
//		   15   7
//		return its level order traversal as:
//		[
//		  [3],
//		  [9,20],
//		  [15,7]
//		]

// BFS method
func levelOrderBFS(root *data.TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	deque := new(list.List)
	deque.PushBack(root)
	for deque.Len() > 0 {
		level := make([]int, 0)
		length := deque.Len()
		for i := 0; i < length; i++ {
			treeNode := deque.Remove(deque.Front()).(*data.TreeNode)
			level = append(level, treeNode.Val)
			if treeNode.Left != nil {
				deque.PushBack(treeNode.Left)
			}
			if treeNode.Right != nil {
				deque.PushBack(treeNode.Right)
			}
		}
		res = append(res, level)
	}
	return res
}

// DFS method
func levelOrderDFS(root *data.TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	dfs(root, 1, &res)
	return res
}

func dfs(root *data.TreeNode, level int, res *[][]int) {
	if root == nil {
		return
	}
	if len(*res) < level {
		*res = append(*res, make([]int, 0))
	}
	(*res)[level-1] = append((*res)[level-1], root.Val)
	dfs(root.Left, level+1, res)
	dfs(root.Right, level+1, res)
}
