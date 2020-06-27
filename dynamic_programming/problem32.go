package dynamic_programming

import (
	"fmt"
	"github.com/leetcodeProblem/utils"
)

// Given a string containing just the characters '(' and ')', find the length of the longest valid (well-formed)
// parentheses substring.
//	Example 1:
//		Input: "(()"
//		Output: 2
//		Explanation: The longest valid parentheses substring is "()"
//	Example 2:
//		Input: ")()())"
//		Output: 4
//		Explanation: The longest valid parentheses substring is "()()"

func longestValidParentheses(s string) int {
	l := len(s)
	if l <= 1 {
		return 0
	}
	dp := make([]int, l)
	dp[0] = 0
	maxLen := 0
	for i := 1; i < l; i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				dp[i] = 2
				if i-2 >= 0 {
					dp[i] = dp[i] + dp[i-2]
				}
			} else if dp[i-1] > 0 {
				if i-dp[i-1]-1 >= 0 && s[i-dp[i-1]-1] == '(' {
					dp[i] = dp[i-1] + 2
					if i-dp[i-1]-2 >= 0 {
						dp[i] = dp[i] + dp[i-dp[i-1]-2]
					}
				}
			}
		}
		maxLen = utils.Max(maxLen, dp[i])
	}
	return maxLen
}

func testProblem32() {
	fmt.Println(longestValidParentheses("(()"))
	fmt.Println(longestValidParentheses(")()())"))
	fmt.Println(longestValidParentheses("((((((((())"))
}
