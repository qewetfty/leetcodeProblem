package dynamic_programming

import (
	"fmt"
	"github.com/leetcodeProblem/utils"
	"math"
)

// Say you have an array for which the ith element is the price of a given stock on day i.
// Design an algorithm to find the maximum profit. You may complete as many transactions as you like
// (ie, buy one and sell one share of the stock multiple times) with the following restrictions:
//	You may not engage in multiple transactions at the same time (ie, you must sell the stock before you buy again).
//	After you sell your stock, you cannot buy stock on next day. (ie, cooldown 1 day)
//	Example:
//		Input: [1,2,3,0,2]
//		Output: 3
//		Explanation: transactions = [buy, sell, cooldown, buy, sell]

func maxProfit309(prices []int) int {
	l := len(prices)
	if l <= 1 {
		return 0
	}
	dp := make([][][]int, l)
	for i := 0; i < l; i++ {
		dp[i] = make([][]int, 2)
		for j := 0; j <= 1; j++ {
			dp[i][j] = make([]int, 2)
		}
	}
	dp[0][0][0] = 0
	dp[0][1][0] = -prices[0]
	dp[0][0][1] = math.MinInt32
	dp[0][1][1] = math.MinInt32
	for i := 1; i < l; i++ {
		dp[i][0][0] = utils.Max(dp[i-1][0][0], dp[i-1][0][1])
		dp[i][0][1] = dp[i-1][1][0] + prices[i]
		dp[i][1][0] = utils.Max(dp[i-1][0][0]-prices[i], dp[i-1][1][0])
		dp[i][1][1] = math.MinInt32
	}
	return utils.Max(dp[l-1][0][1], dp[l-1][0][0])
}

func testProblem309() {
	fmt.Println(maxProfit309([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 5, 7, 88, 10, 2}))
	fmt.Println(maxProfit309([]int{1, 2, 3, 0, 2}))
}
