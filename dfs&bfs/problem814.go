package dfs_bfs

import "github.com/leetcodeProblem/data"

// Given the root of a binary tree, return the same tree where every subtree (of
// the given tree) not containing a 1 has been removed.
// A subtree of a node node is node plus every node that is a descendant of node.
//	Example 1:
//		Input: root = [1,null,0,0,1]
//		Output: [1,null,0,null,1]
//		Explanation:
//			Only the red nodes satisfy the property "every subtree not containing a 1".
//			The diagram on the right represents the answer.
//	Example 2:
//		Input: root = [1,0,1,0,0,0,1]
//		Output: [1,null,1,null,1]
//	Example 3:
//		Input: root = [1,1,0,1,1,0,1,0]
//		Output: [1,1,0,1,1,null,1]
//	Constraints:
//		The number of nodes in the tree is in the range [1, 200].
//		Node.val is either 0 or 1.

func pruneTree(root *data.TreeNode) *data.TreeNode {
	dummyNode := &data.TreeNode{Left: root}
	dfs814(root, dummyNode, true)
	return dummyNode.Left
}

func dfs814(node *data.TreeNode, parentNode *data.TreeNode, left bool) {
	if node.Left != nil {
		dfs814(node.Left, node, true)
	}
	if node.Right != nil {
		dfs814(node.Right, node, false)
	}

	if node.Left == nil && node.Right == nil && node.Val == 0 {
		if left {
			parentNode.Left = nil
		} else {
			parentNode.Right = nil
		}
	}
}
