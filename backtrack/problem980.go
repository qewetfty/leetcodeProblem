package backtrack

import "fmt"

// On a 2-dimensional grid, there are 4 types of squares:
//	1 represents the starting square.  There is exactly one starting square.
//	2 represents the ending square.  There is exactly one ending square.
//	0 represents empty squares we can walk over.
//	-1 represents obstacles that we cannot walk over.
// Return the number of 4-directional walks from the starting square to the ending square, that walk over every
// non-obstacle square exactly once.
//	Example 1:
//		Input: [[1,0,0,0],[0,0,0,0],[0,0,2,-1]]
//		Output: 2
//		Explanation: We have the following two paths:
//			1. (0,0),(0,1),(0,2),(0,3),(1,3),(1,2),(1,1),(1,0),(2,0),(2,1),(2,2)
//			2. (0,0),(1,0),(2,0),(2,1),(1,1),(0,1),(0,2),(0,3),(1,3),(1,2),(2,2)
//	Example 2:
//		Input: [[1,0,0,0],[0,0,0,0],[0,0,0,2]]
//		Output: 4
//		Explanation: We have the following four paths:
//			1. (0,0),(0,1),(0,2),(0,3),(1,3),(1,2),(1,1),(1,0),(2,0),(2,1),(2,2),(2,3)
//			2. (0,0),(0,1),(1,1),(1,0),(2,0),(2,1),(2,2),(1,2),(0,2),(0,3),(1,3),(2,3)
//			3. (0,0),(1,0),(2,0),(2,1),(2,2),(1,2),(1,1),(0,1),(0,2),(0,3),(1,3),(2,3)
//			4. (0,0),(1,0),(2,0),(2,1),(1,1),(0,1),(0,2),(0,3),(1,3),(1,2),(2,2),(2,3)
//	Example 3:
//		Input: [[0,1],[2,0]]
//		Output: 0
//		Explanation:
//			There is no path that walks over every empty square exactly once.
//			Note that the starting and ending square can be anywhere in the grid.
//	Note:
//		1 <= grid.length * grid[0].length <= 20

var views [][]bool
var (
	m          int
	n          int
	validPaths int
	result980  int
	startX     int
	startY     int
)

var dx = []int{0, 0, 1, -1}
var dy = []int{1, -1, 0, 0}

func uniquePathsIII(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m, n = len(grid), len(grid[0])
	validPaths, views, result980 = 0, make([][]bool, m), 0
	for i := 0; i < m; i++ {
		views[i] = make([]bool, n)
	}
	// 确认空方格数量
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				validPaths++
			} else if grid[i][j] == 1 {
				startX, startY = i, j
			}
		}
	}
	views[startX][startY] = true
	backtrack980(grid, startX, startY)
	return result980
}

func backtrack980(grid [][]int, x, y int) {
	if grid[x][y] == 2 {
		if validPaths == 0 {
			result980++
		}
		return
	}
	for k := 0; k < 4; k++ {
		newX, newY := x+dx[k], y+dy[k]
		if 0 <= newX && newX < m && 0 <= newY && newY < n && !views[newX][newY] {
			if grid[newX][newY] == 1 || grid[newX][newY] == -1 {
				continue
			}
			if grid[newX][newY] == 2 {
				backtrack980(grid, newX, newY)
				continue
			}
			views[newX][newY], validPaths = true, validPaths-1
			backtrack980(grid, newX, newY)
			views[newX][newY], validPaths = false, validPaths+1
		}
	}
}

func testProblem980() {
	fmt.Println(uniquePathsIII([][]int{{1, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 2, -1}}))
	fmt.Println(uniquePathsIII([][]int{{1, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 2}}))
	fmt.Println(uniquePathsIII([][]int{{0, 1}, {2, 0}}))
}
