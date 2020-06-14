package dynamic_programming

import (
	"fmt"
	"github.com/leetcodeProblem/utils"
)

// Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest
// sum and return its sum.
//	Example:
//		Input: [-2,1,-3,4,-1,2,1,-5,4],
//		Output: 6
//		Explanation: [4,-1,2,1] has the largest sum = 6.
//	Follow up:
//		If you have figured out the O(n) solution, try coding another solution using the divide and conquer
//		approach, which is more subtle.

func maxSubArray(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}
	dp := make([]int, l)
	dp[0] = nums[0]
	sum := dp[0]
	for i := 1; i < l; i++ {
		dp[i] = utils.Max(0, dp[i-1]) + nums[i]
		sum = utils.Max(sum, dp[i])
	}
	return sum
}

func testProblem53() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}
