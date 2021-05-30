package backtrack

import "fmt"

// The n-queens puzzle is the problem of placing n queens on an n x n chessboard such that no two queens attack each other.
// Given an integer n, return the number of distinct solutions to the n-queens puzzle.
//	Example 1:
//		Input: n = 4
//		Output: 2
//		Explanation: There are two distinct solutions to the 4-queens puzzle as shown.
//	Example 2:
//		Input: n = 1
//		Output: 1
//	Constraints:
//		1 <= n <= 9

var (
	result52 int
)

// 回溯法
func totalNQueens(n int) int {
	result52 = 0
	queens, col, xysum, xydif = make(map[int]int), make(map[int]int), make(map[int]int), make(map[int]int)
	backtrack52(n, 0)
	return result52
}

func backtrack52(n, start int) {
	if len(col) == n {
		result52++
		return
	}
	for i := start; i < n; i++ {
		for j := 0; j < n; j++ {
			if isNotUnderAttack(i, j) {
				placeQueen(i, j)
				backtrack52(n, i+1)
				removeQueen(i, j)
			}
		}
	}
}

// 打表法
var resultMap = []int{1, 0, 0, 2, 10, 4, 40, 92, 352}

func totalNQueens2(n int) int {
	return resultMap[n-1]
}

func testProblem52() {
	fmt.Println(totalNQueens(1))
	fmt.Println(totalNQueens(2))
	fmt.Println(totalNQueens(3))
	fmt.Println(totalNQueens(4))
	fmt.Println(totalNQueens(5))
	fmt.Println(totalNQueens(6))
	fmt.Println(totalNQueens(7))
	fmt.Println(totalNQueens(8))
	fmt.Println(totalNQueens(9))
}
