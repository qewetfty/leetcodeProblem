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

// prefix sum method
var prefixMap map[int]int

func pathSumPrefix(root *data.TreeNode, sum int) int {
	prefixMap = make(map[int]int)
	prefixMap[0] = 1
	return helper(root, sum, 0)
}

func helper(node *data.TreeNode, target, curSum int) int {
	if node == nil {
		return 0
	}
	res := 0
	curSum = curSum + node.Val
	remainNumber := 0
	if _, ok := prefixMap[curSum-target]; ok {
		remainNumber = prefixMap[curSum-target]
	}
	res = res + remainNumber
	curSumPrefixNumber := 0
	if _, ok := prefixMap[curSum]; ok {
		curSumPrefixNumber = prefixMap[curSum]
	}
	curSumPrefixNumber = curSumPrefixNumber + 1
	prefixMap[curSum] = curSumPrefixNumber
	res = res + helper(node.Left, target, curSum)
	res = res + helper(node.Right, target, curSum)
	prefixMap[curSum]--
	return res
}
