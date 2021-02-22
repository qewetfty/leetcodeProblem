package string

import (
	"fmt"
	"sort"
	"strings"
)

// Given a string and a string dictionary, find the longest string in the
// dictionary that can be formed by deleting some characters of the given string.
// If there are more than one possible results, return the longest word with the
// smallest lexicographical order. If there is no possible result, return the
// empty string.
//	Example 1:
//		Input:
//			s = "abpcplea", d = ["ale","apple","monkey","plea"]
//		Output:
//			"apple"
//	Example 2:
//		Input:
//			s = "abpcplea", d = ["a","b","c"]
//		Output:
//			"a"
//	Note:
//		All the strings in the input will only contain lower-case letters.
//		The size of the dictionary won't exceed 1,000.
//		The length of all the strings in the input won't exceed 1,000.

func findLongestWord(s string, d []string) string {
	sort.Slice(d, func(i, j int) bool {
		if len(d[i]) != len(d[j]) {
			return len(d[i]) > len(d[j])
		}
		return strings.Compare(d[i], d[j]) < 0
	})
	l := len(s)
	result := ""
	for _, str := range d {
		strLen := len(str)
		if strLen > l {
			continue
		}
		curLen := 0
		i, j := 0, 0
		for i < l && j < strLen && l-i >= strLen-j {
			if s[i] == str[j] {
				j++
				curLen++
			}
			i++
		}
		if curLen == strLen {
			result = str
			break
		}
	}
	return result
}

func testProblem524() {
	fmt.Println(findLongestWord("abpcplea", []string{"ale", "apple", "monkey", "plea"}))
	fmt.Println(findLongestWord("abpcplea", []string{"a", "b", "c"}))
}
