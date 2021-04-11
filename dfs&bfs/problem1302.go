package dfs_bfs

import "github.com/leetcodeProblem/data"

// Given the root of a binary tree, return the sum of values of its deepest leaves.
//	Example 1:
//		Input: root = [1,2,3,4,5,null,6,7,null,null,null,null,8]
//		Output: 15
//	Example 2:
//		Input: root = [6,7,8,2,7,1,3,9,null,1,4,null,null,null,5]
//		Output: 19
//	Constraints:
//		The number of nodes in the tree is in the range [1, 104].
//		1 <= Node.val <= 100

// dfs 解法
var maxDepth int
var sum int

func deepestLeavesSum(root *data.TreeNode) int {
	maxDepth, sum = -1, 0
	dfs1302(root, 0)
	return sum
}

func dfs1302(node *data.TreeNode, depth int) {
	if node == nil {
		return
	}
	if depth > maxDepth {
		maxDepth, sum = depth, node.Val
	} else if depth == maxDepth {
		sum += node.Val
	}
	dfs1302(node.Left, depth+1)
	dfs1302(node.Right, depth+1)
}

// bfs解法
func deepestLeavesSumBFS(root *data.TreeNode) int {
	queue := make([]*data.TreeNode, 0)
	queue = append(queue, root)
	result := root.Val
	for len(queue) > 0 {
		l, curSum := len(queue), 0
		for i := 0; i < l; i++ {
			curNode := queue[i]
			curSum += curNode.Val
			if curNode.Left != nil {
				queue = append(queue, curNode.Left)
			}
			if curNode.Right != nil {
				queue = append(queue, curNode.Right)
			}
		}
		queue = queue[l:]
		result = curSum
	}
	return result
}
