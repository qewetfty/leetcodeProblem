package dynamic_programming

import "fmt"

// Given a string s, return the longest palindromic substring in s.
//	Example 1:
//		Input: s = "babad"
//		Output: "bab"
//		Note: "aba" is also a valid answer.
//	Example 2:
//		Input: s = "cbbd"
//		Output: "bb"
//	Example 3:
//		Input: s = "a"
//		Output: "a"
//	Example 4:
//		Input: s = "ac"
//		Output: "a"
//	Constraints:
//		1 <= s.length <= 1000
//		s consist of only digits and English letters (lower-case and/or upper-case),

// 中心扩展法，因为所有状态在转移的过程中是唯一的，因此可以通过中心状态扩展来快速处理结果
func longestPalindrome(s string) string {
	if s == "" {
		return s
	}
	start, end, l := 0, 0, len(s)
	for i := 0; i < l; i++ {
		left1, right1 := expand(s, i, i, l)
		left2, right2 := expand(s, i, i+1, l)
		if right1-left1 > end-start {
			start, end = left1, right1
		}
		if right2-left2 > end-start {
			start, end = left2, right2
		}
	}
	return s[start : end+1]
}

func expand(s string, left, right, l int) (int, int) {
	for ; left >= 0 && right < l && s[left] == s[right]; left, right = left-1, right+1 {

	}
	return left + 1, right - 1
}

// 动态规划解法，通过遍历字符串、记忆化搜索获取最长的回文串。时间复杂度O(n^2)
func longestPalindrome2(s string) string {
	l := len(s)
	if l <= 1 {
		return s
	}
	memory := make([][]int, l)
	for i := 0; i < l; i++ {
		memory[i] = make([]int, l)
		for j := 0; j < l; j++ {
			memory[i][j] = -1
		}
	}
	result := 0
	resultStr := ""
	for i := 0; i < l; i++ {
		for j := i; j < l; j++ {
			var curPalindromeLength int
			if memory[i][j] != -1 {
				curPalindromeLength = memory[i][j]
			} else {
				curPalindromeLength = findPalindrome(s, memory, i, j)
			}
			if curPalindromeLength > result {
				result = curPalindromeLength
				resultStr = s[i : j+1]
			}
		}
	}
	return resultStr
}

func findPalindrome(s string, memory [][]int, i, j int) int {
	if memory[i][j] != -1 {
		return memory[i][j]
	}
	if i > j {
		return 0
	}
	if i == j {
		memory[i][j] = 1
		return 1
	}
	if i+1 == j {
		if s[i] == s[j] {
			memory[i][j] = 2
			return 2
		} else {
			memory[i][j] = 0
			return 0
		}
	}
	if s[i] != s[j] {
		memory[i][j] = 0
		return 0
	} else {
		middle := findPalindrome(s, memory, i+1, j-1)
		if middle == 0 {
			memory[i][j] = 0
		} else {
			memory[i][j] = 2 + findPalindrome(s, memory, i+1, j-1)
		}
		return memory[i][j]
	}
}

func testProblem5() {
	fmt.Println(longestPalindrome("cbbdcbbd"))
	fmt.Println(longestPalindrome("babad"))
	fmt.Println(longestPalindrome("cbbd"))
	fmt.Println(longestPalindrome("a"))
	fmt.Println(longestPalindrome("ac"))
}
