package dynamic_programming

import (
	"fmt"
	"github.com/leetcodeProblem/utils"
)

// Your are given an array of integers prices, for which the i-th element is the price of a given stock on day i;
// and a non-negative integer fee representing a transaction fee.
// You may complete as many transactions as you like, but you need to pay the transaction fee for each transaction.
// You may not buy more than 1 share of a stock at a time (ie. you must sell the stock share before you buy again.)
// Return the maximum profit you can make.
//	Example 1:
//		Input: prices = [1, 3, 2, 8, 4, 9], fee = 2
//		Output: 8
//		Explanation: The maximum profit can be achieved by:
//		Buying at prices[0] = 1
//		Selling at prices[3] = 8
//		Buying at prices[4] = 4
//		Selling at prices[5] = 9
//		The total profit is ((8 - 1) - 2) + ((9 - 4) - 2) = 8.
//	Note:
//		0 < prices.length <= 50000.
//		0 < prices[i] < 50000.
//		0 <= fee < 50000.

// 2-dimension dp method
func maxProfit714Dp(prices []int, fee int) int {
	l := len(prices)
	if l <= 1 {
		return 0
	}
	dp := make([][]int, l)
	for i := 0; i < l; i++ {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = 0
	dp[0][1] = -prices[0]
	for i := 1; i < l; i++ {
		dp[i][0] = utils.Max(dp[i-1][1]+prices[i]-fee, dp[i-1][0])
		dp[i][1] = utils.Max(dp[i-1][0]-prices[i], dp[i-1][1])
	}
	return dp[l-1][0]
}

// 1-dimension method
func maxProfit714(prices []int, fee int) int {
	l := len(prices)
	if l <= 1 {
		return 0
	}
	dp := make([]int, 2)
	dp[0] = 0
	dp[1] = -prices[0]
	for i := 1; i < l; i++ {
		dp[0], dp[1] = utils.Max(dp[1]+prices[i]-fee, dp[0]), utils.Max(dp[0]-prices[i], dp[1])
	}
	return dp[0]
}

func testProblem714() {
	fmt.Println(maxProfit714([]int{1, 3, 2, 8, 4, 9}, 2))
}
