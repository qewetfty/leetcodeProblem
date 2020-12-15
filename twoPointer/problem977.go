package twoPointer

import (
	"fmt"
	"github.com/leetcodeProblem/utils"
)

// Given an integer array nums sorted in non-decreasing order, return an array of
// the squares of each number sorted in non-decreasing order.
//	Example 1:
//		Input: nums = [-4,-1,0,3,10]
//		Output: [0,1,9,16,100]
//		Explanation: After squaring, the array becomes [16,1,0,9,100].
//		After sorting, it becomes [0,1,9,16,100].
//	Example 2:
//		Input: nums = [-7,-3,2,3,11]
//		Output: [4,9,9,49,121]
//	Constraints:
//		1 <= nums.length <= 104
//		-104 <= nums[i] <= 104
//		nums is sorted in non-decreasing order.

func sortedSquares(nums []int) []int {
	l := len(nums)
	result := make([]int, l)
	i, left, right := l-1, 0, l-1
	for left <= right {
		if utils.Abs(nums[left]) >= utils.Abs(nums[right]) {
			result[i] = nums[left] * nums[left]
			left++
		} else {
			result[i] = nums[right] * nums[right]
			right--
		}
		i--
	}
	return result
}

func testProblem977() {
	fmt.Println(sortedSquares([]int{-4, -1, 0, 3, 10}))
	fmt.Println(sortedSquares([]int{-7, -3, 2, 3, 11}))
}
