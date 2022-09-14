package main

import "leetcodeProblem/data"

var result int

func pseudoPalindromicPaths(root *data.TreeNode) int {
	if root == nil {
		return 0
	}
	result = 0
	dfs1457(root, 0)
	return result
}

func dfs1457(node *data.TreeNode, temp int) {
	n := temp ^ (1 << node.Val)
	if node.Left == nil && node.Right == nil {
		if n == 0 || n&(n-1) == 0 {
			result++
		}
	}
	if node.Left != nil {
		dfs1457(node.Left, n)
	}
	if node.Right != nil {
		dfs1457(node.Right, n)
	}
}

func main() {

}
