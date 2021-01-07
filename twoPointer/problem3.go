package twoPointer

import (
	"fmt"
	"github.com/leetcodeProblem/utils"
)

// Given a string s, find the length of the longest substring without repeating characters.
//	Example 1:
//		Input: s = "abcabcbb"
//		Output: 3
//		Explanation: The answer is "abc", with the length of 3.
//	Example 2:
//		Input: s = "bbbbb"
//		Output: 1
//		Explanation: The answer is "b", with the length of 1.
//	Example 3:
//		Input: s = "pwwkew"
//		Output: 3
//		Explanation: The answer is "wke", with the length of 3.
//		Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.
//	Example 4:
//		Input: s = ""
//		Output: 0
//	Constraints:
//		0 <= s.length <= 5 * 104
//		s consists of English letters, digits, symbols and spaces.

// sliding window和双指针解法
func lengthOfLongestSubstring(s string) int {
	l := len(s)
	r, charMap := 0, make(map[byte]int)
	res := 0
	for i := 0; i < l; i++ {
		// 移动左边指针
		if i != 0 {
			delete(charMap, s[i-1])
		}
		// 移动右边指针
		for r < l && charMap[s[r]] == 0 {
			charMap[s[r]]++
			r++
		}
		res = utils.Max(res, r-i)
	}
	return res
}

// sliding window做法，使用map存储最后一个字符出现的位置
func lengthOfLongestSubstring2(s string) int {
	l := len(s)
	lastIndex, res := -1, 0
	charMap := make(map[byte]int)
	for i := 0; i < l; i++ {
		curChar := s[i]
		if _, ok := charMap[curChar]; ok {
			lastIndex = charMap[curChar]
			for k, v := range charMap {
				if v < lastIndex {
					delete(charMap, k)
				}
			}
		}
		charMap[curChar] = i
		if i-lastIndex > res {
			res = i - lastIndex
		}
	}
	return res
}

func testProblem3() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
	fmt.Println(lengthOfLongestSubstring(""))
	fmt.Println(lengthOfLongestSubstring("s"))
}
