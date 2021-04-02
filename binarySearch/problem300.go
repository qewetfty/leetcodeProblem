package binarySearch

import (
	"fmt"
	"github.com/leetcodeProblem/utils"
)

// Given an integer array nums, return the length of the longest strictly increasing subsequence.
// A subsequence is a sequence that can be derived from an array by deleting some
// or no elements without changing the order of the remaining elements. For
// example, [3,6,2,7] is a subsequence of the array [0,3,1,6,2,2,7].
//	Example 1:
//		Input: nums = [10,9,2,5,3,7,101,18]
//		Output: 4
//		Explanation: The longest increasing subsequence is [2,3,7,101], therefore the length is 4.
//	Example 2:
//		Input: nums = [0,1,0,3,2,3]
//		Output: 4
//	Example 3:
//		Input: nums = [7,7,7,7,7,7,7]
//		Output: 1
//	Constraints:
//		1 <= nums.length <= 2500
//		-104 <= nums[i] <= 104
//	Follow up:
//		Could you come up with the O(n2) solution?
//		Could you improve it to O(n log(n)) time complexity?

// 二分法查询，使用d[i]记录长度为i的最长子序列的最后一个值
func lengthOfLIS(nums []int) int {
	l, n := 1, len(nums)
	if n < 2 {
		return n
	}
	d := make([]int, n+1)
	d[l] = nums[0]
	for i := 1; i < n; i++ {
		if nums[i] > d[l] {
			l++
			d[l] = nums[i]
		} else {
			lo, hi, pos := 1, l, 0
			for lo <= hi {
				mid := (lo + hi) >> 1
				if d[mid] < nums[i] {
					pos = mid
					lo = mid + 1
				} else {
					hi = mid - 1
				}
			}
			d[pos+1] = nums[i]
		}
	}
	return l
}

// 动态规划，用dp[i]记录以i为结尾的最大可能的长度
func lengthOfLIS2(nums []int) int {
	l := len(nums)
	if l < 2 {
		return l
	}
	dp := make([]int, l)
	dp[0] = 1
	for i := 1; i < l; i++ {
		curMaxCount := 1
		for j := i - 1; j >= 0; j-- {
			curCount := 1
			if nums[i] > nums[j] {
				curCount = dp[j] + 1
			}
			curMaxCount = utils.Max(curCount, curMaxCount)
		}
		dp[i] = curMaxCount
	}
	result := 0
	for _, count := range dp {
		result = utils.Max(count, result)
	}
	return result
}

func testProblem300() {
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
	fmt.Println(lengthOfLIS([]int{0, 1, 0, 3, 2, 3}))
	fmt.Println(lengthOfLIS([]int{7, 7, 7, 7, 7, 7, 7}))
}
