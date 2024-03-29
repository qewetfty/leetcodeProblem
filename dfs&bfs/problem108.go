package dfs_bfs

import "github.com/leetcodeProblem/data"

// Given an integer array nums where the elements are sorted in ascending order,
// convert it to a height-balanced binary search tree.
// A height-balanced binary tree is a binary tree in which the depth of the two
// subtrees of every node never differs by more than one.
//	Example 1:
//		Input: nums = [-10,-3,0,5,9]
//		Output: [0,-3,9,-10,null,5]
//		Explanation: [0,-10,5,null,-3,null,9] is also accepted:
//	Example 2:
//		Input: nums = [1,3]
//		Output: [3,1]
//		Explanation: [1,3] and [3,1] are both a height-balanced BSTs.
//	Constraints:
//		1 <= nums.length <= 104
//		-104 <= nums[i] <= 104
//		nums is sorted in a strictly increasing order.

func sortedArrayToBST(nums []int) *data.TreeNode {
	return dfs108(nums, 0, len(nums)-1)
}

func dfs108(nums []int, left, right int) *data.TreeNode {
	if left > right {
		return nil
	}
	mid := (left + right) >> 1
	midNode := &data.TreeNode{Val: nums[mid]}
	midNode.Left = dfs108(nums, left, mid-1)
	midNode.Right = dfs108(nums, mid+1, right)
	return midNode
}
