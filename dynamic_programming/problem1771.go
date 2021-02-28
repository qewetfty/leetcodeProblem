package dynamic_programming

import (
	"fmt"
	"github.com/leetcodeProblem/utils"
)

// You are given two strings, word1 and word2. You want to construct a string in the following manner:
//	Choose some non-empty subsequence subsequence1 from word1.
//	Choose some non-empty subsequence subsequence2 from word2.
//	Concatenate the subsequences: subsequence1 + subsequence2, to make the string.
// Return the length of the longest palindrome that can be constructed in the described manner. If no palindromes can be constructed, return 0.
// A subsequence of a string s is a string that can be made by deleting some
// (possibly none) characters from s without changing the order of the remaining
// characters.
//	A palindrome is a string that reads the same forward as well as backward.
//	Example 1:
//		Input: word1 = "cacb", word2 = "cbba"
//		Output: 5
//		Explanation: Choose "ab" from word1 and "cba" from word2 to make "abcba", which is a palindrome.
//	Example 2:
//		Input: word1 = "ab", word2 = "ab"
//		Output: 3
//		Explanation: Choose "ab" from word1 and "a" from word2 to make "aba", which is a palindrome.
//	Example 3:
//		Input: word1 = "aa", word2 = "bb"
//		Output: 0
//		Explanation: You cannot construct a palindrome from the described method, so return 0.
//	Constraints:
//		1 <= word1.length, word2.length <= 1000
//		word1 and word2 consist of lowercase English letters.

func longestPalindrome1771(word1 string, word2 string) int {
	word := word1 + word2
	l1, l2 := len(word1), len(word2)
	dp, l := make([][]int, l1+l2), l1+l2
	for i := 0; i < l; i++ {
		dp[i] = make([]int, l)
	}
	result := 0
	for i := l - 1; i >= 0; i-- {
		dp[i][i] = 1
		for j := i + 1; j < l; j++ {
			if word[i] == word[j] {
				dp[i][j] = dp[i+1][j-1] + 2
				if i < l1 && j >= l1 {
					result = utils.Max(dp[i][j], result)
				}
			} else {
				dp[i][j] = utils.Max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return result
}

func testProblem1771() {
	fmt.Println(longestPalindrome1771("aa", "bb"))
	fmt.Println(longestPalindrome1771("cacb", "cbba"))
	fmt.Println(longestPalindrome1771("ab", "ab"))
}
