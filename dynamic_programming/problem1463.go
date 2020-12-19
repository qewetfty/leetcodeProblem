package dynamic_programming

import (
	"fmt"
	"github.com/leetcodeProblem/utils"
)

// Given a rows x cols matrix grid representing a field of cherries. Each cell in
// grid represents the number of cherries that you can collect.
// You have two robots that can collect cherries for you, Robot #1 is located at
// the top-left corner (0,0) , and Robot #2 is located at the top-right corner
// (0, cols-1) of the grid.
// Return the maximum number of cherries collection using both robots  by following the rules below:
//	From a cell (i,j), robots can move to cell (i+1, j-1) , (i+1, j) or (i+1, j+1).
//	When any robot is passing through a cell, It picks it up all cherries, and the cell becomes an empty cell (0).
//	When both robots stay on the same cell, only one of them takes the cherries.
//	Both robots cannot move outside of the grid at any moment.
//	Both robots should reach the bottom row in the grid.
//	Example 1:
//		Input: grid = [[3,1,1],[2,5,1],[1,5,5],[2,1,1]]
//		Output: 24
//		Explanation: Path of robot #1 and #2 are described in color green and blue respectively.
//					 Cherries taken by Robot #1, (3 + 2 + 5 + 2) = 12.
//					 Cherries taken by Robot #2, (1 + 5 + 5 + 1) = 12.
//					 Total of cherries: 12 + 12 = 24.
//	Example 2:
//		Input: grid = [[1,0,0,0,0,0,1],[2,0,0,0,0,3,0],[2,0,9,0,0,0,0],[0,3,0,5,4,0,0],[1,0,2,3,0,0,6]]
//		Output: 28
//		Explanation: Path of robot #1 and #2 are described in color green and blue respectively.
//					 Cherries taken by Robot #1, (1 + 9 + 5 + 2) = 17.
//					 Cherries taken by Robot #2, (1 + 3 + 4 + 3) = 11.
//					 Total of cherries: 17 + 11 = 28.
//	Example 3:
//		Input: grid = [[1,0,0,3],[0,0,0,3],[0,0,3,3],[9,0,3,3]]
//		Output: 22
//	Example 4:
//		Input: grid = [[1,1],[1,1]]
//		Output: 4
//	Constraints:
//		rows == grid.length
//		cols == grid[i].length
//		2 <= rows, cols <= 70
//		0 <= grid[i][j] <= 100

func cherryPickup(grid [][]int) int {
	row, col := len(grid), len(grid[0])
	dp := make([][][]int, row)
	for i := 0; i < row; i++ {
		dp[i] = make([][]int, col)
		for j := 0; j < col; j++ {
			dp[i][j] = make([]int, col)
			for k := 0; k < col; k++ {
				// 这里赋值为-1表示不可能从此情况出发，排除一些可能
				dp[i][j][k] = -1
			}
		}
	}
	dp[0][0][col-1] = grid[0][0] + grid[0][col-1]
	for i := 1; i < row; i++ {
		for j := 0; j < col; j++ {
			for k := 0; k < col; k++ {
				curGridSum := grid[i][j] + grid[i][k]
				if j == k {
					curGridSum = grid[i][j]
				}
				for dj := j - 1; dj <= j+1; dj++ {
					for dk := k - 1; dk <= k+1; dk++ {
						if dj >= 0 && dj < col && dk >= 0 && dk < col && dp[i-1][dj][dk] >= 0 {
							dp[i][j][k] = utils.Max(dp[i][j][k], dp[i-1][dj][dk]+curGridSum)
						}
					}
				}
			}
		}
	}
	// 计算最大值
	max := 0
	for i := 0; i < col; i++ {
		for j := 0; j < col; j++ {
			if dp[row-1][i][j] > max {
				max = dp[row-1][i][j]
			}
		}
	}
	return max
}

func testProblem1463() {
	fmt.Println(cherryPickup([][]int{{3, 1, 1}, {2, 5, 1}, {1, 5, 5}, {2, 1, 1}}))
	fmt.Println(cherryPickup([][]int{{1, 0, 0, 0, 0, 0, 1}, {2, 0, 0, 0, 0, 3, 0}, {2, 0, 9, 0, 0, 0, 0}, {0, 3, 0, 5, 4, 0, 0}, {1, 0, 2, 3, 0, 0, 6}}))
	fmt.Println(cherryPickup([][]int{{1, 0, 0, 3}, {0, 0, 0, 3}, {0, 0, 3, 3}, {9, 0, 3, 3}}))
	fmt.Println(cherryPickup([][]int{{1, 1}, {1, 1}}))
}
