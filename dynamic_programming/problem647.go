package dynamic_programming

import "fmt"

// Given a string, your task is to count how many palindromic substrings in this string.
// The substrings with different start indexes or end indexes are counted as
// different substrings even they consist of same characters.
//	Example 1:
//		Input: "abc"
//		Output: 3
//		Explanation: Three palindromic strings: "a", "b", "c".
//	Example 2:
//		Input: "aaa"
//		Output: 6
//		Explanation: Six palindromic strings: "a", "a", "a", "aa", "aa", "aaa".
//	Note:
//		The input string length won't exceed 1000.

func countSubstrings(s string) int {
	l := len(s)
	if l < 2 {
		return l
	}
	dp := make([][]bool, l)
	for i := 0; i < l; i++ {
		dp[i] = make([]bool, l)
	}
	result := 0
	for i := l - 1; i >= 0; i-- {
		dp[i][i], result = true, result+1
		if i+1 < l && s[i] == s[i+1] {
			dp[i][i+1], result = true, result+1
		}
		for j := i + 2; j < l; j++ {
			if s[i] == s[j] && dp[i+1][j-1] {
				dp[i][j], result = true, result+1
			}
		}
	}
	return result
}

func testProblem647() {
	fmt.Println(countSubstrings("abc"))
	fmt.Println(countSubstrings("aaa"))
}
