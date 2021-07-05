package dynamic_programming

import "fmt"

// Given an integer n, your task is to count how many strings of length n can be formed under the following rules:
//	Each character is a lower case vowel ('a', 'e', 'i', 'o', 'u')
//	Each vowel 'a' may only be followed by an 'e'.
//	Each vowel 'e' may only be followed by an 'a' or an 'i'.
//	Each vowel 'i' may not be followed by another 'i'.
//	Each vowel 'o' may only be followed by an 'i' or a 'u'.
//	Each vowel 'u' may only be followed by an 'a'.
// Since the answer may be too large, return it modulo 10^9 + 7.
//	Example 1:
//		Input: n = 1
//		Output: 5
//		Explanation: All possible strings are: "a", "e", "i" , "o" and "u".
//	Example 2:
//		Input: n = 2
//		Output: 10
//		Explanation: All possible strings are: "ae", "ea", "ei", "ia", "ie", "io", "iu", "oi", "ou" and "ua".
//	Example 3:
//		Input: n = 5
//		Output: 68
//	Constraints:
//		1 <= n <= 2 * 10^4

func countVowelPermutation(n int) int {
	if n == 1 {
		return 5
	}
	dp := make([][]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = make([]int, 5)
	}
	dp[1][0], dp[1][1], dp[1][2], dp[1][3], dp[1][4] = 1, 1, 1, 1, 1
	for i := 2; i <= n; i++ {
		dp[i][0] = (dp[i-1][1] + dp[i-1][2] + dp[i-1][4]) % 1000000007
		dp[i][1] = (dp[i-1][0] + dp[i-1][2]) % 1000000007
		dp[i][2] = (dp[i-1][1] + dp[i-1][3]) % 1000000007
		dp[i][3] = dp[i-1][2]
		dp[i][4] = (dp[i-1][2] + dp[i-1][3]) % 1000000007
	}
	return (dp[n][0] + dp[n][1] + dp[n][2] + dp[n][3] + dp[n][4]) % 1000000007
}

func testProblem1220() {
	fmt.Println(countVowelPermutation(1))
	fmt.Println(countVowelPermutation(2))
	fmt.Println(countVowelPermutation(3))
	fmt.Println(countVowelPermutation(4))
	fmt.Println(countVowelPermutation(5))
}
