package twoPointer

import (
	"fmt"
	"sort"
)

// Given an integer array nums, return the number of triplets chosen from the
// array that can make triangles if we take them as side lengths of a triangle.
//	Example 1:
//		Input: nums = [2,2,3,4]
//		Output: 3
//		Explanation: Valid combinations are:
//			2,3,4 (using the first 2)
//			2,3,4 (using the second 2)
//			2,2,3
//	Example 2:
//		Input: nums = [4,2,3,4]
//		Output: 4
//	Constraints:
//		1 <= nums.length <= 1000
//		0 <= nums[i] <= 1000

func triangleNumber(nums []int) int {
	l := len(nums)
	sort.Ints(nums)
	result := 0
	for i := 0; i < l-2; i++ {
		k := i + 2
		for j := i + 1; j < l-1 && nums[i] != 0; j++ {
			for k < l && nums[i]+nums[j] > nums[k] {
				k++
			}
			result += k - j - 1
		}
	}
	return result
}

func testProblem611() {
	fmt.Println(triangleNumber([]int{2, 2, 3, 4}))
	fmt.Println(triangleNumber([]int{4, 2, 3, 4}))
}
