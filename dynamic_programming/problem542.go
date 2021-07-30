package dynamic_programming

import (
	"fmt"
	"github.com/leetcodeProblem/utils"
	"math"
)

// Given an m x n binary matrix mat, return the distance of the nearest 0 for each cell.
// The distance between two adjacent cells is 1.
//	Example 1:
//		Input: mat = [[0,0,0],[0,1,0],[0,0,0]]
//		Output: [[0,0,0],[0,1,0],[0,0,0]]
//	Example 2:
//		Input: mat = [[0,0,0],[0,1,0],[1,1,1]]
//		Output: [[0,0,0],[0,1,0],[1,2,1]]
//	Constraints:
//		m == mat.length
//		n == mat[i].length
//		1 <= m, n <= 104
//		1 <= m * n <= 104
//		mat[i][j] is either 0 or 1.
//		There is at least one 0 in mat.

func updateMatrix(mat [][]int) [][]int {
	m, n := len(mat), len(mat[0])
	distance := make([][]int, m)
	for i := 0; i < m; i++ {
		distance[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if mat[i][j] == 0 {
				distance[i][j] = 0
			} else {
				distance[i][j] = math.MaxInt32
			}
		}
	}
	//向左上遍历
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i-1 >= 0 {
				distance[i][j] = utils.Min(distance[i][j], distance[i-1][j]+1)
			}
			if j-1 >= 0 {
				distance[i][j] = utils.Min(distance[i][j], distance[i][j-1]+1)
			}
		}
	}

	// 向右下遍历
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i+1 < m {
				distance[i][j] = utils.Min(distance[i][j], distance[i+1][j]+1)
			}
			if j+1 < n {
				distance[i][j] = utils.Min(distance[i][j], distance[i][j+1]+1)
			}
		}
	}
	return distance
}

func testProblem542() {
	fmt.Println(updateMatrix([][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}))
	fmt.Println(updateMatrix([][]int{{0, 0, 0}, {0, 1, 0}, {1, 1, 1}}))
}
