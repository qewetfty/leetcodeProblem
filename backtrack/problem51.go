package backtrack

import (
	"fmt"
	"math"
	"strings"
)

// The n-queens puzzle is the problem of placing n queens on an n×n chessboard such that no two queens attack each other.
// Given an integer n, return all distinct solutions to the n-queens puzzle.
// Each solution contains a distinct board configuration of the n-queens' placement, where 'Q' and '.' both indicate a queen and an empty space respectively.
//	Example:
//		Input: 4
//		Output: [
//		 [".Q..",  // Solution 1
//		  "...Q",
//		  "Q...",
//		  "..Q."],
//
//		 ["..Q.",  // Solution 2
//		  "Q...",
//		  "...Q",
//		  ".Q.."]
//		]
//	Explanation: There exist two distinct solutions to the 4-queens puzzle as shown above.

func solveNQueens(n int) [][]string {
	res := make([][]string, 0)
	queenPos := make(map[int]int, n)
	data := make([]string, 0)
	for i := 0; i < n; i++ {
		data = append(data, strings.Repeat(".", n))
	}
	backtrackMethod1(n, 0, queenPos, &res)
	return res
}

// method 1
func backtrackMethod1(n int, start int, queenPos map[int]int, res *[][]string) {
	if n == start && len(queenPos) == n {
		data := make([]string, 0)
		for i := 0; i < n; i++ {
			pos := strings.Repeat(".", n)
			q := queenPos[i]
			baseStr := pos[:q] + "Q" + pos[q+1:]
			data = append(data, baseStr)
		}
		*res = append(*res, data)
		return
	}
	for i := start; i < n; i++ {
		for j := 0; j < n; j++ {
			queenPos[i] = j
			if isValid(queenPos, i, j) {
				backtrackMethod1(n, i+1, queenPos, res)
			}
			delete(queenPos, i)
		}
	}
}

func isValid(queenPos map[int]int, level int, currentQueenPos int) bool {
	for i := 0; i < level; i++ {
		currentLevelQueenPos := queenPos[i]
		if currentLevelQueenPos == currentQueenPos || int(math.Abs(float64(i-level))) ==
			int(math.Abs(float64(currentLevelQueenPos-currentQueenPos))) {
			return false
		}
	}
	return true
}

var (
	queens = make(map[int]int)
	col    = make(map[int]int)
	xydif  = make(map[int]int)
	xysum  = make(map[int]int)
)

func solveNQueens2(n int) [][]string {
	res := make([][]string, 0)
	backtrackMethod2(n, 0, &res)
	return res
}

func backtrackMethod2(n int, start int, res *[][]string) {
	if len(col) == n {
		data := make([]string, 0)
		for i := 0; i < n; i++ {
			q := queens[i]
			pos := strings.Repeat(".", q) + "Q" + strings.Repeat(".", n-q-1)
			data = append(data, pos)
		}
		*res = append(*res, data)
		return
	}
	for i := start; i < n; i++ {
		for j := 0; j < n; j++ {
			if isNotUnderAttack(i, j) {
				placeQueen(i, j)
				backtrackMethod2(n, i+1, res)
				removeQueen(i, j)
			}
		}
	}
}

func placeQueen(x, y int) {
	queens[x] = y
	col[y] = x
	xydif[x-y] = x
	xysum[x+y] = x
}

func removeQueen(x, y int) {
	delete(queens, x)
	delete(col, y)
	delete(xydif, x-y)
	delete(xysum, x+y)
}

func isNotUnderAttack(x, y int) bool {
	if _, ok := col[y]; ok {
		return false
	}
	if _, ok := xydif[x-y]; ok {
		return false
	}
	if _, ok := xysum[x+y]; ok {
		return false
	}
	return true
}

func testProblem51() {
	fmt.Println(solveNQueens2(1))
	fmt.Println(solveNQueens2(2))
	fmt.Println(solveNQueens2(3))
	fmt.Println(solveNQueens2(4))
	fmt.Println(solveNQueens2(8))
}
