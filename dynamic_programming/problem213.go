package dynamic_programming

import (
	"fmt"
	"github.com/leetcodeProblem/utils"
)

// You are a professional robber planning to rob houses along a street. Each house has a certain amount of money stashed.
// All houses at this place are arranged in a circle. That means the first house is the neighbor of the last one.
// Meanwhile, adjacent houses have security system connected and it will automatically contact the police if two
// adjacent houses were broken into on the same night.
// Given a list of non-negative integers representing the amount of money of each house, determine the maximum amount
// of money you can rob tonight without alerting the police.
//	Example 1:
//		Input: [2,3,2]
//		Output: 3
//		Explanation: You cannot rob house 1 (money = 2) and then rob house 3 (money = 2),
//		             because they are adjacent houses.
//	Example 2:
//		Input: [1,2,3,1]
//		Output: 4
//		Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
//		             Total amount you can rob = 1 + 3 = 4.

func rob213(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	} else if l == 1 {
		return nums[0]
	} else if l == 2 {
		return utils.Max(nums[0], nums[1])
	} else if l == 3 {
		return utils.Max(nums[0], utils.Max(nums[1], nums[2]))
	}
	dp1 := make([]int, l)
	dp2 := make([]int, l)
	dp1[0], dp1[1] = nums[0], utils.Max(nums[0], nums[1])
	dp1[2] = utils.Max(dp1[1], dp1[0]+nums[2])
	dp2[0], dp2[1] = nums[1], utils.Max(nums[1], nums[2])
	res := 0
	for i := 3; i < l; i++ {
		dp1[i] = utils.Max(dp1[i-1], dp1[i-2]+nums[i])
		dp2[i-1] = utils.Max(dp2[i-2], dp2[i-3]+nums[i])
		res = utils.Max(res, utils.Max(dp1[i-1], dp2[i-3]+nums[i]))
	}
	return res
}

// official method
// max(dp[:l-1], dp[1:])
func rob2(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}
	if l == 1 {
		return nums[0]
	}
	return utils.Max(helper(nums[:l-1]), helper(nums[1:]))
}

func helper(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}
	if l == 1 {
		return nums[0]
	}
	if l == 2 {
		return utils.Max(nums[0], nums[1])
	}
	dp := make([]int, l)
	dp[0] = nums[0]
	dp[1] = utils.Max(nums[0], nums[1])
	for i := 2; i < l; i++ {
		dp[i] = utils.Max(dp[i-2]+nums[i], dp[i-1])
	}
	return dp[l-1]
}

func testProblem213() {
	fmt.Println(rob2([]int{1, 3, 1, 3, 100}))
	fmt.Println(rob2([]int{2, 3, 2}))
	fmt.Println(rob2([]int{1, 2, 3, 1}))
}
