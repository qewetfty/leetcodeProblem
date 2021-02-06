package dfs_bfs

import "github.com/leetcodeProblem/data"

// Given a binary tree, imagine yourself standing on the right side of it, return
// the values of the nodes you can see ordered from top to bottom.
//	Example:
//		Input: [1,2,3,null,5,null,4]
//		Output: [1, 3, 4]
//		Explanation:
//			   1            <---
//			 /   \
//			2     3         <---
//			 \     \
//			  5     4       <---

// bfs解法
func rightSideView(root *data.TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}
	queue := make([]*data.TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		curLen := len(queue)
		result = append(result, queue[0].Val)
		for i := 0; i < curLen; i++ {
			node := queue[i]
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
		}
		queue = queue[curLen:]
	}
	return result
}
