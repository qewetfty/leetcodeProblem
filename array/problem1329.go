package array

import (
	"fmt"
	"sort"
)

// A matrix diagonal is a diagonal line of cells starting from some cell in
// either the topmost row or leftmost column and going in the bottom-right
// direction until reaching the matrix's end. For example, the matrix diagonal
// starting from mat[2][0], where mat is a 6 x 3 matrix, includes cells
// mat[2][0], mat[3][1], and mat[4][2].
// Given an m x n matrix mat of integers, sort each matrix diagonal in ascending order and return the resulting matrix.
//	Example 1:
//		Input: mat = [[3,3,1,1],[2,2,1,2],[1,1,1,2]]
//		Output: [[1,1,1,1],[1,2,2,2],[1,2,3,3]]
//	Constraints:
//		m == mat.length
//		n == mat[i].length
//		1 <= m, n <= 100
//		1 <= mat[i][j] <= 100

func diagonalSort(mat [][]int) [][]int {
	diagonalMap := make(map[int][]int)
	m, n := len(mat), len(mat[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			diagonalMap[i-j] = append(diagonalMap[i-j], mat[i][j])
		}
	}
	for val, values := range diagonalMap {
		sort.Ints(values)
		var i, j int
		if val < 0 {
			i, j = 0, -val
		} else {
			i, j = val, 0
		}
		for _, value := range values {
			mat[i][j] = value
			i++
			j++
		}
	}
	return mat
}

func testProblem1329() {
	fmt.Println(diagonalSort([][]int{{3, 3, 1, 1}, {2, 2, 1, 2}, {1, 1, 1, 2}}))
}
