package dfs_bfs

import "github.com/leetcodeProblem/data"

// Given a non-empty binary tree, return the average value of the nodes on each level in the form of an array.
//	Example 1:
//		Input:
//		  	  3
//		  	 / \
//		  	9  20
//		  	  /  \
//		  	 15   7
//		Output: [3, 14.5, 11]
//		Explanation:
//			The average value of nodes on level 0 is 3,  on level 1 is 14.5, and on level 2 is 11. Hence return [3, 14.5, 11].
//	Note:
//		The range of node's value is in the range of 32-bit signed integer.

func averageOfLevels(root *data.TreeNode) []float64 {
	queue, result := make([]*data.TreeNode, 0), make([]float64, 0)
	if root == nil {
		return result
	}
	queue = append(queue, root)
	for len(queue) > 0 {
		l, sum := len(queue), 0.0
		for i := 0; i < l; i++ {
			node := queue[i]
			sum = sum + float64(node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		sum = sum / float64(l)
		queue = queue[l:]
		result = append(result, sum)
	}
	return result
}
