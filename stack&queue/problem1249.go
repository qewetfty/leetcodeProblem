package stack_queue

import "fmt"

// Given a string s of '(' , ')' and lowercase English characters.
// Your task is to remove the minimum number of parentheses ( '(' or ')', in any
// positions ) so that the resulting parentheses string is valid and return any
// valid string.
// Formally, a parentheses string is valid if and only if:
//	It is the empty string, contains only lowercase characters, or
//	It can be written as AB (A concatenated with B), where A and B are valid strings, or
//	It can be written as (A), where A is a valid string.
//	Example 1:
//		Input: s = "lee(t(c)o)de)"
//		Output: "lee(t(c)o)de"
//		Explanation: "lee(t(co)de)" , "lee(t(c)ode)" would also be accepted.
//	Example 2:
//		Input: s = "a)b(c)d"
//		Output: "ab(c)d"
//	Example 3:
//		Input: s = "))(("
//		Output: ""
//		Explanation: An empty string is also valid.
//	Example 4:
//		Input: s = "(a(b(c)d)"
//		Output: "a(b(c)d)"
//	Constraints:
//		1 <= s.length <= 10^5
//		s[i] is one of  '(' , ')' and lowercase English letters.

func minRemoveToMakeValid(s string) string {
	removeIndex, stack := make(map[int]bool), make([]int, 0)
	l := len(s)
	for i := 0; i < l; i++ {
		char := s[i]
		if char == '(' {
			stack = append(stack, i)
		} else if char == ')' {
			if len(stack) == 0 {
				removeIndex[i] = true
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}
	for _, index := range stack {
		removeIndex[index] = true
	}
	resultByte := make([]byte, 0)
	for i := 0; i < l; i++ {
		if !removeIndex[i] {
			resultByte = append(resultByte, s[i])
		}
	}
	return string(resultByte)
}

func testProblem1249() {
	fmt.Println(minRemoveToMakeValid("lee(t(c)o)de)"))
	fmt.Println(minRemoveToMakeValid("a)b(c)d"))
	fmt.Println(minRemoveToMakeValid("(a(b(c)d)"))
	fmt.Println(minRemoveToMakeValid("))(("))
}
