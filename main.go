package main

import (
	"fmt"
	"github.com/leetcodeProblem/data"
)

// You are given the root of a binary tree with n nodes, where each node is
// uniquely assigned a value from 1 to n. You are also given a sequence of n
// values voyage, which is the desired pre-order traversal of the binary tree.
// Any node in the binary tree can be flipped by swapping its left and right
// subtrees. For example, flipping node 1 will have the following effect:
// Flip the smallest number of nodes so that the pre-order traversal of the tree matches voyage.
// Return a list of the values of all flipped nodes. You may return the answer in
// any order. If it is impossible to flip the nodes in the tree to make the
// pre-order traversal match voyage, return the list [-1].
//	Example 1:
//		Input: root = [1,2], voyage = [2,1]
//		Output: [-1]
//		Explanation: It is impossible to flip the nodes such that the pre-order traversal matches voyage.
//	Example 2:
//		Input: root = [1,2,3], voyage = [1,3,2]
//		Output: [1]
//		Explanation: Flipping node 1 swaps nodes 2 and 3, so the pre-order traversal matches voyage.
//	Example 3:
//		Input: root = [1,2,3], voyage = [1,2,3]
//		Output: []
//		Explanation: The tree's pre-order traversal already matches voyage, so no nodes need to be flipped.
//	Constraints:
//		The number of nodes in the tree is n.
//		n == voyage.length
//		1 <= n <= 100
//		1 <= Node.val, voyage[i] <= n
//		All the values in the tree are unique.
//		All the values in voyage are unique.

var number int

func flipMatchVoyage(root *data.TreeNode, voyage []int) []int {
	number = 0
	result := make([]int, 0)
	dfs971(root, voyage, &result)
	return result
}

func dfs971(node *data.TreeNode, voyage []int, result *[]int) {
	if len(*result) > 0 && (*result)[0] == -1 {
		return
	}
	if node.Val != voyage[number] {
		*result = []int{-1}
		return
	}
	number++
	if node.Left != nil {
		if node.Left.Val != voyage[number] {
			if node.Right != nil {
				*result = append(*result, node.Val)
				dfs971(node.Right, voyage, result)
				dfs971(node.Left, voyage, result)
			} else {
				*result = []int{-1}
				return
			}
		} else {
			dfs971(node.Left, voyage, result)
			if node.Right != nil {
				dfs971(node.Right, voyage, result)
			}
		}
	} else if node.Right != nil {
		dfs971(node.Right, voyage, result)
	}
}

func main() {
	node1 := &data.TreeNode{Val: 3}
	node2 := &data.TreeNode{Val: 2, Left: node1}
	node3 := &data.TreeNode{Val: 1, Left: node2}
	fmt.Println(flipMatchVoyage(node3, []int{1, 2, 3}))
}
