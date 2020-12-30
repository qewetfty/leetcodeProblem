package array

import "fmt"

// According to Wikipedia's article: "The Game of Life, also known simply as
// Life, is a cellular automaton devised by the British mathematician John Horton
// Conway in 1970."
// The board is made up of an m x n grid of cells, where each cell has an initial
// state: live (represented by a 1) or dead (represented by a 0). Each cell
// interacts with its eight neighbors (horizontal, vertical, diagonal) using the
// following four rules (taken from the above Wikipedia article):
//	Any live cell with fewer than two live neighbors dies as if caused by under-population.
//	Any live cell with two or three live neighbors lives on to the next generation.
//	Any live cell with more than three live neighbors dies, as if by over-population.
//	Any dead cell with exactly three live neighbors becomes a live cell, as if by reproduction.
//	The next state is created by applying the above rules simultaneously to every cell in the current state, where births
//		and deaths occur simultaneously. Given the current state of the m x n grid board, return the next state.
//	Example 1:
//		Input: board = [[0,1,0],[0,0,1],[1,1,1],[0,0,0]]
//		Output: [[0,0,0],[1,0,1],[0,1,1],[0,1,0]]
//	Example 2:
//		Input: board = [[1,1],[1,0]]
//		Output: [[1,1],[1,1]]
//	Constraints:
//		m == board.length
//		n == board[i].length
//		1 <= m, n <= 25
//		board[i][j] is 0 or 1.
//	Follow up:
//		Could you solve it in-place? Remember that the board needs to be updated
//		simultaneously: You cannot update some cells first and then use their updated
//		values to update other cells. In this question, we represent the board using a
//		2D array.
//		In principle, the board is infinite, which would cause problems when
//		the active area encroaches upon the border of the array (i.e., live cells
//		reach the border). How would you address these problems?

// 如果空间复杂度为O(1)，则需要记录以前的状态。这里通过修改原始board来完成。
// 如果原始为0： a. 如果0旁边恰海有3个为1的，状态置为-1，
//             b. 其他情况则仍旧保持0
// 这样只需判断是否小于等于0即可确认原始状态。
// 如果原始为1： a. 如果1旁边小于2个1，状态置为2
//             b. 如果1旁边是2个或者3个1，状态不变，仍为1
//             c. 如果1旁边大于3个1，状态置为3
// 这样大于等于1则表示原始状态为1。
// 最终通过判断每个格子的数值来确认最终的结果

var dx289 = []int{1, 1, 0, -1, -1, -1, 0, 1}
var dy289 = []int{0, 1, 1, 1, 0, -1, -1, -1}

func gameOfLife(board [][]int) {
	m, n := len(board), len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			board[i][j] = judge(board, m, n, i, j)
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			switch board[i][j] {
			case 2, 3:
				board[i][j] = 0
			case -1:
				board[i][j] = 1
			}
		}
	}
}

func judge(board [][]int, m, n, i, j int) int {
	one := 0
	for k := 0; k < 8; k++ {
		newI, newJ := i+dx289[k], j+dy289[k]
		if newI >= 0 && newI < m && newJ >= 0 && newJ < n {
			if board[newI][newJ] >= 1 {
				one++
			}
		}
	}
	// 返回最终结果
	if board[i][j] == 0 {
		if one == 3 {
			return -1
		} else {
			return 0
		}
	}
	if one < 2 {
		return 2
	}
	if one > 3 {
		return 3
	}
	return 1
}

func testProblem289() {
	a := [][]int{{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0}}
	gameOfLife(a)
	fmt.Println(a)
}
