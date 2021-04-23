package string

import (
	"fmt"
	"github.com/leetcodeProblem/utils"
)

// Give a string s, count the number of non-empty (contiguous) substrings that
// have the same number of 0's and 1's, and all the 0's and all the 1's in these
// substrings are grouped consecutively.
// Substrings that occur multiple times are counted the number of times they occur.
//	Example 1:
//		Input: "00110011"
//		Output: 6
//		Explanation: There are 6 substrings that have equal number of consecutive 1's and 0's: "0011", "01", "1100", "10", "0011", and "01".
//					 Notice that some of these substrings repeat and are counted the number of times they occur.
//					 Also, "00110011" is not a valid substring because all the 0's (and 1's) are not grouped together.
//	Example 2:
//		Input: "10101"
//		Output: 4
//		Explanation: There are 4 substrings: "10", "01", "10", "01" that have equal number of consecutive 1's and 0's.
//	Note:
//		s.length will be between 1 and 50,000.
//		s will only consist of "0" or "1" characters.

func countBinarySubstrings(s string) int {
	l := len(s)
	init := true
	countList := make([]int, 2)
	lastChar := s[0]
	countList[lastChar-'0']++
	result := 0
	for i := 1; i < l; i++ {
		curChar := s[i]
		if curChar != lastChar {
			if init {
				init = false
			} else {
				result = result + utils.Min(countList[0], countList[1])
				countList[curChar-'0'] = 0
			}
			lastChar = curChar
		}
		countList[curChar-'0']++
	}
	result = result + utils.Min(countList[0], countList[1])
	return result
}

func testProblem696() {
	fmt.Println(countBinarySubstrings("00110011"))
	fmt.Println(countBinarySubstrings("10101"))
}
