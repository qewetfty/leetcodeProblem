package dynamic_programming

import (
	"fmt"
	"sort"
)

// Given a set of distinct positive integers nums, return the largest subset
// answer such that every pair (answer[i], answer[j]) of elements in this subset
// satisfies:
//	answer[i] % answer[j] == 0, or
//	answer[j] % answer[i] == 0
//	If there are multiple solutions, return any of them.
//	Example 1:
//		Input: nums = [1,2,3]
//		Output: [1,2]
//		Explanation: [1,3] is also accepted.
//	Example 2:
//		Input: nums = [1,2,4,8]
//		Output: [1,2,4,8]
//	Constraints:
//		1 <= nums.length <= 1000
//		1 <= nums[i] <= 2 * 109
//		All the integers in nums are unique.

func largestDivisibleSubset(nums []int) []int {
	sort.Ints(nums)

	l := len(nums)
	dp := make([]int, l)
	for i := 0; i < l; i++ {
		dp[i] = 1
	}
	maxSize, maxValue := 1, 1
	for i := 1; i < l; i++ {
		for j := 0; j < i; j++ {
			if nums[i]%nums[j] == 0 && dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
			}
		}
		if dp[i] > maxSize {
			maxSize, maxValue = dp[i], nums[i]
		}
	}
	if maxSize == 1 {
		return []int{nums[0]}
	}
	result := make([]int, 0)
	for i := l - 1; i >= 0 && maxSize > 0; i-- {
		if dp[i] == maxSize && maxValue%nums[i] == 0 {
			result = append(result, nums[i])
			maxValue = nums[i]
			maxSize--
		}
	}
	return result
}

func testProblem368() {
	fmt.Println(largestDivisibleSubset([]int{1, 2, 3}))
	fmt.Println(largestDivisibleSubset([]int{1, 2, 4, 8}))
}
