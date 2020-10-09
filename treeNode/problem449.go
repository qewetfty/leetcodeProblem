package treeNode

import (
	"github.com/leetcodeProblem/data"
	"strconv"
	"strings"
)

// Serialization is converting a data structure or object into a sequence of bits
// so that it can be stored in a file or memory buffer, or transmitted across a
// network connection link to be reconstructed later in the same or another
// computer environment.
// Design an algorithm to serialize and deserialize a binary search tree. There
// is no restriction on how your serialization/deserialization algorithm should
// work. You need to ensure that a binary search tree can be serialized to a
// string, and this string can be deserialized to the original tree structure.
// The encoded string should be as compact as possible.
//	Example 1:
//		Input: root = [2,1,3]
//		Output: [2,1,3]
//	Example 2:
//		Input: root = []
//		Output: []
//	Constraints:
//		The number of nodes in the tree is in the range [0, 104].
//		0 <= Node.val <= 104
//		The input tree is guaranteed to be a binary search tree.

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *data.TreeNode) string {
	if root == nil {
		return ""
	}
	valueList := make([]string, 0)
	this.serializeTreeNode(root, &valueList)
	return strings.Join(valueList, ",")
}

func (this *Codec) serializeTreeNode(root *data.TreeNode, valueList *[]string) {
	if root == nil {
		*valueList = append(*valueList, "null")
		return
	}
	*valueList = append(*valueList, strconv.Itoa(root.Val))
	this.serializeTreeNode(root.Left, valueList)
	this.serializeTreeNode(root.Right, valueList)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *data.TreeNode {
	valueList := strings.Split(data, ",")
	if len(data) == 0 || len(valueList) == 0 {
		return nil
	}
	pos := 0
	return this.buildTreeNode(valueList, &pos)
}

func (this *Codec) buildTreeNode(valueList []string, pos *int) *data.TreeNode {
	if strings.EqualFold("null", valueList[*pos]) {
		*pos++
		return nil
	}
	val, _ := strconv.Atoi(valueList[*pos])
	node := &data.TreeNode{Val: val}
	*pos++
	node.Left = this.buildTreeNode(valueList, pos)
	node.Right = this.buildTreeNode(valueList, pos)
	return node
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor()
 * deser := Constructor()
 * tree := ser.serialize(root)
 * ans := deser.deserialize(tree)
 * return ans
 */
