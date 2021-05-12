package dynamic_programming

// Given a 2D matrix matrix, handle multiple queries of the following type:
// Calculate the sum of the elements of matrix inside the rectangle defined by
// its upper left corner (row1, col1) and lower right corner (row2, col2).
// Implement the NumMatrix class:
//	NumMatrix(int[][] matrix) Initializes the object with the integer matrix matrix.
// 	int sumRegion(int row1, int col1, int row2, int col2) Returns the sum of the
// 	elements of matrix inside the rectangle defined by its upper left corner
// 	(row1, col1) and lower right corner (row2, col2).
//	Example 1:
//		Input
//			["NumMatrix", "sumRegion", "sumRegion", "sumRegion"]
//			[[[[3, 0, 1, 4, 2], [5, 6, 3, 2, 1], [1, 2, 0, 1, 5], [4, 1, 0, 1, 7], [1, 0, 3, 0, 5]]], [2, 1, 4, 3], [1, 1, 2, 2], [1, 2, 2, 4]]
//		Output
//			[null, 8, 11, 12]
//		Explanation
//			NumMatrix numMatrix = new NumMatrix([[3, 0, 1, 4, 2], [5, 6, 3, 2, 1], [1, 2, 0, 1, 5], [4, 1, 0, 1, 7], [1, 0, 3, 0, 5]]);
//			numMatrix.sumRegion(2, 1, 4, 3); // return 8 (i.e sum of the red rectangle)
//			numMatrix.sumRegion(1, 1, 2, 2); // return 11 (i.e sum of the green rectangle)
//			numMatrix.sumRegion(1, 2, 2, 4); // return 12 (i.e sum of the blue rectangle)
//	Constraints:
//		m == matrix.length
//		n == matrix[i].length
//		1 <= m, n <= 200
//		-105 <= matrix[i][j] <= 105
//		0 <= row1 <= row2 < m
//		0 <= col1 <= col2 < n
//		At most 104 calls will be made to sumRegion.

type NumMatrix struct {
	sumMatrix [][]int
}

func Constructor304(matrix [][]int) NumMatrix {
	m, n := len(matrix), len(matrix[0])
	preMatrix := make([][]int, m)
	for i := 0; i < m; i++ {
		preMatrix[i] = make([]int, n)
	}
	preMatrix[0][0] = matrix[0][0]
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j > 0 {
				preMatrix[i][j] = preMatrix[i][j-1] + matrix[i][j]
			} else if j == 0 && i > 0 {
				preMatrix[i][j] = preMatrix[i-1][j] + matrix[i][j]
			} else if i > 0 && j > 0 {
				preMatrix[i][j] = preMatrix[i-1][j] + preMatrix[i][j-1] - preMatrix[i-1][j-1] + matrix[i][j]
			}
		}
	}
	return NumMatrix{
		sumMatrix: preMatrix,
	}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	result := this.sumMatrix[row2][col2]
	if row1 == 0 && col1 == 0 {
		return result
	}
	if row1 > 0 && col1 > 0 {
		result = result - this.sumMatrix[row1-1][col2] - this.sumMatrix[row2][col1-1] + this.sumMatrix[row1-1][col1-1]
	} else if row1 == 0 && col1 > 0 {
		result = result - this.sumMatrix[row2][col1-1]
	} else if row1 > 0 && col1 == 0 {
		result = result - this.sumMatrix[row1-1][col2]
	}
	return result
}
