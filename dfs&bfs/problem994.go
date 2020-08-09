package dfs_bfs

import (
	"container/list"
	"fmt"
)

// In a given grid, each cell can have one of three values:
//	the value 0 representing an empty cell;
//	the value 1 representing a fresh orange;
//	the value 2 representing a rotten orange.
//	Every minute, any fresh orange that is adjacent (4-directionally) to a rotten orange becomes rotten.
// Return the minimum number of minutes that must elapse until no cell has a fresh orange.  If this is impossible,
// return -1 instead.
//	Example 1:
//		Input: [[2,1,1],[1,1,0],[0,1,1]]
//		Output: 4
//	Example 2:
//		Input: [[2,1,1],[0,1,1],[1,0,1]]
//		Output: -1
//		Explanation:  The orange in the bottom left corner (row 2, column 0) is never rotten, because rotting only happens 4-directionally.
//	Example 3:
//		Input: [[0,2]]
//		Output: 0
//		Explanation:  Since there are already no fresh oranges at minute 0, the answer is just 0.
//	Note:
//		1 <= grid.length <= 10
//		1 <= grid[0].length <= 10
//		grid[i][j] is only 0, 1, or 2.

type point struct {
	x int
	y int
}

func orangesRotting(grid [][]int) int {
	var (
		dx = []int{-1, 1, 0, 0}
		dy = []int{0, 0, 1, -1}
	)
	m, n := len(grid), len(grid[0])
	deque := new(list.List)
	// first add rotten orange into deque
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 2 {
				deque.PushFront(point{
					x: i,
					y: j,
				})
			}
		}
	}
	time := 0
	// bfs rotten orange to make fresh orange rotten
	for deque.Len() > 0 {
		curLen := deque.Len()
		for i := 0; i < curLen; i++ {
			rottenPoint := deque.Remove(deque.Back()).(point)
			curX, curY := rottenPoint.x, rottenPoint.y
			for j := 0; j < 4; j++ {
				newX, newY := curX+dx[j], curY+dy[j]
				if newX >= 0 && newX < m && newY >= 0 && newY < n && grid[newX][newY] == 1 {
					grid[newX][newY] = 2
					deque.PushFront(point{
						x: newX,
						y: newY,
					})
				}
			}
		}
		if deque.Len() > 0 {
			time++
		}
	}
	// if there is still fresh orange, return -1. else return time
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				return -1
			}
		}
	}
	return time
}

func testProblem994() {
	fmt.Println(orangesRotting([][]int{{2, 2}, {1, 1}, {0, 0}, {2, 0}}))
	fmt.Println(orangesRotting([][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}))
	fmt.Println(orangesRotting([][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}}))
	fmt.Println(orangesRotting([][]int{{0, 2}}))
}
