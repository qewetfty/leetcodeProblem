package treeNode

import "github.com/leetcodeProblem/data"

// Given two binary search trees root1 and root2.
// Return a list containing all the integers from both trees sorted in ascending order.
//	Example 1:
//		Input: root1 = [2,1,4], root2 = [1,0,3]
//		Output: [0,1,1,2,3,4]
//	Example 2:
//		Input: root1 = [0,-10,10], root2 = [5,1,7,0,2]
//		Output: [-10,0,0,1,2,5,7,10]
//	Example 3:
//		Input: root1 = [], root2 = [5,1,7,0,2]
//		Output: [0,1,2,5,7]
//	Example 4:
//		Input: root1 = [0,-10,10], root2 = []
//		Output: [-10,0,10]
//	Example 5:
//		Input: root1 = [1,null,8], root2 = [8,1]
//		Output: [1,1,8,8]
//	Constraints:
//		Each tree has at most 5000 nodes.
//		Each node's value is between [-10^5, 10^5].

// 中序遍历获取两个树的所有值，然后合并结果
func getAllElements(root1 *data.TreeNode, root2 *data.TreeNode) []int {
	nums1, nums2 := make([]int, 0), make([]int, 0)
	inorderTree(root1, &nums1)
	inorderTree(root2, &nums2)
	l1, l2 := len(nums1), len(nums2)
	res, i, j, k := make([]int, l1+l2), l1-1, l2-1, l1+l2-1
	for i >= 0 && j >= 0 {
		if nums1[i] >= nums2[j] {
			res[k] = nums1[i]
			i--
			k--
		} else {
			res[k] = nums2[j]
			j--
			k--
		}
	}
	if i >= 0 {
		for o := i; o >= 0; o-- {
			res[o] = nums1[o]
		}
	}
	if j >= 0 {
		for o := j; o >= 0; o-- {
			res[o] = nums1[o]
		}
	}
	return res
}

func inorderTree(root *data.TreeNode, res *[]int) {
	if root == nil {
		return
	}
	inorderTree(root.Left, res)
	*res = append(*res, root.Val)
	inorderTree(root.Right, res)
}
