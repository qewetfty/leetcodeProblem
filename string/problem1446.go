package string

import (
	"fmt"
	"github.com/leetcodeProblem/utils"
)

// Given a string s, the power of the string is the maximum length of a non-empty
// substring that contains only one unique character.
// Return the power of the string.
//	Example 1:
//		Input: s = "leetcode"
//		Output: 2
//		Explanation: The substring "ee" is of length 2 with the character 'e' only.
//	Example 2:
//		Input: s = "abbcccddddeeeeedcba"
//		Output: 5
//		Explanation: The substring "eeeee" is of length 5 with the character 'e' only.
//	Example 3:
//		Input: s = "triplepillooooow"
//		Output: 5
//	Example 4:
//		Input: s = "hooraaaaaaaaaaay"
//		Output: 11
//	Example 5:
//		Input: s = "tourist"
//		Output: 1
//	Constraints:
//		1 <= s.length <= 500
//		s contains only lowercase English letters.

func maxPower(s string) int {
	l := len(s)
	if l < 2 {
		return l
	}
	curChar, length, res := s[0], 1, 1
	for i := 1; i < l; i++ {
		if s[i] == curChar {
			length++
			res = utils.Max(res, length)
		} else {
			res = utils.Max(res, length)
			curChar, length = s[i], 1
		}
	}
	return res
}

func testProblem1446() {
	fmt.Println(maxPower("leetcode"))
	fmt.Println(maxPower("abbcccddddeeeeedcba"))
	fmt.Println(maxPower("triplepillooooow"))
	fmt.Println(maxPower("hooraaaaaaaaaaay"))
	fmt.Println(maxPower("tourist"))
}
