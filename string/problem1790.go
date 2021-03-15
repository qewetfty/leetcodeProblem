package string

import "fmt"

// You are given two strings s1 and s2 of equal length. A string swap is an
// operation where you choose two indices in a string (not necessarily different)
// and swap the characters at these indices.
// Return true if it is possible to make both strings equal by performing at most
// one string swap on exactly one of the strings. Otherwise, return false.
//	Example 1:
//		Input: s1 = "bank", s2 = "kanb"
//		Output: true
//		Explanation: For example, swap the first character with the last character of s2 to make "bank".
//	Example 2:
//		Input: s1 = "attack", s2 = "defend"
//		Output: false
//		Explanation: It is impossible to make them equal with one string swap.
//	Example 3:
//		Input: s1 = "kelb", s2 = "kelb"
//		Output: true
//		Explanation: The two strings are already equal, so no string swap operation is required.
//	Example 4:
//		Input: s1 = "abcd", s2 = "dcba"
//		Output: false
//	Constraints:
//		1 <= s1.length, s2.length <= 100
//		s1.length == s2.length
//		s1 and s2 consist of only lowercase English letters.

func areAlmostEqual(s1 string, s2 string) bool {
	l := len(s1)
	if l == 1 {
		return s1 == s2
	}
	swapPosition, swap := -1, false
	for i := 0; i < l; i++ {
		if s1[i] != s2[i] {
			if swapPosition < 0 {
				swapPosition = i
			} else {
				if !swap && s1[i] == s2[swapPosition] && s2[i] == s1[swapPosition] {
					swap = true
				} else {
					return false
				}
			}
		}
	}
	return (swapPosition < 0 && !swap) || (swapPosition >= 0 && swap)
}

func testProblem1790() {
	fmt.Println(areAlmostEqual("bank", "kanb"))
	fmt.Println(areAlmostEqual("attack", "defend"))
	fmt.Println(areAlmostEqual("kelb", "kelb"))
	fmt.Println(areAlmostEqual("abcd", "dcba"))
}
