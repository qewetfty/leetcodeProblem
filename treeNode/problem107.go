package treeNode

import "github.com/leetcodeProblem/data"

// Given the root of a binary tree, return the bottom-up level order traversal of
// its nodes' values. (i.e., from left to right, level by level from leaf to
// root).
//	Example 1:
//		Input: root = [3,9,20,null,null,15,7]
//		Output: [[15,7],[9,20],[3]]
//	Example 2:
//		Input: root = [1]
//		Output: [[1]]
//	Example 3:
//		Input: root = []
//		Output: []
//	Constraints:
//		The number of nodes in the tree is in the range [0, 2000].
//		-1000 <= Node.val <= 1000

func levelOrderBottom(root *data.TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	queue := make([]*data.TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		l := len(queue)
		curResult := make([]int, 0)
		for i := 0; i < l; i++ {
			node := queue[i]
			curResult = append(curResult, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[l:]
		result = append([][]int{curResult}, result...)
	}
	return result
}
