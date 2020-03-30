package recursive

import (
	"fmt"
	"github.com/leetcodeProblem/data"
	"math"
)

// Given a binary tree, determine if it is a valid binary search tree (BST).
// Assume a BST is defined as follows:
//	The left subtree of a node contains only nodes with keys less than the node's key.
//	The right subtree of a node contains only nodes with keys greater than the node's key.
//	Both the left and right subtrees must also be binary search trees.

func isValidBST(root *data.TreeNode) bool {
	return isValid(root, math.MinInt64, math.MaxInt64)
}

func isValid(root *data.TreeNode, leftBound int, rightBound int) bool {
	if root == nil {
		return true
	}
	if root.Val <= leftBound || root.Val >= rightBound {
		return false
	}
	return isValid(root.Left, leftBound, root.Val) && isValid(root.Right, root.Val, rightBound)
}

func testProblem98() {
	node1 := new(data.TreeNode)
	node1.Val = 2147483647

	fmt.Println(isValidBST(node1))
}
