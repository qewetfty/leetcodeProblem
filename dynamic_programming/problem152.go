package dynamic_programming

import (
	"fmt"
	"github.com/leetcodeProblem/utils"
)

// Given an integer array nums, find the contiguous subarray within an array (containing at least one number)
// which has the largest product.
//	Example 1:
//		Input: [2,3,-2,4]
//		Output: 6
//		Explanation: [2,3] has the largest product 6.
//	Example 2:
//		Input: [-2,0,-1]
//		Output: 0
//		Explanation: The result cannot be 2, because [-2,-1] is not a subarray.

func maxProduct(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}
	min, max, res := nums[0], nums[0], nums[0]
	for i := 1; i < l; i++ {
		minF, maxF := min, max
		min = utils.Min(minF*nums[i], utils.Min(maxF*nums[i], nums[i]))
		max = utils.Max(minF*nums[i], utils.Max(maxF*nums[i], nums[i]))
		res = utils.Max(res, max)
	}
	return res
}

func testProblem152() {
	fmt.Println(maxProduct([]int{2, 3, -2, 4}))
	fmt.Println(maxProduct([]int{-2, 0, -1}))
}
