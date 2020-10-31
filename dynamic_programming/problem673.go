package dynamic_programming

import (
	"fmt"
	"github.com/leetcodeProblem/utils"
)

// Given an integer array nums, return the number of longest increasing subsequences.
// Example 1:
//		Input: nums = [1,3,5,4,7]
//		Output: 2
//		Explanation: The two longest increasing subsequences are [1, 3, 4, 7] and [1, 3, 5, 7].
//	Example 2:
//		Input: nums = [2,2,2,2,2]
//		Output: 5
//		Explanation: The length of longest continuous increasing subsequence is 1, and there are 5 subsequences'
//		length is 1, so output 5.
//	Constraints:
//		0 <= nums.length <= 2000
//		-106 <= nums[i] <= 106

func findNumberOfLIS(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}
	// 构造dp和count数组
	// dp[i]表示：以nums[i]结尾的最长递增字串的长度
	// count[i]表示：以nums[i]结尾的最长递增字串的数量
	// max记录最长的递增字串的长度，方便之后计算结果
	dp, count, max := make([]int, l), make([]int, l), 1
	for i := 0; i < l; i++ {
		// 每一个自己都初始化为1
		dp[i], count[i] = 1, 1
	}
	for i := 0; i < l; i++ {
		for j := 0; j < i; j++ {
			// 只有当nums[i]比nums[j]大的时候，才会组成[...j i]的递增序列
			if nums[i] > nums[j] {
				// 如果是大于，则表示是第一次找到长度是dp[j]+1的序列，个数就是dp[j]的个数
				if dp[j]+1 > dp[i] {
					dp[i], count[i] = utils.Max(dp[i], dp[j]+1), count[j]
					// 如果dp[j]+1是dp[i]，则表示之前已经找到过了dp[i]长度，只需要把这次找到的数量加上就好了
				} else if dp[j]+1 == dp[i] {
					count[i] += count[j]
				}
			}
		}
		max = utils.Max(max, dp[i])
	}
	res := 0
	for i := 0; i < l; i++ {
		if dp[i] == max {
			res += count[i]
		}
	}
	return res
}

func main() {
	fmt.Println(findNumberOfLIS([]int{1, 3, 5, 4, 7}))
	fmt.Println(findNumberOfLIS([]int{2, 2, 2, 2, 2}))
}
