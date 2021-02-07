package string

import (
	"fmt"
	"github.com/leetcodeProblem/utils"
	"math"
)

// Given a string s and a character c that occurs in s, return an array of
// integers answer where answer.length == s.length and answer[i] is the shortest
// distance from s[i] to the character c in s.
//	Example 1:
//		Input: s = "loveleetcode", c = "e"
//		Output: [3,2,1,0,1,0,0,1,2,2,1,0]
//	Example 2:
//		Input: s = "aaab", c = "b"
//		Output: [3,2,1,0]
//	Constraints:
//		1 <= s.length <= 104
//		s[i] and c are lowercase English letters.
//		c occurs at least once in s.

func shortestToChar(s string, c byte) []int {
	charPos, l := make([]int, 0), len(s)
	for i := 0; i < l; i++ {
		if s[i] == c {
			charPos = append(charPos, i)
		}
	}
	result := make([]int, 0)
	leftPosition, rightPosition := math.MinInt32, charPos[0]
	curPos := 0
	for i := 0; i < l; i++ {
		result = append(result, utils.Min(i-leftPosition, rightPosition-i))
		if i == rightPosition {
			leftPosition = i
			curPos++
			if curPos == len(charPos) {
				rightPosition = math.MaxInt32
			} else {
				rightPosition = charPos[curPos]
			}
		}
	}
	return result
}

func testProblem821() {
	fmt.Println(shortestToChar("loveleetcode", 'e'))
	fmt.Println(shortestToChar("aaab", 'b'))
}
