package divideAndConquer

import "fmt"

// A robot is located at the top-left corner of a m x n grid (marked 'Start' in the diagram below).
// The robot can only move either down or right at any point in time. The robot is trying to reach the bottom-right
// corner of the grid (marked 'Finish' in the diagram below).
// How many possible unique paths are there?
//	Example 1:
//		Input: m = 3, n = 2
//		Output: 3
//		Explanation:
//			From the top-left corner, there are a total of 3 ways to reach the bottom-right corner:
//				1. Right -> Right -> Down
//				2. Right -> Down -> Right
//				3. Down -> Right -> Right
//	Example 2:
//		Input: m = 7, n = 3
//		Output: 28
//	Constraints:=
//		1 <= m, n <= 100
//		It's guaranteed that the answer will be less than or equal to 2 * 10 ^ 9.

var paths [][]int

func uniquePaths(m int, n int) int {
	paths = make([][]int, m)
	for i := 0; i < m; i++ {
		paths[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		paths[i][0] = 1
	}
	for j := 0; j < n; j++ {
		paths[0][j] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			paths[i][j] = paths[i-1][j] + paths[i][j-1]
		}
	}
	return paths[m-1][n-1]
}

func testProblem62() {
	fmt.Println(uniquePaths(3, 2))
	fmt.Println(uniquePaths(7, 3))
	fmt.Println(uniquePaths(1, 1))
}
