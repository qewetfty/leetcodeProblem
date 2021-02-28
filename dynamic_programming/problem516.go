package dynamic_programming

import (
	"fmt"
	"github.com/leetcodeProblem/utils"
)

// Given a string s, find the longest palindromic subsequence's length in s.
// A subsequence is a sequence that can be derived from another sequence by
// deleting some or no elements without changing the order of the remaining
// elements.
//	Example 1:
//		Input: s = "bbbab"
//		Output: 4
//		Explanation: One possible longest palindromic subsequence is "bbbb".
//	Example 2:
//		Input: s = "cbbd"
//		Output: 2
//		Explanation: One possible longest palindromic subsequence is "bb".
//	Constraints:
//		1 <= s.length <= 1000
//		s consists only of lowercase English letters.

func longestPalindromeSubseq(s string) int {
	l := len(s)
	dp := make([][]int, l)
	for i := 0; i < l; i++ {
		dp[i] = make([]int, l)
	}
	for i := l - 1; i >= 0; i-- {
		dp[i][i] = 1
		for j := i + 1; j < l; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = utils.Max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return dp[0][l-1]
}

func testProblem516() {
	fmt.Println(longestPalindromeSubseq("bbbab"))
	fmt.Println(longestPalindromeSubseq("cbbd"))
}
