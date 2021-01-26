package binarySearch

import (
	"fmt"
	"github.com/leetcodeProblem/utils"
)

// You are a hiker preparing for an upcoming hike. You are given heights, a 2D
// array of size rows x columns, where heights[row][col] represents the height of
// cell (row, col). You are situated in the top-left cell, (0, 0), and you hope
// to travel to the bottom-right cell, (rows-1, columns-1) (i.e., 0-indexed). You
// can move up, down, left, or right, and you wish to find a route that requires
// the minimum effort.
// A route's effort is the maximum absolute difference in heights between two consecutive cells of the route.
// Return the minimum effort required to travel from the top-left cell to the bottom-right cell.
//	Example 1:
//		Input: heights = [[1,2,2],[3,8,2],[5,3,5]]
//		Output: 2
//		Explanation: The route of [1,3,5,3,5] has a maximum absolute difference of 2 in consecutive cells.
//					This is better than the route of [1,2,2,2,5], where the maximum absolute difference is 3.
//	Example 2:
//		Input: heights = [[1,2,3],[3,8,4],[5,3,5]]
//		Output: 1
//		Explanation: The route of [1,2,3,4,5] has a maximum absolute difference of 1 in consecutive cells, which is better than route [1,3,5,3,5].
//	Example 3:
//		Input: heights = [[1,2,1,1,1],[1,2,1,2,1],[1,2,1,2,1],[1,2,1,2,1],[1,1,1,2,1]]
//		Output: 0
//		Explanation: This route does not require any effort.
//	Constraints:
//		rows == heights.length
//		columns == heights[i].length
//		1 <= rows, columns <= 100
//		1 <= heights[i][j] <= 106

var dx = []int{0, 0, 1, -1}
var dy = []int{1, -1, 0, 0}

func minimumEffortPath(heights [][]int) int {
	m, n := len(heights), len(heights[0])
	left, right, result := 0, 999999, 0
	for left <= right {
		mid := (left + right) / 2
		visited, queue := make([]bool, m*n), make([]int, 0)
		queue, visited[0] = append(queue, 0), true
		for len(queue) > 0 {
			cur := queue[0]
			queue = queue[1:]
			x, y := cur/n, cur%n
			for i := 0; i < 4; i++ {
				newX, newY := x+dx[i], y+dy[i]
				if newX >= 0 && newX < m && newY >= 0 && newY < n && !visited[newX*n+newY] && utils.Abs(heights[x][y]-heights[newX][newY]) <= mid {
					queue = append(queue, newX*n+newY)
					visited[newX*n+newY] = true
				}
			}
		}
		if visited[m*n-1] {
			result = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return result
}

func testProblem1631() {
	fmt.Println(minimumEffortPath([][]int{{1, 2, 2}, {3, 8, 2}, {5, 3, 5}}))
	fmt.Println(minimumEffortPath([][]int{{1, 2, 3}, {3, 8, 4}, {5, 3, 5}}))
	fmt.Println(minimumEffortPath([][]int{{1, 2, 1, 1, 1}, {1, 2, 1, 2, 1}, {1, 2, 1, 2, 1}, {1, 2, 1, 2, 1}, {1, 1, 1, 2, 1}}))
}
