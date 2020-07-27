package recursive

import "github.com/leetcodeProblem/data"

// Given preorder and inorder traversal of a tree, construct the binary tree.
//	Note:
//		You may assume that duplicates do not exist in the tree.
//	For example, given
//		preorder = [3,9,20,15,7]
//		inorder = [9,3,15,20,7]
//	Return the following binary tree:
//		    3
//		   / \
//		  9  20
//		    /  \
//		   15   7

var (
	inorderMap105 map[int]int
	pre           []int
	preIndex      int
)

func buildTree105(preorder []int, inorder []int) *data.TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	inorderMap105, pre, preIndex = make(map[int]int, 0), preorder, 0
	for i, val := range inorder {
		inorderMap105[val] = i
	}
	return recursive105(0, len(inorder)-1)
}

func recursive105(left, right int) *data.TreeNode {
	if left > right {
		return nil
	}
	val := pre[preIndex]
	preIndex++
	root := data.NewTreeNode(val)
	rootIndex := inorderMap105[val]
	root.Left = recursive105(left, rootIndex-1)
	root.Right = recursive105(rootIndex+1, right)
	return root
}
