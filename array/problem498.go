package array

import "fmt"

// Given a matrix of M x N elements (M rows, N columns), return all elements of
// the matrix in diagonal order as shown in the below image.
//	Example:
//	Input:
//		[
//		 [ 1, 2, 3 ],
//		 [ 4, 5, 6 ],
//		 [ 7, 8, 9 ]
//		]
//	Output:  [1,2,4,7,5,3,6,8,9]
//	Explanation:
//	Note:
//		The total number of elements of the given matrix will not exceed 10,000.

var dx = []int{-1, 1}
var dy = []int{1, -1}

func findDiagonalOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	m, n := len(matrix), len(matrix[0])
	num := m * n
	result := make([]int, 0)
	x, y, step := 0, 0, 0
	for len(result) < num {
		move := step % 2
		for x >= 0 && x < m && y >= 0 && y < n {
			result = append(result, matrix[x][y])
			x, y = x+dx[move], y+dy[move]
		}
		if x < 0 && y >= n {
			x, y = 1, n-1
		} else if y < 0 && x >= m {
			x, y = m-1, 1
		} else if x < 0 {
			x = 0
		} else if x >= m {
			x, y = m-1, y+2
		} else if y < 0 {
			y = 0
		} else if y >= n {
			x, y = x+2, n-1
		}
		step++
	}
	return result
}

func testProblem498() {
	fmt.Println(findDiagonalOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
	fmt.Println(findDiagonalOrder([][]int{{1, 2}, {4, 5}, {7, 8}}))
	fmt.Println(findDiagonalOrder([][]int{{1}}))
	fmt.Println(findDiagonalOrder([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}))
	fmt.Println(findDiagonalOrder([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}))
}
