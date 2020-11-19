package stack_queue

import (
	"fmt"
	"strconv"
	"strings"
)

// Given an encoded string, return its decoded string.
// The encoding rule is: k[encoded_string], where the encoded_string inside the
// square brackets is being repeated exactly k times. Note that k is guaranteed
// to be a positive integer.
// You may assume that the input string is always valid; No extra white spaces, square brackets are well-formed, etc.
// Furthermore, you may assume that the original data does not contain any digits
// and that digits are only for those repeat numbers, k. For example, there won't
// be input like 3a or 2[4].
//	Example 1:
//		Input: s = "3[a]2[bc]"
//		Output: "aaabcbc"
//	Example 2:
//		ZInput: s = "3[a2[c]]"
//		Output: "accaccacc"
//	Example 3:
//		Input: s = "2[abc]3[cd]ef"
//		Output: "abcabccdcdcdef"
//	Example 4:
//		Input: s = "abc3[cd]xyz"
//		Output: "abccdcdcdxyz"
//	Constraints:
//		1 <= s.length <= 30
//		s consists of lowercase English letters, digits, and square brackets '[]'.
//		s is guaranteed to be a valid input.
//		All the integers in s are in the range [1, 300].

func decodeString(s string) string {
	stack, l := make([]string, 0), len(s)
	for i := 0; i < l; i++ {
		curStr := s[i : i+1]
		switch curStr {
		// 如果遇到括号结束，就开始从栈里弹出，直到遇到'['
		case "]":
			repeatStr := ""
			for stack[len(stack)-1] != "[" {
				repeatStr = stack[len(stack)-1] + repeatStr
				stack = stack[:len(stack)-1]
			}
			// 弹出'['并且获取'['前的数字
			repeatTimeStr, cur := "", 2
			for cur <= len(stack) && stack[len(stack)-cur] >= "0" && stack[len(stack)-cur] <= "9" {
				repeatTimeStr = stack[len(stack)-cur] + repeatTimeStr
				cur++
			}
			repeatTime, _ := strconv.Atoi(repeatTimeStr)
			stack = stack[:len(stack)-cur+1]
			appendStr := strings.Repeat(repeatStr, repeatTime)
			// 最后的结果也要入栈，可能外层还有括号匹配
			stack = append(stack, appendStr)
		default:
			stack = append(stack, curStr)
		}
	}
	return strings.Join(stack, "")
}

func testProblem394() {
	fmt.Println(decodeString("100[leetcode]"))
	fmt.Println(decodeString("3[a]2[bc]"))
	fmt.Println(decodeString("3[a2[c]]"))
	fmt.Println(decodeString("2[abc]3[cd]ef"))
	fmt.Println(decodeString("abc3[cd]xyz"))
}
