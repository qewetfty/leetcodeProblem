package dynamic_programming

import (
	"fmt"
	"github.com/leetcodeProblem/utils"
)

// Given a m x n grid filled with non-negative numbers, find a path from top left to bottom right which minimizes the
// sum of all numbers along its path.
//	Note: You can only move either down or right at any point in time.
//	Example:
//		Input:
//		[
//		  [1,3,1],
//		  [1,5,1],
//		  [4,2,1]
//		]
//		Output: 7
//		Explanation: Because the path 1→3→1→1→1 minimizes the sum.

// 2-dimension dp method
func minPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m := len(grid)
	n := len(grid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		dp[i][0] = grid[i][0]
		if i > 0 {
			dp[i][0] = dp[i][0] + dp[i-1][0]
		}
	}
	for i := 0; i < n; i++ {
		dp[0][i] = grid[0][i]
		if i > 0 {
			dp[0][i] = dp[0][i-1] + dp[0][i]
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = utils.Min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[m-1][n-1]
}

func minPathSum2(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m := len(grid)
	n := len(grid[0])
	dp := make([]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if j == 0 {
				dp[j] = dp[j] + grid[i][j]
			} else {
				if i == 0 {
					dp[j] = dp[j-1] + grid[i][j]
				} else {
					dp[j] = utils.Min(dp[j-1], dp[j]) + grid[i][j]
				}
			}
		}
	}
	return dp[n-1]
}

func testProblem64() {
	fmt.Println(minPathSum2([][]int{{1, 2, 5}, {3, 2, 1}}))
	fmt.Println(minPathSum2([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}))
}
