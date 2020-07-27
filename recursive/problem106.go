package recursive

import (
	"fmt"
	"github.com/leetcodeProblem/data"
)

// Given inorder and postorder traversal of a tree, construct the binary tree.
//	Note:
//		You may assume that duplicates do not exist in the tree.
//	For example, given
//		inorder = [9,3,15,20,7]
//		postorder = [9,15,7,20,3]
//	Return the following binary tree:
// 		   3
// 		  / \
// 		 9  20
// 		   /  \
// 		  15   7

var (
	inorderMap map[int]int
	in         []int
	post       []int
	postIndex  int
)

func buildTree(inorder []int, postorder []int) *data.TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}
	inorderMap, in, post = make(map[int]int, 0), inorder, postorder
	for i, val := range inorder {
		inorderMap[val] = i
	}
	postIndex = len(postorder) - 1
	return recursive(0, postIndex)
}

/**
left, right 为中序遍历的起止位置
*/
func recursive(left, right int) *data.TreeNode {
	if left > right {
		return nil
	}
	val := post[postIndex]
	postIndex--
	root := data.NewTreeNode(val)
	rootIndex := inorderMap[val]
	root.Right = recursive(rootIndex+1, right)
	root.Left = recursive(left, rootIndex-1)
	return root
}

func testProblem106() {
	fmt.Println(buildTree([]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3}))
}
