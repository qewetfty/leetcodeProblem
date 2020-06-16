package dynamic_programming

import (
	"fmt"
	"github.com/leetcodeProblem/utils"
	"math"
)

// Say you have an array for which the ith element is the price of a given stock on day i.
// If you were only permitted to complete at most one transaction (i.e., buy one and sell one share of the stock),
// design an algorithm to find the maximum profit.
// Note that you cannot sell a stock before you buy one.
//	Example 1:
//		Input: [7,1,5,3,6,4]
//		Output: 5
//		Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
//		             Not 7-1 = 6, as selling price needs to be larger than buying price.
//	Example 2:
//		Input: [7,6,4,3,1]
//		Output: 0
//		Explanation: In this case, no transaction is done, i.e. max profit = 0.

// dp method
func maxProfit(prices []int) int {
	l := len(prices)
	if l <= 1 {
		return 0
	}
	dp := make([]int, l)
	dp[0] = 0
	min := prices[0]
	for i := 1; i < l; i++ {
		dp[i] = utils.Max(dp[i-1], prices[i]-min)
		if prices[i] < min {
			min = prices[i]
		}
	}
	return dp[l-1]
}

func maxProfit2(prices []int) int {
	maxProfit, min := 0, math.MaxInt32
	for _, price := range prices {
		if price < min {
			min = price
		} else if price-min > maxProfit {
			maxProfit = price - min
		}
	}
	return maxProfit
}

func testProblem121() {
	fmt.Println(maxProfit2([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfit2([]int{7, 6, 4, 3, 1}))
}
