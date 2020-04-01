package recursive

import (
	"fmt"
	"github.com/leetcodeProblem/data"
	"strconv"
	"strings"
)

// Serialization is the process of converting a data structure or object into a
// sequence of bits so that it can be stored in a file or memory buffer,
// or transmitted across a network connection link to be reconstructed later in the same
// or another computer environment.
// Design an algorithm to serialize and deserialize a binary tree.
// There is no restriction on how your serialization/deserialization algorithm should work.
// You just need to ensure that a binary tree can be serialized to a string and this string can
// be deserialized to the original tree structure.
//	Example:
//		You may serialize the following tree:
//
//		    1
//		   / \
//		  2   3
//		     / \
//		    4   5
//		as "[1,2,3,null,null,4,5]"
//	Clarification: The above format is the same as how LeetCode serializes a binary tree.
//		You do not necessarily need to follow this format, so please be creative and come
//		up with different approaches yourself.
//	Note: Do not use class member/global/static variables to store states.
//		Your serialize and deserialize algorithms should be stateless.

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
}

func Constructor() Codec {
	codec := new(Codec)
	return *codec
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *data.TreeNode) string {
	if root == nil {
		return "null"
	}
	valueList := make([]string, 0)
	serializeTreeNode(root, &valueList)
	return strings.Join(valueList, ",")
}

func serializeTreeNode(root *data.TreeNode, valueList *[]string) {
	if root == nil {
		*valueList = append(*valueList, "null")
		return
	}
	*valueList = append(*valueList, strconv.Itoa(root.Val))
	serializeTreeNode(root.Left, valueList)
	serializeTreeNode(root.Right, valueList)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *data.TreeNode {
	valueList := strings.Split(data, ",")
	if len(data) == 0 || len(valueList) == 0 {
		return nil
	}
	pos := 0
	return buildTreeNode(valueList, &pos)
}

func buildTreeNode(valueList []string, pos *int) *data.TreeNode {
	if strings.EqualFold("null", valueList[*pos]) {
		*pos++
		return nil
	}
	node := new(data.TreeNode)
	node.Val, _ = strconv.Atoi(valueList[*pos])
	*pos++
	node.Left = buildTreeNode(valueList, pos)
	node.Right = buildTreeNode(valueList, pos)
	return node
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * data := obj.serialize(root);
 * ans := obj.deserialize(data);
 */

func testProblem297() {
	obj := Constructor()
	ans := obj.deserialize("")
	fmt.Println(obj.serialize(ans))
}
