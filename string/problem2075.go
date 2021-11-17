package string

import (
	"fmt"
	"strings"
)

// A string originalText is encoded using a slanted transposition cipher to a
// string encodedText with the help of a matrix having a fixed number of rows
// rows.
// originalText is placed first in a top-left to bottom-right manner.
// The blue cells are filled first, followed by the red cells, then the yellow
// cells, and so on, until we reach the end of originalText. The arrow indicates
// the order in which the cells are filled. All empty cells are filled with ' '.
// The number of columns is chosen such that the rightmost column will not be
// empty after filling in originalText.
// encodedText is then formed by appending all characters of the matrix in a row-wise fashion.
// The characters in the blue cells are appended first to encodedText, then the
// red cells, and so on, and finally the yellow cells. The arrow indicates the
// order in which the cells are accessed.
// For example, if originalText = "cipher" and rows = 3, then we encode it in the following manner:
// The blue arrows depict how originalText is placed in the matrix, and the red
// arrows denote the order in which encodedText is formed. In the above example,
// encodedText = "ch ie pr".
// Given the encoded string encodedText and number of rows rows, return the original string originalText.
//	Note: originalText does not have any trailing spaces ' '. The test cases are generated such that there is only one possible originalText.
//	Example 1:
//		Input: encodedText = "ch   ie   pr", rows = 3
//		Output: "cipher"
//		Explanation: This is the same example described in the problem description.
//	Example 2:
//		Input: encodedText = "iveo    eed   l te   olc", rows = 4
//		Output: "i love leetcode"
//		Explanation: The figure above denotes the matrix that was used to encode originalText.
//					 The blue arrows show how we can find originalText from encodedText.
//	Example 3:
//		Input: encodedText = "coding", rows = 1
//		Output: "coding"
//		Explanation: Since there is only 1 row, both originalText and encodedText are the same.
//	Example 4:
//		Input: encodedText = " b  ac", rows = 2
//		Output: " abc"
//		Explanation: originalText cannot have trailing spaces, but it may be preceded by one or more spaces.
//	Constraints:
//		0 <= encodedText.length <= 106
//		encodedText consists of lowercase English letters and ' ' only.
//		encodedText is a valid encoding of some originalText that does not have trailing spaces.
//		1 <= rows <= 1000
//		The testcases are generated such that there is only one possible originalText.

func decodeCiphertext(encodedText string, rows int) string {
	l := len(encodedText)
	col := l / rows
	matrix := make([][]byte, rows)
	for i := 0; i < rows; i++ {
		matrix[i] = make([]byte, col)
	}
	// 还原矩阵
	for i := 0; i < l; i++ {
		curRow, curCol := i/col, i%col
		matrix[curRow][curCol] = encodedText[i]
	}
	// 通过矩阵获取原字符串
	resultByte := make([]byte, 0)
	for i := 0; i < col; i++ {
		curByte := make([]byte, 0)
		for j := 0; j < rows; j++ {
			if i+j >= col {
				break
			}
			curByte = append(curByte, matrix[j][i+j])
		}
		curStr := string(curByte)
		if i+rows >= col && strings.EqualFold("", strings.TrimLeft(curStr, " ")) {
			break
		}
		resultByte = append(resultByte, curByte...)
	}
	return strings.TrimRight(string(resultByte), " ")
}

func testProblem2075() {
	fmt.Println(decodeCiphertext("ch   ie   pr", 3))
	fmt.Println(decodeCiphertext("iveo    eed   l te   olc", 4))
	fmt.Println(decodeCiphertext(" b  ac", 2))
	fmt.Println(decodeCiphertext("", 10))
}
