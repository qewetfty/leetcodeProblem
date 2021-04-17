package dynamic_programming

import "fmt"

// Given a matrix and a target, return the number of non-empty submatrices that sum to target.
// A submatrix x1, y1, x2, y2 is the set of all cells matrix[x][y] with x1 <= x <= x2 and y1 <= y <= y2.
// Two submatrices (x1, y1, x2, y2) and (x1', y1', x2', y2') are different if they have some coordinate that is different: for example, if x1 != x1'.
//	Example 1:
//		Input: matrix = [[0,1,0],[1,1,1],[0,1,0]], target = 0
//		Output: 4
//		Explanation: The four 1x1 submatrices that only contain 0.
//	Example 2:
//		Input: matrix = [[1,-1],[-1,1]], target = 0
//		Output: 5
//		Explanation: The two 1x2 submatrices, plus the two 2x1 submatrices, plus the 2x2 submatrix.
//	Example 3:
//		Input: matrix = [[904]], target = 0
//		Output: 0
//	Constraints:
//		1 <= matrix.length <= 100
//		1 <= matrix[0].length <= 100
//		-1000 <= matrix[i] <= 1000
//		-10^8 <= target <= 10^8

func numSubmatrixSumTarget(matrix [][]int, target int) int {
	m, n := len(matrix), len(matrix[0])
	sumPrefix := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		sumPrefix[i] = make([]int, n+1)
	}
	for i := 0; i < m; i++ {
		t := 0
		for j := 0; j < n; j++ {
			t = t + matrix[i][j]
			sumPrefix[i+1][j+1] = t + sumPrefix[i][j+1]
		}
	}
	result := 0
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			for k := 1; k <= i; k++ {
				for p := 1; p <= j; p++ {
					curSum := sumPrefix[i][j] - sumPrefix[k-1][j] - sumPrefix[i][p-1] + sumPrefix[k-1][p-1]
					if curSum == target {
						result++
					}
				}
			}
		}
	}
	return result
}

func testProblem1074() {
	fmt.Println(numSubmatrixSumTarget([][]int{{0, 1, 0}, {1, 1, 1}, {0, 1, 0}}, 0))
	fmt.Println(numSubmatrixSumTarget([][]int{{1, -1}, {-1, 1}}, 0))
	fmt.Println(numSubmatrixSumTarget([][]int{{111}}, 0))
}
