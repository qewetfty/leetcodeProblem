package dfs_bfs

import (
	"fmt"
	"github.com/leetcodeProblem/utils"
)

// You are given an m x n binary matrix grid. An island is a group of 1's
// (representing land) connected 4-directionally (horizontal or vertical.) You
// may assume all four edges of the grid are surrounded by water.
// The area of an island is the number of cells with a value 1 in the island.
// Return the maximum area of an island in grid. If there is no island, return 0.
//	Example 1:
//		Input: grid = [[0,0,1,0,0,0,0,1,0,0,0,0,0],[0,0,0,0,0,0,0,1,1,1,0,0,0],[0,1,1,0,1,0,0,0,0,0,0,0,0],[0,1,0,0,1,1,0,0,1,0,1,0,0],[0,1,0,0,1,1,0,0,1,1,1,0,0],[0,0,0,0,0,0,0,0,0,0,1,0,0],[0,0,0,0,0,0,0,1,1,1,0,0,0],[0,0,0,0,0,0,0,1,1,0,0,0,0]]
//		Output: 6
//		Explanation: The answer is not 11, because the island must be connected 4-directionally.
//	Example 2:
//		Input: grid = [[0,0,0,0,0,0,0,0]]
//		Output: 0
//	Constraints:
//		m == grid.length
//		n == grid[i].length
//		1 <= m, n <= 50
//		grid[i][j] is either 0 or 1.

func maxAreaOfIsland(grid [][]int) int {
	result := 0
	m, n := len(grid), len(grid[0])
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			result = utils.Max(result, dfs695(visited, grid, i, j))
		}
	}
	return result
}

func dfs695(visited [][]bool, grid [][]int, i, j int) int {
	if visited[i][j] || grid[i][j] == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	visited[i][j] = true
	result := 1
	for k := 0; k < 4; k++ {
		newI, newJ := i+xgrid[k], j+ygrid[k]
		if newI >= 0 && newI < m && newJ >= 0 && newJ < n {
			result = result + dfs695(visited, grid, newI, newJ)
		}
	}
	return result
}

func testProblem695() {
	fmt.Println(maxAreaOfIsland([][]int{{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0}, {0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0}, {0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0}, {0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0}}))
}
