package recursive

import "github.com/leetcodeProblem/data"

// Given a binary tree, find the lowest common ancestor (LCA) of two given nodes in the tree.
// According to the definition of LCA on Wikipedia: “The lowest common ancestor is defined between two nodes p and q
// as the lowest node in T that has both p and q as descendants (where we allow a node to be a descendant of itself).”
// Given the following binary tree:  root = [3,5,1,6,2,0,8,null,null,7,4]
//	Example 1:
//		Input: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
//		Output: 3
//		Explanation: The LCA of nodes 5 and 1 is 3.
//	Example 2:
//		Input: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
//		Output: 5
//		Explanation: The LCA of nodes 5 and 4 is 5, since a node can be a descendant of itself according to the LCA definition.

// Way to think:
//	1. if root is p or q, then their lowest common ancestor is root.
//	2. if p and q are in different side of tree node, then their lowest common ancestor is root.
//	3. if p and q are in the same side, we recursive in that side and do search again.

func lowestCommonAncestor(root, p, q *data.TreeNode) *data.TreeNode {
	if root == nil || p == root || q == root {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	// if left find is null, then p and q must in the right side. So result is right.
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	// Both left and right are not null means that they are in different side.
	// So their lowest common ancestor is root.
	return root
}
