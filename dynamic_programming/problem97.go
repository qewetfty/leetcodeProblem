package dynamic_programming

import "fmt"

// Given strings s1, s2, and s3, find whether s3 is formed by an interleaving of s1 and s2.
// An interleaving of two strings s and t is a configuration where they are divided into non-empty substrings such that:
//	s = s1 + s2 + ... + sn
//	t = t1 + t2 + ... + tm
//	|n - m| <= 1
//	The interleaving is s1 + t1 + s2 + t2 + s3 + t3 + ... or t1 + s1 + t2 + s2 + t3 + s3 + ...
//	Note: a + b is the concatenation of strings a and b.
//	Example 1:
//		Input: s1 = "aabcc", s2 = "dbbca", s3 = "aadbbcbcac"
//		Output: true
//	Example 2:
//		Input: s1 = "aabcc", s2 = "dbbca", s3 = "aadbbbaccc"
//		Output: false
//	Example 3:
//		Input: s1 = "", s2 = "", s3 = ""
//		Output: true
//	Constraints:
//		0 <= s1.length, s2.length <= 100
//		0 <= s3.length <= 200
//		s1, s2, and s3 consist of lowercase English letters.
//	Follow up: Could you solve it using only O(s2.length) additional memory space?

func isInterleave(s1 string, s2 string, s3 string) bool {
	l1, l2, l3 := len(s1), len(s2), len(s3)
	if l1+l2 != l3 {
		return false
	}
	dp := make([]bool, l2+1)
	dp[0] = true
	for i := 0; i <= l1; i++ {
		for j := 0; j <= l2; j++ {
			p := i + j - 1
			if i > 0 {
				dp[j] = dp[j] && s1[i-1] == s3[p]
			}
			if j > 0 {
				dp[j] = dp[j] || dp[j-1] && s2[j-1] == s3[p]
			}
		}
	}
	return dp[l1]
}

func testProblem97() {
	fmt.Println(isInterleave("aabcc", "dbbca", "aadbbbaccc"))
	fmt.Println(isInterleave("", "", ""))
}
