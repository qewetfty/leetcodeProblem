package string

import (
	"fmt"
	"strings"
)

// Given a non-empty string check if it can be constructed by taking a substring of it and appending multiple copies of
// the substring together. You may assume the given string consists of lowercase English letters only and its length
// will not exceed 10000.
//	Example 1:
//		Input: "abab"
//		Output: True
//		Explanation: It's the substring "ab" twice.
//	Example 2:
//		Input: "aba"
//		Output: False
//	Example 3:
//		Input: "abcabcabcabc"
//		Output: True
//		Explanation: It's the substring "abc" four times. (And the substring "abcabc" twice.)

// 暴力法
func repeatedSubstringPatternFource(s string) bool {
	l := len(s)
	if l <= 1 {
		return false
	}
	for i := 1; i <= l/2; i++ {
		if l%i != 0 {
			continue
		}
		sub := s[:i]
		if strings.Count(s, sub) == l/i {
			return true
		}
	}
	return false
}

// 字符串匹配算法
func repeatedSubstringPattern(s string) bool {
	return strings.Contains((s + s)[1:len(s)*2-1], s)
}

func testProblem459() {
	fmt.Println(repeatedSubstringPattern("abab"))
	fmt.Println(repeatedSubstringPattern("aba"))
	fmt.Println(repeatedSubstringPattern("abcabcabcabc"))
	fmt.Println(repeatedSubstringPattern("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab"))
}
