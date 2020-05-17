package main

import "fmt"

// Write an efficient algorithm that searches for a value in an m x n matrix. This matrix has the following properties:
//	Integers in each row are sorted from left to right.
//	The first integer of each row is greater than the last integer of the previous row.
//	Example 1:
//		Input:
//			matrix = [
//			  [1,   3,  5,  7],
//			  [10, 11, 16, 20],
//			  [23, 30, 34, 50]
//			]
//			target = 3
//		Output: true
//	Example 2:
//		Input:
//			matrix = [
//			  [1,   3,  5,  7],
//			  [10, 11, 16, 20],
//			  [23, 30, 34, 50]
//			]
//			target = 13
//		Output: false

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	height := len(matrix)
	width := len(matrix[0])
	lo, hi := 0, height*width-1
	for lo <= hi {
		mid := (lo + hi) / 2
		x := mid / width
		y := mid % width
		num := matrix[x][y]
		if target == num {
			return true
		} else if target < num {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return false
}

func main() {
	fmt.Println(searchMatrix([][]int{{1}}, 2))
}
