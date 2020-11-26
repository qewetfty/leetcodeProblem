package divideAndConquer

import (
	"fmt"
	"github.com/leetcodeProblem/utils"
	"strings"
)

// Given a string s and an integer k, return the length of the longest substring
// of s such that the frequency of each character in this substring is less than
// or equal to k.
//	Example 1:
//		Input: s = "aaabb", k = 3
//		Output: 3
//		Explanation: The longest substring is "aaa", as 'a' is repeated 3 times.
//	Example 2:
//		Input: s = "ababbc", k = 2
//		Output: 5
//		Explanation: The longest substring is "ababb", as 'a' is repeated 2 times and 'b' is repeated 3 times.
//	Constraints:
//		1 <= s.length <= 104
//		s consists of only lowercase English letters.
//		1 <= k <= 105

// 滑动窗口解法
func longestSubstring(s string, k int) int {
	l := len(s)
	if k > l {
		return 0
	}
	charMap := make(map[byte]int, 0)
	for i := 0; i < l; i++ {
		charMap[s[i]]++
	}
	uniqueCharCount := len(charMap)
	res := 0
	for i := 1; i <= uniqueCharCount; i++ {
		curCharMap := make(map[byte]int)
		windowStart, windowEnd, differentCount, count := 0, 0, 0, 0
		for windowEnd < l {
			// 向右扩大窗口
			curChar := s[windowEnd]
			curCharMap[curChar]++
			if curCharMap[curChar] == 1 {
				differentCount++
			}
			if curCharMap[curChar] == k {
				count++
			}
			windowEnd++
			// 窗口中元素最多的个数为i
			for windowStart < windowEnd && differentCount > i {
				leftChar := s[windowStart]
				if curCharMap[leftChar] == 1 {
					differentCount--
				}
				if curCharMap[leftChar] == k {
					count--
				}
				curCharMap[leftChar]--
				windowStart++
			}
			if differentCount == i && differentCount == count {
				res = utils.Max(res, windowEnd-windowStart)
			}
		}
	}
	return res
}

// 递归解法（分治）
func longestSubstring2(s string, k int) int {
	l := len(s)
	if k > l {
		return 0
	}
	// 记录每个字符的个数，接下来判断时，如果字符的个数小于K，能够直接快速判断
	charCount := make(map[byte]int)
	for i := 0; i < l; i++ {
		charCount[s[i]]++
	}
	// 处理不符合条件的字符
	for i := 0; i < l; i++ {
		if charCount[s[i]] < k {
			s = s[:i] + "," + s[i+1:]
		}
	}
	// 根据不符合条件的字符串分割，每一组重新进行判断
	split := strings.Split(s, ",")
	// 如果仅剩一组了，说明这一组符合条件，直接返回第0个元素长度
	if len(split) == 1 {
		return len(split[0])
	}
	max := 0
	// 递归处理每一段
	for _, str := range split {
		max = utils.Max(max, longestSubstring(str, k))
	}
	return max
}

func testProblem395() {
	fmt.Println(longestSubstring("aaabb", 3))
	fmt.Println(longestSubstring("ababbc", 2))
	fmt.Println(longestSubstring("bbaaacbd", 3))
}
