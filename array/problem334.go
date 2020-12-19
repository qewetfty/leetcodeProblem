package array

import (
	"fmt"
	"math"
)

// Given an integer array nums, return true if there exists a triple of indices
// (i, j, k) such that i < j < k and nums[i] < nums[j] < nums[k]. If no such
// indices exists, return false.
//	Example 1:
//		Input: nums = [1,2,3,4,5]
//		Output: true
//		Explanation: Any triplet where i < j < k is valid.
//	Example 2:
//		Input: nums = [5,4,3,2,1]
//		Output: false
//		Explanation: No triplet exists.
//	Example 3:
//		Input: nums = [2,1,5,0,4,6]
//		Output: true
//		Explanation: The triplet (3, 4, 5) is valid because nums[3] == 0 < nums[4] == 4 < nums[5] == 6.
//	Constraints:
//		1 <= nums.length <= 105
//		-231 <= nums[i] <= 231 - 1

func increasingTriplet(nums []int) bool {
	l := len(nums)
	if l < 3 {
		return false
	}
	low1, low2 := nums[0], math.MaxInt32
	for i := 1; i < l; i++ {
		curNum := nums[i]
		if curNum < low1 {
			low1 = curNum
		} else if curNum > low1 {
			if curNum < low2 {
				low2 = curNum
			} else if curNum > low2 {
				return true
			}
		}
	}
	return false
}

func testProblem334() {
	fmt.Println(increasingTriplet([]int{1, 2, 3, 4, 5}))
	fmt.Println(increasingTriplet([]int{5, 4, 3, 2, 1}))
	fmt.Println(increasingTriplet([]int{2, 1, 5, 0, 4, 6}))
}
