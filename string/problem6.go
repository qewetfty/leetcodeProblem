package string

import "fmt"

// The string "PAYPALISHIRING" is written in a zigzag pattern on a given number
// of rows like this: (you may want to display this pattern in a fixed font for
// better legibility)
//		P   A   H   N
//		A P L S I I G
//		Y   I   R
// And then read line by line: "PAHNAPLSIIGYIR"
// Write the code that will take a string and make this conversion given a number of rows:
//		string convert(string s, int numRows);
//	Example 1:
//		Input: s = "PAYPALISHIRING", numRows = 3
//		Output: "PAHNAPLSIIGYIR"
//	Example 2:
//		Input: s = "PAYPALISHIRING", numRows = 4
//		Output: "PINALSIGYAHRPI"
//		Explanation:
//			P     I    N
//			A   L S  I G
//			Y A   H R
//			P     I
//	Example 3:
//		Input: s = "A", numRows = 1
//		Output: "A"
//	Constraints:
//		1 <= s.length <= 1000
//		s consists of English letters (lower-case and upper-case), ',' and '.'.
//		1 <= numRows <= 1000

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	bytesList := make([][]byte, numRows)
	for i := 0; i < numRows; i++ {
		bytesList[i] = make([]byte, 0)
	}
	startRow, l, down := 0, len(s), true
	for i := 0; i < l; i++ {
		char := s[i]
		bytesList[startRow] = append(bytesList[startRow], char)
		if down {
			startRow++
			if startRow == numRows {
				startRow = startRow - 2
				down = false
			}
		} else {
			startRow--
			if startRow < 0 {
				startRow = startRow + 2
				down = true
			}
		}
	}
	resultByte := make([]byte, 0)
	for _, bytes := range bytesList {
		resultByte = append(resultByte, bytes...)
	}
	return string(resultByte)
}

func testProblem6() {
	fmt.Println(convert("PAYPALISHIRING", 3))
	fmt.Println(convert("PAYPALISHIRING", 4))
	fmt.Println(convert("A", 1))
}
