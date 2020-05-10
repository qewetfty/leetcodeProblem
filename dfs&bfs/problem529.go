package dfs_bfs

import (
	"container/list"
	"fmt"
)

// Let's play the minesweeper game (Wikipedia, online game)!
// You are given a 2D char matrix representing the game board. 'M' represents an unrevealed mine,
// 'E' represents an unrevealed empty square, 'B' represents a revealed blank square that has no adjacent
// (above, below, left, right, and all 4 diagonals) mines, digit ('1' to '8') represents how many mines
// are adjacent to this revealed square, and finally 'X' represents a revealed mine.
// Now given the next click position (row and column indices) among all the unrevealed squares ('M' or 'E'),
// return the board after revealing this position according to the following rules:
//	If a mine ('M') is revealed, then the game is over - change it to 'X'.
//	If an empty square ('E') with no adjacent mines is revealed, then change it to revealed blank ('B') and all of
//		its adjacent unrevealed squares should be revealed recursively.
//	If an empty square ('E') with at least one adjacent mine is revealed, then change it to a digit ('1' to '8')
//		representing the number of adjacent mines.
//	Return the board when no more squares will be revealed.

//	Example 1:
//		Input:
//			[['E', 'E', 'E', 'E', 'E'],
//			 ['E', 'E', 'M', 'E', 'E'],
//			 ['E', 'E', 'E', 'E', 'E'],
//			 ['E', 'E', 'E', 'E', 'E']]
//			Click : [3,0]
//		Output:
//			[['B', '1', 'E', '1', 'B'],
//			 ['B', '1', 'M', '1', 'B'],
//			 ['B', '1', '1', '1', 'B'],
//			 ['B', 'B', 'B', 'B', 'B']]
//	Example 2:
//		Input:
//			[['B', '1', 'E', '1', 'B'],
//			 ['B', '1', 'M', '1', 'B'],
//			 ['B', '1', '1', '1', 'B'],
//			 ['B', 'B', 'B', 'B', 'B']]
//			Click : [1,2]
//		Output:
//			[['B', '1', 'E', '1', 'B'],
//			 ['B', '1', 'X', '1', 'B'],
//			 ['B', '1', '1', '1', 'B'],
//			 ['B', 'B', 'B', 'B', 'B']]
//	Note:
//		The range of the input matrix's height and width is [1,50].
//		The click position will only be an unrevealed square ('M' or 'E'), which also means the input board
//			contains at least one clickable square.
//		The input board won't be a stage when game is over (some mines have been revealed).
//		For simplicity, not mentioned rules should be ignored in this problem. For example, you don't need to reveal
//			all the unrevealed mines when the game is over, consider any cases that you will win the game or flag any squares.

var xgrids = []int{1, 1, 0, -1, -1, -1, 0, 1}
var ygrids = []int{0, -1, -1, -1, 0, 1, 1, 1}

// DFS method
func updateBoard(board [][]byte, click []int) [][]byte {
	x := click[0]
	y := click[1]
	if board[x][y] == 'M' {
		board[x][y] = 'X'
		return board
	}
	mineCount := 0
	for i := 0; i < len(xgrids); i++ {
		newX := x + xgrids[i]
		newY := y + ygrids[i]
		if newX >= 0 && newX < len(board) && newY >= 0 && newY < len(board[0]) && board[newX][newY] == 'M' {
			mineCount++
		}
	}
	if mineCount > 0 {
		board[x][y] = byte(mineCount) + '0'
		return board
	}
	board[x][y] = 'B'
	for i := 0; i < len(xgrids); i++ {
		newX := x + xgrids[i]
		newY := y + ygrids[i]
		if newX >= 0 && newX < len(board) && newY >= 0 && newY < len(board[0]) && board[newX][newY] == 'E' {
			updateBoard(board, []int{newX, newY})
		}
	}
	return board
}

// BFS method
func updateBoard2(board [][]byte, click []int) [][]byte {
	if board[click[0]][click[1]] == 'M' {
		board[click[0]][click[1]] = 'X'
		return board
	}
	deque := new(list.List)
	deque.PushBack(click)
	for deque.Len() > 0 {
		newClick := deque.Remove(deque.Front()).([]int)
		mineCount := 0
		for i := 0; i < len(xgrids); i++ {
			newX := newClick[0] + xgrids[i]
			newY := newClick[1] + ygrids[i]
			if newX >= 0 && newX < len(board) && newY >= 0 && newY < len(board[0]) && board[newX][newY] == 'M' {
				mineCount++
			}
		}
		if mineCount > 0 {
			board[newClick[0]][newClick[1]] = byte(mineCount) + '0'
			continue
		}
		board[newClick[0]][newClick[1]] = 'B'
		for i := 0; i < len(xgrids); i++ {
			newX := newClick[0] + xgrids[i]
			newY := newClick[1] + ygrids[i]
			if newX >= 0 && newX < len(board) && newY >= 0 && newY < len(board[0]) && board[newX][newY] == 'E' {
				deque.PushBack([]int{newX, newY})
				// to avoid deal the same grid
				board[newX][newY] = 'B'
			}
		}
	}
	return board
}

func testProblem529() {
	fmt.Println('B')
	fmt.Println('E')
	fmt.Println('M')
	fmt.Println(updateBoard2([][]byte{
		{'E', 'E', 'E', 'E', 'E'},
		{'E', 'E', 'M', 'E', 'E'},
		{'E', 'E', 'E', 'E', 'E'},
		{'E', 'E', 'E', 'E', 'E'}}, []int{3, 0}))
}
