package treeNode

import "github.com/leetcodeProblem/data"

//Given the root node of a binary search tree,
// return the sum of values of all nodes with value between L and R (inclusive).
//
//The binary search tree is guaranteed to have unique values.

//Note:
//The number of nodes in the tree is at most 10000.
//The final answer is guaranteed to be less than 2^31.

func rangeSumBST(root *data.TreeNode, L int, R int) int {
	if root == nil {
		return 0
	}
	res := make([]int, 0)
	searchTreeNode(root, &res)
	sum := 0
	for _, i := range res {
		if L <= i && i <= R {
			sum += i
		}
	}
	return sum
}
