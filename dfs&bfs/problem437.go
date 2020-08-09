package dfs_bfs

import (
	"container/list"
	"github.com/leetcodeProblem/data"
)

func pathSum437(root *data.TreeNode, sum int) int {
	res := 0
	if root == nil {
		return res
	}
	dfsProblem437(root, &res, sum)
	return res
}

func dfsProblem437(node *data.TreeNode, res *int, sum int) {
	*res = *res + hasSumPath(node, sum)
	if node.Left != nil {
		dfsProblem437(node.Left, res, sum)
	}
	if node.Right != nil {
		dfsProblem437(node.Right, res, sum)
	}
}

type nodeSum struct {
	node *data.TreeNode
	sum  int
}

func hasSumPath(root *data.TreeNode, sum int) int {
	res := 0
	if root == nil {
		return res
	}
	deque := new(list.List)
	deque.PushFront(nodeSum{
		node: root,
		sum:  0,
	})
	for deque.Len() > 0 {
		curLen := deque.Len()
		for i := 0; i < curLen; i++ {
			currNodeSum := deque.Remove(deque.Back()).(nodeSum)
			node, prevSum := currNodeSum.node, currNodeSum.sum
			curSum := prevSum + node.Val
			if curSum == sum {
				res++
			}
			if node.Left != nil {
				deque.PushFront(nodeSum{
					node: node.Left,
					sum:  curSum,
				})
			}
			if node.Right != nil {
				deque.PushFront(nodeSum{
					node: node.Right,
					sum:  curSum,
				})
			}
		}
	}
	return res
}
