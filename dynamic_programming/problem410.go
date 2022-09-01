package dynamic_programming

import (
	"leetcodeProblem/utils"
	"math"
)

//Given an array nums which consists of non-negative integers and an integer m,
//you can split the array into m non-empty continuous subarrays.
//
// Write an algorithm to minimize the largest sum among these m subarrays.
//
//
// Example 1:
//
//
//Input: nums = [7,2,5,10,8], m = 2
//Output: 18
//Explanation:
//There are four ways to split nums into two subarrays.
//The best way is to split it into [7,2,5] and [10,8],
//where the largest sum among the two subarrays is only 18.
//
//
// Example 2:
//
//
//Input: nums = [1,2,3,4,5], m = 2
//Output: 9
//
//
// Example 3:
//
//
//Input: nums = [1,4,4], m = 3
//Output: 4
//
//
//
// Constraints:
//
//
// 1 <= nums.length <= 1000
// 0 <= nums[i] <= 10⁶
// 1 <= m <= min(50, nums.length)
//
// Related Topics Array Binary Search Dynamic Programming Greedy 👍 4575 👎 128

// dp解法
func splitArray(nums []int, m int) int {
	l := len(nums)
	dp := make([][]int, l+1)
	for i := 0; i <= l; i++ {
		dp[i] = make([]int, m+1)
		for j := 0; j <= m; j++ {
			dp[i][j] = math.MaxInt32
		}
	}
	sub := make([]int, l+1)
	for i := 0; i < l; i++ {
		sub[i+1] = sub[i] + nums[i]
	}
	dp[0][0] = 0
	for i := 1; i <= l; i++ {
		for j := 1; j <= utils.Min(i, m); j++ {
			for k := 0; k < i; k++ {
				dp[i][j] = utils.Min(dp[i][j], utils.Max(dp[k][j-1], sub[i]-sub[k]))
			}
		}
	}
	return dp[l][m]
}

// 二分法解法
//leetcode submit region begin(Prohibit modification and deletion)
func splitArray2(nums []int, m int) int {
	left, right := 0, 0
	for _, num := range nums {
		right += num
		if left < num {
			left = num
		}
	}
	for left < right {
		mid := (left + right) >> 1
		if check410(nums, mid, m) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func check410(nums []int, x, m int) bool {
	sum, cnt := 0, 1
	for _, num := range nums {
		if sum+num > x {
			cnt++
			sum = num
		} else {
			sum += num
		}
	}
	return cnt <= m
}

//leetcode submit region end(Prohibit modification and deletion)
