package dynamic_programming

import (
	"fmt"
	"github.com/leetcodeProblem/utils"
)

// Given a 2D binary matrix filled with 0's and 1's, find the largest square containing only 1's and return its area.
//	Example:
//	Input:
//		1 0 1 0 0
//		1 0 1 1 1
//		1 1 1 1 1
//		1 0 0 1 0
//	Output: 4

func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	m := len(matrix)
	n := len(matrix[0])
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}
	max := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '0' {
				dp[i+1][j+1] = 0
			} else {
				dp[i+1][j+1] = utils.Min(dp[i][j], utils.Min(dp[i+1][j], dp[i][j+1])) + 1
				max = utils.Max(max, dp[i+1][j+1])
			}
		}
	}
	return max * max
}

func testProblem221() {
	fmt.Println(maximalSquare([][]byte{{'0'}}))
	fmt.Println(maximalSquare([][]byte{{'1', '0', '1', '0', '0'}, {'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'}, {'1', '0', '0', '1', '0'}}))
}
