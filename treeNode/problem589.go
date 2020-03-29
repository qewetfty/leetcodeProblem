package treeNode

import "github.com/leetcodeProblem/data"

// Given an n-ary tree, return the postorder traversal of its nodes' values.
// Nary-Tree input serialization is represented in their level order traversal,
// each group of children is separated by the null value (See examples).
//	Follow up:
//		Recursive solution is trivial, could you do it iteratively?

// non-recursive solution
func preorder(root *Node) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	stack := new(data.Stack)
	stack.Push(root)
	for !stack.IsEmpty() {
		node, _ := stack.Pop()
		n := node.(*Node)
		res = append(res, n.Val)
		for i := len(n.Children) - 1; i >= 0; i-- {
			child := n.Children[i]
			stack.Push(child)
		}
	}
	return res
}
