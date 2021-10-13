package dfs_bfs

import (
	"leetcodeProblem/data"
	"sort"
)

// Given an array of integers preorder, which represents the preorder traversal
// of a BST (i.e., binary search tree), construct the tree and return its root.
// It is guaranteed that there is always possible to find a binary search tree
// with the given requirements for the given test cases.
// A binary search tree is a binary tree where for every node, any descendant of
// Node.left has a value strictly less than Node.val, and any descendant of
// Node.right has a value strictly greater than Node.val.
// A preorder traversal of a binary tree displays the value of the node first,
// then traverses Node.left, then traverses Node.right.
//	Example 1:
//		Input: preorder = [8,5,1,7,10,12]
//		Output: [8,5,10,1,7,null,12]
//	Example 2:
//		Input: preorder = [1,3]
//		Output: [1,null,3]
//	Constraints:
//		1 <= preorder.length <= 100
//		1 <= preorder[i] <= 108
//		All the values of preorder are unique.

// 对preorder进行排序，获取中序遍历结果。通过前序遍历和中序遍历结果恢复BST的结构
var preIndex int

func bstFromPreorder(preorder []int) *data.TreeNode {
	inorder := make([]int, len(preorder))
	copy(inorder, preorder)
	sort.Ints(inorder)
	inorderMap := make(map[int]int)
	for i, val := range inorder {
		inorderMap[val] = i
	}
	preIndex = 0
	return buildTree1008(preorder, inorder, inorderMap, 0, len(preorder)-1)
}

func buildTree1008(preorder, inorder []int, inorderMap map[int]int, left, right int) *data.TreeNode {
	if left > right {
		return nil
	}
	val := preorder[preIndex]
	preIndex++
	root := &data.TreeNode{Val: val}
	rootIndex := inorderMap[val]
	root.Left = buildTree1008(preorder, inorder, inorderMap, left, rootIndex-1)
	root.Right = buildTree1008(preorder, inorder, inorderMap, rootIndex+1, right)
	return root
}
