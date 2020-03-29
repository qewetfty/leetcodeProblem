package treeNode

// Given an n-ary tree, return the postorder traversal of its nodes' values.
// Nary-Tree input serialization is represented in their level order traversal,
// each group of children is separated by the null value (See examples).
//	Follow up:
//		Recursive solution is trivial, could you do it iteratively?

type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) []int {
	res := make([]int, 0)
	helper(root, &res)
	return res
}

func helper(root *Node, res *[]int) {
	if root == nil {
		return
	}
	for _, childNode := range root.Children {
		helper(childNode, res)
	}
	*res = append(*res, root.Val)
}
