package dynamic_programming

import (
	"fmt"
	"github.com/leetcodeProblem/utils"
	"math"
)

// Say you have an array for which the ith element is the price of a given stock on day i.
// Design an algorithm to find the maximum profit. You may complete at most two transactions.
//	Note: You may not engage in multiple transactions at the same time (i.e., you must sell the stock before you buy again).
//	Example 1:
//		Input: [3,3,5,0,0,3,1,4]
//		Output: 6
//		Explanation: Buy on day 4 (price = 0) and sell on day 6 (price = 3), profit = 3-0 = 3.
//		             Then buy on day 7 (price = 1) and sell on day 8 (price = 4), profit = 4-1 = 3.
//	Example 2:
//		Input: [1,2,3,4,5]
//		Output: 4
//		Explanation: Buy on day 1 (price = 1) and sell on day 5 (price = 5), profit = 5-1 = 4.
//		             Note that you cannot buy on day 1, buy on day 2 and sell them later, as you are
//		             engaging multiple transactions at the same time. You must sell before buying again.
//	Example 3:
//		Input: [7,6,4,3,1]
//		Output: 0
//		Explanation: In this case, no transaction is done, i.e. max profit = 0.

// divide method
func maxProfit123(prices []int) int {
	l := len(prices)
	if l <= 1 {
		return 0
	}
	min, max := prices[0], prices[l-1]
	dp1, dp2 := make([]int, l), make([]int, l)
	dp1[0], dp2[l-1] = 0, 0
	for i := 1; i < l; i++ {
		dp1[i] = utils.Max(dp1[i-1], prices[i]-min)
		dp2[l-i-1] = utils.Max(dp2[l-1], max-prices[l-i-1])
		if prices[i] < min {
			min = prices[i]
		}
		if prices[l-i-1] > max {
			max = prices[l-i-1]
		}
	}
	maxProfit := 0
	for i := 0; i < l; i++ {
		maxProfit = utils.Max(maxProfit, dp1[i]+dp2[i])
	}
	return maxProfit
}

// dp method
func maxProfit123Dp(prices []int) int {
	l := len(prices)
	if l <= 1 {
		return 0
	}
	dp := make([][][]int, l)
	for i := 0; i < l; i++ {
		dp[i] = make([][]int, 2)
		for j := 0; j <= 1; j++ {
			dp[i][j] = make([]int, 3)
		}
	}
	dp[0][0][0] = 0
	dp[0][1][0] = -prices[0]
	dp[0][0][1] = math.MinInt32
	dp[0][1][1] = math.MinInt32
	dp[0][0][2] = math.MinInt32
	dp[0][1][2] = math.MinInt32
	for i := 1; i < l; i++ {
		price := prices[i]
		dp[i][0][0] = 0
		dp[i][1][0] = utils.Max(dp[i-1][0][0]-price, dp[i-1][1][0])
		dp[i][0][1] = utils.Max(dp[i-1][1][0]+price, dp[i-1][0][1])
		dp[i][1][1] = utils.Max(dp[i-1][0][1]-price, dp[i-1][1][1])
		dp[i][0][2] = utils.Max(dp[i-1][1][1]+price, dp[i-1][0][2])
		dp[i][1][2] = math.MinInt32
	}
	return utils.Max(dp[l-1][0][1], utils.Max(dp[l-1][0][2], 0))
}

func testProblem123() {
	fmt.Println(maxProfit123([]int{3, 3, 5, 0, 0, 3, 1, 4}))
	fmt.Println(maxProfit123([]int{1, 2, 3, 4, 5}))
	fmt.Println(maxProfit123([]int{7, 6, 4, 3, 1}))
}
