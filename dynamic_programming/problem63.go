package dynamic_programming

import "fmt"

// A robot is located at the top-left corner of a m x n grid (marked 'Start' in the diagram below).
// The robot can only move either down or right at any point in time. The robot
// is trying to reach the bottom-right corner of the grid (marked 'Finish' in the
// diagram below).
// Now consider if some obstacles are added to the grids. How many unique paths would there be?
// An obstacle and space is marked as 1 and 0 respectively in the grid.
//	Example 1:
//		Input: obstacleGrid = [[0,0,0],[0,1,0],[0,0,0]]
//		Output: 2
//		Explanation: There is one obstacle in the middle of the 3x3 grid above.
//			There are two ways to reach the bottom-right corner:
//			1. Right -> Right -> Down -> Down
//			2. Down -> Down -> Right -> Right
//	Example 2:
//		Input: obstacleGrid = [[0,1],[0,0]]
//		Output: 1
//	Constraints:
//		m == obstacleGrid.length
//		n == obstacleGrid[i].length
//		1 <= m, n <= 100
//		obstacleGrid[i][j] is 0 or 1.

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if obstacleGrid[0][0] == 1 {
		return 0
	}
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([]int, n)
	dp[0] = 1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[j] = 0
			} else {
				if j > 0 {
					dp[j] = dp[j] + dp[j-1]
				}
			}
		}
	}
	return dp[n-1]
}

func testProblem63() {
	fmt.Println(uniquePathsWithObstacles([][]int{{0, 1}}))
	fmt.Println(uniquePathsWithObstacles([][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}))
	fmt.Println(uniquePathsWithObstacles([][]int{{0, 1}, {0, 0}}))
}
