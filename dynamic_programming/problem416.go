package dynamic_programming

import "fmt"

// Given a non-empty array nums containing only positive integers, find if the
// array can be partitioned into two subsets such that the sum of elements in
// both subsets is equal.
//	Example 1:
//		Input: nums = [1,5,11,5]
//		Output: true
//		Explanation: The array can be partitioned as [1, 5, 5] and [11].
//	Example 2:
//		Input: nums = [1,2,3,5]
//		Output: false
//		Explanation: The array cannot be partitioned into equal sum subsets.
//	Constraints:
//		1 <= nums.length <= 200
//		1 <= nums[i] <= 100

// O(target)时间复杂度
func canPartition(nums []int) bool {
	l := len(nums)
	if l < 2 {
		return false
	}
	sum, max := 0, 0
	for _, num := range nums {
		sum += num
		if num > max {
			max = num
		}
	}
	if sum%2 != 0 {
		return false
	}
	target := sum / 2
	if max > target {
		return false
	}
	dp := make([]bool, target+1)
	dp[0], dp[nums[0]] = true, true
	for i := 1; i < l; i++ {
		curNum := nums[i]
		for j := target; j >= curNum; j-- {
			dp[j] = dp[j] || dp[j-curNum]
		}
	}
	return dp[target]
}

// O(n * target)空间复杂度
func canPartition2(nums []int) bool {
	l := len(nums)
	if l < 2 {
		return false
	}
	sum, max := 0, 0
	for _, num := range nums {
		sum += num
		if num > max {
			max = num
		}
	}
	if sum%2 != 0 {
		return false
	}
	target := sum / 2
	if max > target {
		return false
	}
	dp := make([][]bool, l)
	for i := range dp {
		dp[i] = make([]bool, target+1)
	}
	for i := 0; i < l; i++ {
		dp[i][0] = true
	}
	dp[0][nums[0]] = true
	for i := 1; i < l; i++ {
		curNum := nums[i]
		for j := 1; j <= target; j++ {
			if j >= curNum {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-curNum]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[l-1][target]
}

// 记忆化搜索
var memories [][]*bool

func canPartition3(nums []int) bool {
	l := len(nums)
	if l < 2 {
		return false
	}
	sum, max := 0, 0
	for _, num := range nums {
		sum += num
		if num > max {
			max = num
		}
	}
	if sum%2 != 0 {
		return false
	}
	target := sum / 2
	if max > target {
		return false
	}
	memories = make([][]*bool, l)
	for i := range memories {
		memories[i] = make([]*bool, target+1)
	}
	return dfs416(nums, l-1, target)
}

func dfs416(nums []int, n, target int) bool {
	if target == 0 {
		return true
	}
	if n == 0 || target < 0 {
		return false
	}
	if memories[n][target] != nil {
		return *memories[n][target]
	}
	result := dfs416(nums, n-1, target-nums[n-1]) || dfs416(nums, n-1, target)
	memories[n][target] = &result
	return result
}

func testProblem416() {
	fmt.Println(canPartition3([]int{1, 5, 11, 5}))
	fmt.Println(canPartition3([]int{1, 2, 3, 5}))
}
