package array

import "sort"

// Given two integer arrays nums1 and nums2, return an array of their
// intersection. Each element in the result must appear as many times as it shows
// in both arrays and you may return the result in any order.
//	Example 1:
//		Input: nums1 = [1,2,2,1], nums2 = [2,2]
//		Output: [2,2]
//	Example 2:
//		Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
//		Output: [4,9]
//		Explanation: [9,4] is also accepted.
//	Constraints:
//		1 <= nums1.length, nums2.length <= 1000
//		0 <= nums1[i], nums2[i] <= 1000

func intersect(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	p1, p2 := 0, 0
	l1, l2 := len(nums1), len(nums2)
	result := make([]int, 0)
	for p1 < l1 && p2 < l2 {
		if nums1[p1] == nums2[p2] {
			result = append(result, nums1[p1])
			p1++
			p2++
		} else if nums1[p1] < nums2[p2] {
			p1++
		} else {
			p2++
		}
	}
	return result
}
