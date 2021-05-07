package dynamic_programming

import (
	"fmt"
	"github.com/leetcodeProblem/utils"
)

// Given two strings word1 and word2, return the minimum number of steps required to make word1 and word2 the same.
// In one step, you can delete exactly one character in either string.
//	Example 1:
//		Input: word1 = "sea", word2 = "eat"
//		Output: 2
//		Explanation: You need one step to make "sea" to "ea" and another step to make "eat" to "ea".
//	Example 2:
//		Input: word1 = "leetcode", word2 = "etco"
//		Output: 4
//	Constraints:
//		1 <= word1.length, word2.length <= 500
//		word1 and word2 consist of only lowercase English letters.

// 最长公共字串解法
func minDistance583_2(word1 string, word2 string) int {
	l1, l2 := len(word1), len(word2)
	dp := make([][]int, l1+1)
	for i := 0; i <= l1; i++ {
		dp[i] = make([]int, l2+1)
	}
	for i := 0; i <= l1; i++ {
		for j := 0; j <= l2; j++ {
			if i == 0 || j == 0 {
				continue
			}
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = utils.Max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return l1 + l2 - 2*dp[l1][l2]
}

// 记忆化搜索方法
var memo [][]int

func minDistance583(word1 string, word2 string) int {
	l1, l2 := len(word1), len(word2)
	memo = make([][]int, l1+1)
	for i := 0; i <= l1; i++ {
		memo[i] = make([]int, l2+1)
		for j := 0; j <= l2; j++ {
			memo[i][j] = -1
		}
	}
	return getDistance(word1, word2, 0, 0)
}

func getDistance(word1, word2 string, i, j int) int {
	if memo[i][j] >= 0 {
		return memo[i][j]
	}
	l1, l2 := len(word1), len(word2)
	if i == l1 {
		memo[i][j] = l2 - j
		return memo[i][j]
	}
	if l2 == j {
		memo[i][j] = l1 - i
		return memo[i][j]
	}
	minDiff := utils.Min(getDistance(word1, word2, i+1, j), getDistance(word1, word2, i, j+1)) + 1
	if word1[i] != word2[j] {
		memo[i][j] = minDiff
		return memo[i][j]
	}
	memo[i][j] = utils.Min(minDiff, getDistance(word1, word2, i+1, j+1))
	return memo[i][j]
}

func testProblem583() {
	fmt.Println(minDistance583("sea", "eat"))
	fmt.Println(minDistance583("leetcode", "etco"))
}
