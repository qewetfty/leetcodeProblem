package dfs_bfs

import "fmt"

// Given a 2d grid map of '1's (land) and '0's (water), count the number of islands. An island is surrounded by water
// and is formed by connecting adjacent lands horizontally or vertically. You may assume all four edges of the grid
// are all surrounded by water.
//	Example 1:
//	Input:
//		11110
//		11010
//		11000
//		00000
//	Output: 1
//	Example 2:
//	Input:
//		11000
//		11000
//		00100
//		00011
//	Output: 3

// DFS method
var xgrid = []int{-1, 1, 0, 0}
var ygrid = []int{0, 0, -1, 1}

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	number := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '0' {
				continue
			}
			number = number + sink(i, j, &grid)
		}
	}
	return number
}

func sink(x, y int, grid *[][]byte) int {
	(*grid)[x][y] = '0'
	for i := 0; i < len(xgrid); i++ {
		newX := x + xgrid[i]
		newY := y + ygrid[i]
		if newX >= 0 && newX < len(*grid) && newY >= 0 && newY < len((*grid)[0]) {
			if (*grid)[newX][newY] == '0' {
				continue
			}
			sink(newX, newY, grid)
		}
	}
	return 1
}

func testProblem200() {
	num := [][]byte{{'1', '1', '0', '0', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '1', '0', '0'}, {'0', '0', '0', '1', '1'}}
	fmt.Println(numIslands(num))
}
