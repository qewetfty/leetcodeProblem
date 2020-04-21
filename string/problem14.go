package string

import (
	"fmt"
	"math"
)

// Write a function to find the longest common prefix string amongst an array of strings.
// If there is no common prefix, return an empty string "".
//	Example 1:
//		Input: ["flower","flow","flight"]
//		Output: "fl"
//	Example 2:
//		Input: ["dog","racecar","car"]
//		Output: ""
//	Explanation: There is no common prefix among the input strings.
//	Note:
//		All given inputs are in lowercase letters a-z.

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	length := math.MaxInt32
	for _, s := range strs {
		if len(s) < length {
			length = len(s)
		}
	}
	commonPre := ""
	for i := 0; i < length; i++ {
		char := strs[0][i]
		for j := 1; j < len(strs); j++ {
			compareChar := strs[j][i]
			if char != compareChar {
				return commonPre
			}
		}
		commonPre = commonPre + string(char)
	}
	return commonPre
}

func testProblem14() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))
}
