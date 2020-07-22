package dfs_bfs

import (
	list2 "container/list"
	"github.com/leetcodeProblem/data"
)

// Given a binary tree, return the zigzag level order traversal of its nodes' values.
// (ie, from left to right, then right to left for the next level and alternate between).
//	For example:
//		Given binary tree [3,9,20,null,null,15,7],
//		  	  3
//		  	 / \
//		  	9  20
//		  	  /  \
//		  	 15   7
//		return its zigzag level order traversal as:
//			[
//			  [3],
//			  [20,9],
//			  [15,7]
//			]

func zigzagLevelOrder(root *data.TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	list := new(list2.List)
	list.PushBack(root)
	level := 1
	for list.Len() > 0 {
		listLen := list.Len()
		thisLevelRes := make([]int, 0)
		for i := 0; i < listLen; i++ {
			if level%2 == 1 {
				node := list.Remove(list.Front()).(*data.TreeNode)
				thisLevelRes = append(thisLevelRes, node.Val)
				if node.Left != nil {
					list.PushBack(node.Left)
				}
				if node.Right != nil {
					list.PushBack(node.Right)
				}
			} else {
				node := list.Remove(list.Back()).(*data.TreeNode)
				thisLevelRes = append(thisLevelRes, node.Val)
				if node.Right != nil {
					list.PushFront(node.Right)
				}
				if node.Left != nil {
					list.PushFront(node.Left)
				}
			}
		}
		res = append(res, thisLevelRes)
		level++
	}
	return res
}
