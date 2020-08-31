package treeNode

import "github.com/leetcodeProblem/data"

// Given a root node reference of a BST and a key, delete the node with the given key in the BST. Return the root
// node reference (possibly updated) of the BST.
// Basically, the deletion can be divided into two stages:
// Search for a node to remove.
// If the node is found, delete the node.
//	Note: Time complexity should be O(height of tree).
//	Example:
//		root = [5,3,6,2,4,null,7]
//		key = 3
//
//		    5
//		   / \
//		  3   6
//		 / \   \
//		2   4   7
//	Given key to delete is 3. So we find the node with value 3 and delete it.
//	One valid answer is [5,4,6,2,null,null,7], shown in the following BST.
//		    5
//		   / \
//		  4   6
//		 /     \
//		2       7
//	Another valid answer is [5,2,6,null,4,null,7].
//		    5
//		   / \
//		  2   6
//		   \   \
//		    4   7

func deleteNode(root *data.TreeNode, key int) *data.TreeNode {
	if root == nil {
		return nil
	}
	if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	} else if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else {
		if root.Left == nil && root.Right == nil {
			root = nil
		} else if root.Right != nil {
			root.Val = successor(root)
			root.Right = deleteNode(root.Right, root.Val)
		} else {
			root.Val = predecessor(root)
			root.Left = deleteNode(root.Left, root.Val)
		}
	}
	return root
}

func successor(root *data.TreeNode) int {
	root = root.Right
	for root.Left != nil {
		root = root.Left
	}
	return root.Val
}

func predecessor(root *data.TreeNode) int {
	root = root.Left
	for root.Right != nil {
		root = root.Right
	}
	return root.Val
}
