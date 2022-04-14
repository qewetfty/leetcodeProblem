package treeNode

import "leetcodeProblem/data"

//You are given the root of a binary search tree (BST) and an integer val.
//
// Find the node in the BST that the node's value equals val and return the
//subtree rooted with that node. If such a node does not exist, return null.
//
//
// Example 1:
//
//
//Input: root = [4,2,7,1,3], val = 2
//Output: [2,1,3]
//
//
// Example 2:
//
//
//Input: root = [4,2,7,1,3], val = 5
//Output: []
//
//
//
// Constraints:
//
//
// The number of nodes in the tree is in the range [1, 5000].
// 1 <= Node.val <= 10⁷
// root is a binary search tree.
// 1 <= val <= 10⁷
//
// Related Topics Tree Binary Search Tree Binary Tree 👍 2700 👎 141

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func searchBST(root *data.TreeNode, val int) *data.TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}
	if val < root.Val {
		return searchBST(root.Left, val)
	} else {
		return searchBST(root.Right, val)
	}
}

//leetcode submit region end(Prohibit modification and deletion)
