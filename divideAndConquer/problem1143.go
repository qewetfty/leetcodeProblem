package divideAndConquer

import "fmt"

// Given two strings text1 and text2, return the length of their longest common subsequence.
// A subsequence of a string is a new string generated from the original string with some characters(can be none)
// deleted without changing the relative order of the remaining characters. (eg, "ace" is a subsequence of "abcde"
// while "aec" is not). A common subsequence of two strings is a subsequence that is common to both strings.
// If there is no common subsequence, return 0.
//	Example 1:
//		Input: text1 = "abcde", text2 = "ace"
//		Output: 3
//		Explanation: The longest common subsequence is "ace" and its length is 3.
//	Example 2:
//		Input: text1 = "abc", text2 = "abc"
//		Output: 3
//		Explanation: The longest common subsequence is "abc" and its length is 3.
//	Example 3:
//		Input: text1 = "abc", text2 = "def"
//		Output: 0
//		Explanation: There is no such common subsequence, so the result is 0.
//	Constraints:
//		1 <= text1.length <= 1000
//		1 <= text2.length <= 1000
//		The input strings consist of lowercase English characters only.

func longestCommonSubsequence(text1 string, text2 string) int {
	if len(text1) == 0 || len(text2) == 0 {
		return 0
	}
	m, n := len(text1), len(text2)
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[m][n]
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {
	fmt.Println(longestCommonSubsequence("abcde", "ace"))
	fmt.Println(longestCommonSubsequence("abc", "abc"))
	fmt.Println(longestCommonSubsequence("abc", "def"))
	fmt.Println(longestCommonSubsequence("abc", "ace"))
}
