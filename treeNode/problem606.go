package treeNode

import (
	"fmt"
	"github.com/leetcodeProblem/data"
	"strconv"
	"strings"
)

//Given the root of a binary tree, construct a string consisting of parenthesis
//and integers from a binary tree with the preorder traversal way, and return it.
//
//
// Omit all the empty parenthesis pairs that do not affect the one-to-one
//mapping relationship between the string and the original binary tree.
//
//
// Example 1:
//
//
//Input: root = [1,2,3,4]
//Output: "1(2(4))(3)"
//Explanation: Originally, it needs to be "1(2(4)())(3()())", but you need to
//omit all the unnecessary empty parenthesis pairs. And it will be "1(2(4))(3)"
//
//
// Example 2:
//
//
//Input: root = [1,2,3,null,4]
//Output: "1(2()(4))(3)"
//Explanation: Almost the same as the first example, except we cannot omit the
//first parenthesis pair to break the one-to-one mapping relationship between the
//input and the output.
//
//
//
// Constraints:
//
//
// The number of nodes in the tree is in the range [1, 10⁴].
// -1000 <= Node.val <= 1000
//
//
// Related Topics String Tree Depth-First Search Binary Tree 👍 1829 👎 2326

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func tree2str(root *data.TreeNode) string {
	var sb strings.Builder
	if root == nil {
		return ""
	}
	dfs606(root, &sb)
	return sb.String()
}

func dfs606(node *data.TreeNode, sb *strings.Builder) {
	fmt.Fprint(sb, strconv.Itoa(node.Val))
	if node.Left == nil && node.Right == nil {
		return
	}
	fmt.Fprint(sb, "(")
	if node.Left != nil {
		dfs606(node.Left, sb)
	}
	fmt.Fprint(sb, ")")

	if node.Right != nil {
		fmt.Fprint(sb, "(")
		dfs606(node.Right, sb)
		fmt.Fprint(sb, ")")
	}
}
