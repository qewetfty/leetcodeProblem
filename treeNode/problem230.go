package treeNode

import "leetcodeProblem/data"

//Given the root of a binary search tree, and an integer k, return the kᵗʰ
//smallest value (1-indexed) of all the values of the nodes in the tree.
//
//
// Example 1:
//
//
//Input: root = [3,1,4,null,2], k = 1
//Output: 1
//
//
// Example 2:
//
//
//Input: root = [5,3,6,2,4,null,null,1], k = 3
//Output: 3
//
//
//
// Constraints:
//
//
// The number of nodes in the tree is n.
// 1 <= k <= n <= 10⁴
// 0 <= Node.val <= 10⁴
//
//
//
// Follow up: If the BST is modified often (i.e., we can do insert and delete
//operations) and you need to find the kth smallest frequently, how would you
//optimize?
// Related Topics Tree Depth-First Search Binary Search Tree Binary Tree 👍 6167
// 👎 124

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var inorderNum []int

func kthSmallest(root *data.TreeNode, k int) int {
	inorderNum = make([]int, 0)
	inorder230(root)
	return inorderNum[k-1]
}

func inorder230(node *data.TreeNode) {
	if node == nil {
		return
	}
	inorder230(node.Left)
	inorderNum = append(inorderNum, node.Val)
	inorder230(node.Right)
}

//leetcode submit region end(Prohibit modification and deletion)
