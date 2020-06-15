package dynamic_programming

import (
	"fmt"
	"github.com/leetcodeProblem/utils"
)

// You are given coins of different denominations and a total amount of money amount. Write a function to compute the
// fewest number of coins that you need to make up that amount. If that amount of money cannot be made up by any
// combination of the coins, return -1.
//	Example 1:
//		Input: coins = [1, 2, 5], amount = 11
//		Output: 3
//		Explanation: 11 = 5 + 5 + 1
//	Example 2:
//		Input: coins = [2], amount = 3
//		Output: -1

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = amount + 1
	}
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if coin <= i {
				dp[i] = utils.Min(dp[i-coin]+1, dp[i])
			}
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

func testProblem322() {
	fmt.Println(coinChange([]int{1, 2, 5}, 11))
	fmt.Println(coinChange([]int{2}, 3))
}
