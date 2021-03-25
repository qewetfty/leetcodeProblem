package dfs_bfs

import "fmt"

// Given an m x n matrix of non-negative integers representing the height of each
// unit cell in a continent, the "Pacific ocean" touches the left and top edges
// of the matrix and the "Atlantic ocean" touches the right and bottom edges.
// Water can only flow in four directions (up, down, left, or right) from a cell to another one with height equal or lower.
// Find the list of grid coordinates where water can flow to both the Pacific and Atlantic ocean.
//	Note:
//		The order of returned grid coordinates does not matter.
//		Both m and n are less than 150.
//	Example:
//		Given the following 5x5 matrix:
//  		Pacific ~   ~   ~   ~   ~
//  		     ~  1   2   2   3  (5) *
//  		     ~  3   2   3  (4) (4) *
//  		     ~  2   4  (5)  3   1  *
//  		     ~ (6) (7)  1   4   5  *
//  		     ~ (5)  1   1   2   4  *
//  		        *   *   *   *   * Atlantic
//	Return:
//		[[0, 4], [1, 3], [1, 4], [2, 2], [3, 0], [3, 1], [4, 0]] (positions with parentheses in above matrix).

var dx417 = []int{1, 0, -1, 0}
var dy417 = []int{0, 1, 0, -1}

func pacificAtlantic(matrix [][]int) [][]int {
	result := make([][]int, 0)
	m := len(matrix)
	if m <= 0 {
		return result
	}
	n := len(matrix[0])
	if n <= 0 {
		return result
	}
	canGetPa, canGetAt := make([][]bool, m), make([][]bool, m)
	for i := 0; i < m; i++ {
		canGetPa[i], canGetAt[i] = make([]bool, n), make([]bool, n)
	}
	// 遍历左和右
	for i := 0; i < m; i++ {
		dfs417(matrix, &canGetPa, i, 0)
		dfs417(matrix, &canGetAt, i, n-1)
	}
	// 遍历上和下
	for j := 0; j < n; j++ {
		dfs417(matrix, &canGetPa, 0, j)
		dfs417(matrix, &canGetAt, m-1, j)
	}
	// 汇总结果
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if canGetPa[i][j] && canGetAt[i][j] {
				result = append(result, []int{i, j})
			}
		}
	}
	return result
}

func dfs417(matrix [][]int, visited *[][]bool, x, y int) {
	if !(*visited)[x][y] {
		(*visited)[x][y] = true
		for i := 0; i < 4; i++ {
			newX, newY := x+dx417[i], y+dy417[i]
			if newX >= 0 && newX < len(matrix) && newY >= 0 && newY < len(matrix[0]) && matrix[x][y] <= matrix[newX][newY] {
				dfs417(matrix, visited, newX, newY)
			}
		}
	}
}

func testProblem417() {
	fmt.Println(pacificAtlantic([][]int{{1, 2, 2, 3, 5}, {3, 2, 3, 4, 4}, {2, 4, 5, 3, 1}, {6, 7, 1, 4, 5}, {5, 1, 1, 2, 4}}))
}
