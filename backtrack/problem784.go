package backtrack

import "fmt"

// Given a string S, we can transform every letter individually to be lowercase or uppercase to create another string.
// Return a list of all possible strings we could create. You can return the output in any order.
//	Example 1:
//		Input: S = "a1b2"
//		Output: ["a1b2","a1B2","A1b2","A1B2"]
//	Example 2:
//		Input: S = "3z4"
//		Output: ["3z4","3Z4"]
//	Example 3:
//		Input: S = "12345"
//		Output: ["12345"]
//	Example 4:
//		Input: S = "0"
//		Output: ["0"]
//	Constraints:
//		S will be a string with length between 1 and 12.
//		S will consist only of letters or digits.

func letterCasePermutation(S string) []string {
	result := make([]string, 0)
	backtrack784(S, "", &result)
	return result
}

func backtrack784(s, curStr string, result *[]string) {
	strLen, curLen := len(s), len(curStr)
	if curLen == strLen {
		*result = append(*result, curStr)
		return
	}
	char := s[curLen]
	curStr = curStr + string(char)
	backtrack784(s, curStr, result)
	if 'a' <= char && char <= 'z' {
		// 增加大写的
		curStr = curStr[:curLen]
		curStr = curStr + string(char-32)
		backtrack784(s, curStr, result)
	} else if 'A' <= char && char <= 'Z' {
		// 增加小写的
		curStr = curStr[:curLen]
		curStr = curStr + string(char+32)
		backtrack784(s, curStr, result)
	}
}

func testProblem784() {
	fmt.Println(letterCasePermutation("a1b2"))
	fmt.Println(letterCasePermutation("3z4"))
	fmt.Println(letterCasePermutation("12345"))
	fmt.Println(letterCasePermutation("0"))
}
