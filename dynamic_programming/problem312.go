package dynamic_programming

import (
	"fmt"
	"github.com/leetcodeProblem/utils"
)

// Given n balloons, indexed from 0 to n-1. Each balloon is painted with a number
// on it represented by array nums. You are asked to burst all the balloons. If
// the you burst balloon i you will get nums[left] * nums[i] * nums[right] coins.
// Here left and right are adjacent indices of i. After the burst, the left and
// right then becomes adjacent.
// Find the maximum coins you can collect by bursting the balloons wisely.
//	Note:
//		You may imagine nums[-1] = nums[n] = 1. They are not real therefore you can not burst them.
//		0 ≤ n ≤ 500, 0 ≤ nums[i] ≤ 100
//	Example:
//		Input: [3,1,5,8]
//		Output: 167
//		Explanation: nums = [3,1,5,8] --> [3,5,8] -->   [3,8]   -->  [8]  --> []
//		             coins =  3*1*5      +  3*5*8    +  1*3*8      + 1*8*1   = 167

// dp解法，思路同记忆化搜索一样，需要自底向上计算，计算时候需要注意dp依赖的顺序
func maxCoins(nums []int) int {
	l := len(nums)
	// 预先处理nums
	nums = append([]int{1}, nums...)
	nums = append(nums, 1)
	// dp为开区间(i,j)中，可以获取的最大的币的数量
	dp := make([][]int, l+2)
	for i := 0; i < l+2; i++ {
		dp[i] = make([]int, l+2)
	}
	for i := l - 1; i >= 0; i-- {
		for j := i + 2; j <= l+1; j++ {
			for k := i + 1; k < j; k++ {
				sum := nums[i] * nums[k] * nums[j]
				sum = dp[i][k] + sum + dp[k][j]
				dp[i][j] = utils.Max(dp[i][j], sum)
			}
		}
	}
	return dp[0][l+1]
}

// 记忆化搜索方法，设solve(i,j)是区间i到j的最优解，每次都去遍历。
// 与暴力法不同的是，将暴力法中重复计算的结果进行保存
var midResult [][]int

func maxCoins2(nums []int) int {
	l := len(nums)
	// 预先处理nums，向两边插入值
	newNums := make([]int, l+2)
	for i := 1; i <= l; i++ {
		newNums[i] = nums[i-1]
	}
	newNums[0], newNums[l+1] = 1, 1
	midResult = make([][]int, l+2)
	for i := 0; i < l+2; i++ {
		midResult[i] = make([]int, l+2)
		for j := 0; j < l+2; j++ {
			midResult[i][j] = -1
		}
	}
	return solve(0, l+1, newNums)
}

func solve(left, right int, nums []int) int {
	if left >= right-1 {
		return 0
	}
	if midResult[left][right] != -1 {
		return midResult[left][right]
	}
	for i := left + 1; i < right; i++ {
		sum := nums[left] * nums[i] * nums[right]
		sum = solve(left, i, nums) + sum + solve(i, right, nums)
		midResult[left][right] = utils.Max(midResult[left][right], sum)
	}
	return midResult[left][right]
}

func testProblem312() {
	fmt.Println(maxCoins([]int{3, 1, 5, 8}))
	fmt.Println(maxCoins([]int{1, 2, 3, 4, 5, 6, 7, 8}))
	fmt.Println(maxCoins([]int{1}))
	fmt.Println(maxCoins([]int{1, 4}))
	fmt.Println(maxCoins([]int{9, 9, 6}))
}
