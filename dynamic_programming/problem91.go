package dynamic_programming

import "fmt"

// A message containing letters from A-Z is being encoded to numbers using the following mapping:
//	'A' -> 1
//	'B' -> 2
//	...
//	'Z' -> 26
// Given a non-empty string containing only digits, determine the total number of ways to decode it.
//	Example 1:
//		Input: "12"
//		Output: 2
//		Explanation: It could be decoded as "AB" (1 2) or "L" (12).
//	Example 2:
//		Input: "226"
//		Output: 3
//		Explanation: It could be decoded as "BZ" (2 26), "VF" (22 6), or "BBF" (2 2 6).

func numDecodings(s string) int {
	l := len(s)
	if l == 0 {
		return l
	}
	if s[0] == '0' {
		return 0
	}
	dp := make([]int, l+1)
	dp[0], dp[1] = 1, 1
	for i := 1; i < l; i++ {
		if s[i] != '0' {
			dp[i+1] = dp[i]
		}
		lastTwo := 10*(s[i-1]-'0') + (s[i] - '0')
		if 10 <= lastTwo && lastTwo <= 26 {
			dp[i+1] = dp[i+1] + dp[i-1]
		}
	}
	return dp[l]
}

func testProblem91() {
	fmt.Println(numDecodings("17"))
	fmt.Println(numDecodings("101"))
	fmt.Println(numDecodings("0"))
	fmt.Println(numDecodings("4"))
	fmt.Println(numDecodings("07123124"))
	fmt.Println(numDecodings("226"))
	fmt.Println(numDecodings("12"))
	fmt.Println(numDecodings("29847612684712970312"))
}
