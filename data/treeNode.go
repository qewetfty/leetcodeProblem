package data

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(val int) *TreeNode {
	treeNode := new(TreeNode)
	treeNode.Val = val
	return treeNode
}
