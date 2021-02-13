package dfs_bfs

import "fmt"

// In an N by N square grid, each cell is either empty (0) or blocked (1).
// A clear path from top-left to bottom-right has length k if and only if it is composed of cells C_1, C_2, ..., C_k such that:
//		Adjacent cells C_i and C_{i+1} are connected 8-directionally (ie., they are different and share an edge or corner)
//		C_1 is at location (0, 0) (ie. has value grid[0][0])
//		C_k is at location (N-1, N-1) (ie. has value grid[N-1][N-1])
//		If C_i is located at (r, c), then grid[r][c] is empty (ie. grid[r][c] == 0).
// Return the length of the shortest such clear path from top-left to bottom-right.  If such a path does not exist, return -1.
//	Example 1:
//		Input: [[0,1],[1,0]]
//		Output: 2
//	Example 2:
//		Input: [[0,0,0],[1,1,0],[1,1,0]]
//		Output: 4
//	Note:
//		1 <= grid.length == grid[0].length <= 100
//		grid[r][c] is 0 or 1

var dx1091 = []int{1, 1, 0, -1, -1, -1, 0, 1}
var dy1091 = []int{0, 1, 1, 1, 0, -1, -1, -1}

func shortestPathBinaryMatrix(grid [][]int) int {
	if grid[0][0] == 1 {
		return -1
	}
	m, n := len(grid), len(grid[0])
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	queue := make([]int, 0)
	queue = append(queue, 0)
	path := 0
	for len(queue) > 0 {
		path++
		l := len(queue)
		for i := 0; i < l; i++ {
			cur := queue[i]
			x, y := cur/n, cur%n
			if x == n-1 && y == n-1 && grid[x][y] == 0 {
				return path
			}
			for j := 0; j < 8; j++ {
				newX, newY := x+dx1091[j], y+dy1091[j]
				if newX >= 0 && newX < m && newY >= 0 && newY < n && !visited[newX][newY] {
					visited[newX][newY] = true
					if grid[newX][newY] == 0 {
						queue = append(queue, newX*n+newY)
					}
				}
			}
		}
		queue = queue[l:]
	}
	return -1
}

func testProblem1091() {
	fmt.Println(shortestPathBinaryMatrix([][]int{{0, 1}, {1, 0}}))
	fmt.Println(shortestPathBinaryMatrix([][]int{{0, 0, 0}, {1, 1, 0}, {1, 1, 0}}))
}
