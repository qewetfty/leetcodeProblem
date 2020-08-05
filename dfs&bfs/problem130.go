package dfs_bfs

import "fmt"

// Given a 2D board containing 'X' and 'O' (the letter O), capture all regions surrounded by 'X'.
//A region is captured by flipping all 'O's into 'X's in that surrounded region.
//	Example:
//		X X X X
//		X O O X
//		X X O X
//		X O X X
//	After running your function, the board should be:
//		X X X X
//		X X X X
//		X X X X
//		X O X X
//	Explanation:
//		Surrounded regions shouldn’t be on the border, which means that any 'O' on the border of the board are not
//		flipped to 'X'. Any 'O' that is not on the border and it is not connected to an 'O' on the border will be
//		flipped to 'X'. Two cells are connected if they are adjacent cells connected horizontally or vertically.

var (
	dx = []int{1, -1, 0, 0}
	dy = []int{0, 0, 1, -1}
)

// 搜索和边界'O'相连的'O',特殊标记为'#'.待最终扫描时，将'#'变为'O',将'O'变为'X'
func solve(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}
	m, n := len(board), len(board[0])
	if m <= 2 || n <= 2 {
		return
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// 判断是否在边缘
			edge := i == 0 || i == m-1 || j == 0 || j == n-1
			if edge && board[i][j] == 'O' {
				dfs(board, i, j)
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == '#' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
}

func dfs(board [][]byte, x, y int) {
	if x < 0 || x >= len(board) || y < 0 || y >= len(board[0]) || board[x][y] == 'X' || board[x][y] == '#' {
		return
	}
	board[x][y] = '#'
	for i := 0; i < 4; i++ {
		newX, newY := x+dx[i], y+dy[i]
		dfs(board, newX, newY)
	}
}

func testProblem130() {
	a := [][]byte{{'X', 'X', 'X', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'X', 'O', 'X'}, {'X', 'O', 'X', 'X'}}
	fmt.Println(a)
	solve(a)
	fmt.Println(a)
}
