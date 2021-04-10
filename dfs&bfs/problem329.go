package dfs_bfs

import (
	"fmt"
	"github.com/leetcodeProblem/utils"
)

// Given an m x n integers matrix, return the length of the longest increasing path in matrix.
// From each cell, you can either move in four directions: left, right, up, or
// down. You may not move diagonally or move outside the boundary (i.e.,
// wrap-around is not allowed).
//	Example 1:
//		Input: matrix = [[9,9,4],[6,6,8],[2,1,1]]
//		Output: 4
//		Explanation: The longest increasing path is [1, 2, 6, 9].
//	Example 2:
//		Input: matrix = [[3,4,5],[3,2,6],[2,2,1]]
//		Output: 4
//		Explanation: The longest increasing path is [3, 4, 5, 6]. Moving diagonally is not allowed.
//	Example 3:
//		Input: matrix = [[1]]
//		Output: 1
//	Constraints:
//		m == matrix.length
//		n == matrix[i].length
//		1 <= m, n <= 200
//		0 <= matrix[i][j] <= 231 - 1

var dx329 = []int{0, 0, 1, -1}
var dy329 = []int{1, -1, 0, 0}

func longestIncreasingPath(matrix [][]int) int {
	if len(matrix) == 0 && len(matrix[0]) == 0 {
		return 0
	}
	m, n := len(matrix), len(matrix[0])
	paths := make([][]int, m)
	for i := 0; i < m; i++ {
		paths[i] = make([]int, n)
	}
	result := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			result = utils.Max(result, dfs329(matrix, i, j, paths))
		}
	}
	return result
}

func dfs329(matrix [][]int, i, j int, paths [][]int) int {
	if paths[i][j] != 0 {
		return paths[i][j]
	}
	paths[i][j]++
	for k := 0; k < 4; k++ {
		newI, newJ := i+dx329[k], j+dy329[k]
		if newI >= 0 && newI < len(matrix) && newJ >= 0 && newJ < len(matrix[0]) && matrix[newI][newJ] > matrix[i][j] {
			paths[i][j] = utils.Max(paths[i][j], dfs329(matrix, newI, newJ, paths)+1)
		}
	}
	return paths[i][j]
}

func testProblem329() {
	fmt.Println(longestIncreasingPath([][]int{{9, 9, 4}, {6, 6, 8}, {2, 1, 1}}))
	fmt.Println(longestIncreasingPath([][]int{{3, 4, 5}, {3, 2, 6}, {2, 2, 1}}))
	fmt.Println(longestIncreasingPath([][]int{{1}}))
}
