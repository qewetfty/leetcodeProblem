package hashTable

import (
	"fmt"
	"strings"
)

// Given a pattern and a string str, find if str follows the same pattern.
// Here follow means a full match, such that there is a bijection between a letter in pattern and a non-empty word in str.
//	Example 1:
//		Input: pattern = "abba", str = "dog cat cat dog"
//		Output: true
//	Example 2:
//		Input:pattern = "abba", str = "dog cat cat fish"
//		Output: false
//	Example 3:
//		Input: pattern = "aaaa", str = "dog cat cat dog"
//		Output: false
//	Example 4:
//		Input: pattern = "abba", str = "dog dog dog dog"
//		Output: false
//	Notes:
//		You may assume pattern contains only lowercase letters, and str contains lowercase letters that may be separated
//		by a single space.

func wordPattern(pattern string, str string) bool {
	wordList := strings.Split(str, " ")
	if len(pattern) != len(wordList) {
		return false
	}
	patternMap, wordUse := make(map[byte]string), make(map[string]bool)
	for i, word := range wordList {
		if _, ok := patternMap[pattern[i]]; ok {
			if !strings.EqualFold(word, patternMap[pattern[i]]) {
				return false
			}
		} else {
			if wordUse[word] {
				return false
			}
			patternMap[pattern[i]], wordUse[word] = word, true
		}
	}
	return true
}

func testProblem290() {
	fmt.Println(wordPattern("aaa", "aa aa aa aa"))
	fmt.Println(wordPattern("abba", "dog cat cat dog"))
	fmt.Println(wordPattern("abba", "dog cat cat fish"))
	fmt.Println(wordPattern("aaaa", "dog cat cat dog"))
	fmt.Println(wordPattern("abba", "dog dog dog dog"))
}
