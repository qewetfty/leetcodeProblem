package dfs_bfs

import "github.com/leetcodeProblem/data"

// Given a binary tree where node values are digits from 1 to 9. A path in the
// binary tree is said to be pseudo-palindromic if at least one permutation of
// the node values in the path is a palindrome.
// Return the number of pseudo-palindromic paths going from the root node to leaf nodes.
//	Example 1:
//		Input: root = [2,3,1,3,1,null,1]
//		Output: 2
//		Explanation: The figure above represents the given binary tree. There are three paths going from the root
//					node to leaf nodes: the red path [2,3,3], the green path [2,1,1], and the path [2,3,1]. Among these paths
//					only red path and green path are pseudo-palindromic paths since the red path [2,3,3] can be rearranged in
//					[3,2,3] (palindrome) and the green path [2,1,1] can be rearranged in [1,2,1] (palindrome).
//	Example 2:
//		Input: root = [2,1,1,1,3,null,null,null,null,null,1]
//		Output: 1
//		Explanation: The figure above represents the given binary tree. There are three paths going from the root node
//					to leaf nodes: the green path [2,1,1], the path [2,1,3,1], and the path [2,1]. Among these paths
//					only the green path is pseudo-palindromic since [2,1,1] can be rearranged in [1,2,1] (palindrome).
//	Example 3:
//		Input: root = [9]
//		Output: 1
//	Constraints:
//		The given binary tree will have between 1 and 10^5 nodes.
//		Node values are digits from 1 to 9.

var result1457 int

func pseudoPalindromicPaths(root *data.TreeNode) int {
	result1457 = 0
	numberMap := make(map[int]int)
	dfs1457(root, numberMap)
	return result1457
}

func dfs1457(node *data.TreeNode, numberMap map[int]int) {
	if node == nil {
		return
	}
	numberMap[node.Val]++
	// 遍历到叶子节点，处理结果
	if node.Left == nil && node.Right == nil {
		oddNumber := 0
		for _, val := range numberMap {
			if val%2 == 1 {
				oddNumber++
			}
		}
		if oddNumber <= 1 {
			result1457++
		}
		numberMap[node.Val]--
		return
	}
	dfs1457(node.Left, numberMap)
	dfs1457(node.Right, numberMap)
	numberMap[node.Val]--
}

func pseudoPalindromicPathsBit(root *data.TreeNode) int {
	if root == nil {
		return 0
	}
	result1457 = 0
	dfs1457Bit(root, 0)
	return result1457
}

func dfs1457Bit(node *data.TreeNode, temp int) {
	n := temp ^ (1 << node.Val)
	if node.Left == nil && node.Right == nil {
		if n == 0 || n&(n-1) == 0 {
			result1457++
		}
	}
	if node.Left != nil {
		dfs1457Bit(node.Left, n)
	}
	if node.Right != nil {
		dfs1457Bit(node.Right, n)
	}
}
