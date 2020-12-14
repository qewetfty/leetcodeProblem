package dynamic_programming

import "fmt"

// Given a string s, partition s such that every substring of the partition is a
// palindrome. Return all possible palindrome partitioning of s.
// A palindrome string is a string that reads the same backward as forward.
//	Example 1:
//		Input: s = "aab"
//		Output: [["a","a","b"],["aa","b"]]
//	Example 2:
//		Input: s = "a"
//		Output: [["a"]]
//	Constraints:
//		1 <= s.length <= 16
//		s contains only lowercase English letters.

var result [][]string
var l int

// dp解法，主要使用dp去判断一个字符串是否是回文。
func partition(s string) [][]string {
	result, l = make([][]string, 0), len(s)
	curList := make([]string, 0)
	palindromeList := make([][]bool, l)
	for i := 0; i < len(s); i++ {
		palindromeList[i] = make([]bool, l)
	}
	dfs131_dp(s, 0, curList, palindromeList)
	return result
}

func dfs131_dp(s string, start int, curList []string, palindromeList [][]bool) {
	if start >= l {
		newList := make([]string, len(curList))
		copy(newList, curList)
		result = append(result, newList)
	}
	for end := start; end < l; end++ {
		if s[start] == s[end] && (end-start <= 2 || palindromeList[start+1][end-1]) {
			palindromeList[start][end] = true
			curList = append(curList, s[start:end+1])
			dfs131_dp(s, end+1, curList, palindromeList)
			curList = curList[:len(curList)-1]
		}
	}
}

// dfs遍历方法，判断每一种情况
func partition2(s string) [][]string {
	result = make([][]string, 0)
	curList := make([]string, 0)
	dfs131(s, 0, curList)
	return result
}

func dfs131(s string, start int, curList []string) {
	if start >= len(s) {
		newList := make([]string, len(curList))
		copy(newList, curList)
		result = append(result, newList)
	}
	for end := start; end < len(s); end++ {
		if isPalindrome(s, start, end) {
			curList = append(curList, s[start:end+1])
			dfs131(s, end+1, curList)
			curList = curList[:len(curList)-1]
		}
	}
}

func isPalindrome(s string, start, end int) bool {
	for start < end {
		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}
	return true
}

func testProblem131() {
	fmt.Println(partition("aab"))
	fmt.Println(partition("a"))
	fmt.Println(partition("abba"))
}
