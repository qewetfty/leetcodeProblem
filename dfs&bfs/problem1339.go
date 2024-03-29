package dfs_bfs

import (
	"fmt"
	"github.com/leetcodeProblem/data"
	"github.com/leetcodeProblem/utils"
)

// Given the root of a binary tree, split the binary tree into two subtrees by
// removing one edge such that the product of the sums of the subtrees is
// maximized.
// Return the maximum product of the sums of the two subtrees. Since the answer may be too large, return it modulo 109 + 7.
// Note that you need to maximize the answer before taking the mod and not after taking it.
//	Example 1:
//		Input: root = [1,2,3,4,5,6]
//		Output: 110
//		Explanation: Remove the red edge and get 2 binary trees with sum 11 and 10. Their product is 110 (11*10)
//	Example 2:
//		Input: root = [1,null,2,3,4,null,null,5,6]
//		Output: 90
//		Explanation: Remove the red edge and get 2 binary trees with sum 15 and 6.Their product is 90 (15*6)
//	Example 3:
//		Input: root = [2,3,9,10,7,8,6,5,4,11,1]
//		Output: 1025
//	Example 4:
//		Input: root = [1,1]
//		Output: 1
//	Constraints:
//		The number of nodes in the tree is in the range [2, 5 * 104].
//		1 <= Node.val <= 104

var result1339 int

func maxProduct(root *data.TreeNode) int {
	result1339 = 0
	sum := getSumOfTree(root)
	getMaxProduct(root, sum)
	return result1339 % 1000000007
}

func getSumOfTree(root *data.TreeNode) int {
	if root == nil {
		return 0
	}
	sum := 0
	if root.Left != nil {
		sum += getSumOfTree(root.Left)
	}
	if root.Right != nil {
		sum += getSumOfTree(root.Right)
	}
	root.Val = sum + root.Val
	return root.Val
}

func getMaxProduct(root *data.TreeNode, sum int) {
	if root == nil {
		return
	}
	if root.Left != nil {
		getMaxProduct(root.Left, sum)
		curProduct := root.Left.Val * (sum - root.Left.Val)
		result1339 = utils.Max(curProduct, result1339)
	}
	if root.Right != nil {
		getMaxProduct(root.Right, sum)
		curProduct := root.Right.Val * (sum - root.Right.Val)
		result1339 = utils.Max(curProduct, result1339)
	}
}

func testProblem1339() {
	node2 := &data.TreeNode{Val: 3}
	node1 := &data.TreeNode{Val: 2, Left: node2}
	node3 := &data.TreeNode{Val: 1, Right: node1}
	fmt.Println(maxProduct(node3))
}
