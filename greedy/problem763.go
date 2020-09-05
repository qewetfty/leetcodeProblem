package greedy

import (
	"fmt"
	"github.com/leetcodeProblem/utils"
	"strings"
)

// A string S of lowercase English letters is given. We want to partition this string into as many parts as possible so
// that each letter appears in at most one part, and return a list of integers representing the size of these parts.
//	Example 1:
//		Input: S = "ababcbacadefegdehijhklij"
//		Output: [9,7,8]
//		Explanation:
//			The partition is "ababcbaca", "defegde", "hijhklij".
//			This is a partition so that each letter appears in at most one part.
//			A partition like "ababcbacadefegde", "hijhklij" is incorrect, because it splits S into less parts.
//	Note:
//		S will have length in range [1, 500].
//		S will consist of lowercase English letters ('a' to 'z') only.

// 保存index，贪心做法
func partitionLabels(S string) []int {
	l := len(S)
	lastIndexMap := make(map[byte]int)
	for i := 0; i < l; i++ {
		lastIndexMap[S[i]] = i
	}
	result := make([]int, 0)
	start, end := 0, 0
	for i := 0; i < l; i++ {
		if lastIndexMap[S[i]] > end {
			end = lastIndexMap[S[i]]
		}
		if i == end {
			result = append(result, end-start+1)
			start = end + 1
		}
	}
	return result
}

// 每步都计算LastIndex，耗费时间
func partitionLabels2(S string) []int {
	start, l := 0, len(S)
	letterMap := make(map[string]bool)
	resultStrList := make([]string, 0)
	for start < l {
		startStr := S[start : start+1]
		letterMap[startStr] = true
		lastIndex := strings.LastIndex(S, startStr)
		for i := start + 1; i < lastIndex; i++ {
			curStr := S[i : i+1]
			if letterMap[curStr] {
				continue
			}
			curStrLastIndex := strings.LastIndex(S, curStr)
			lastIndex = utils.Max(lastIndex, curStrLastIndex)
		}
		resultStrList = append(resultStrList, S[start:lastIndex+1])
		start = lastIndex + 1
	}
	result := make([]int, len(resultStrList))
	for i, str := range resultStrList {
		result[i] = len(str)
	}
	return result
}

func testProblem763() {
	fmt.Println(partitionLabels("ababcbacadefegdehijhklij"))
	fmt.Println(partitionLabels("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"))
	fmt.Println(partitionLabels("aaaaaaaaaaaabaaaaaaaaaaaaaaaaaacccddawe"))
	fmt.Println(partitionLabels("a"))
}
