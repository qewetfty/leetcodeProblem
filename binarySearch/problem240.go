package binarySearch

import "fmt"

// Write an efficient algorithm that searches for a target value in an m x n integer matrix. The matrix has the following properties:
//	Integers in each row are sorted in ascending from left to right.
//	Integers in each column are sorted in ascending from top to bottom.
//	Example 1:
//		Input: matrix = [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]], target = 5
//		Output: true
//	Example 2:
//		Input: matrix = [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]], target = 20
//		Output: false
//	Constraints:
//		m == matrix.length
//		n == matrix[i].length
//		1 <= n, m <= 300
//		-109 <= matix[i][j] <= 109
//		All the integers in each row are sorted in ascending order.
//		All the integers in each column are sorted in ascending order.
//		-109 <= target <= 109

// 利用矩阵特性从坐下向右上进行遍历
func searchMatrix240(matrix [][]int, target int) bool {
	m, n := len(matrix)-1, 0
	for m >= 0 && n < len(matrix[0]) {
		if matrix[m][n] == target {
			return true
		} else if matrix[m][n] < target {
			n++
		} else {
			m--
		}
	}
	return false
}

// 分治法+二分法
func searchMatrix2(matrix [][]int, target int) bool {
	return divide(matrix, 0, len(matrix)-1, 0, len(matrix[0])-1, target)
}

func divide(matrix [][]int, startRow, endRow, startCol, endCol, target int) bool {
	if startRow >= len(matrix) || endRow < 0 || startCol >= len(matrix[0]) || endCol < 0 || startRow > endRow || startCol > endCol {
		return false
	}
	if startRow == endRow && startCol == endCol {
		return matrix[startRow][startCol] == target
	}
	midRow, midCol := (startRow+endRow)>>1, (startCol+endCol)>>1
	curNum := matrix[midRow][midCol]
	if curNum == target {
		return true
	} else if curNum < target {
		return divide(matrix, midRow+1, endRow, startCol, endCol, target) || divide(matrix, startRow, midRow, midCol+1, endCol, target)
	} else {
		return divide(matrix, startRow, endRow, startCol, midCol-1, target) || divide(matrix, startRow, midRow-1, midCol, endCol, target)
	}
}

func testProblem240() {
	fmt.Println(searchMatrix240([][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}, 5))
	fmt.Println(searchMatrix240([][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}, 20))
}
