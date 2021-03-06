package string

import (
	"fmt"
	"strings"
)

// A valid encoding of an array of words is any reference string s and array of indices indices such that:
//	words.length == indices.length
//	The reference string s ends with the '#' character.
//	For each index indices[i], the substring of s starting from indices[i] and up to (but not including) the next '#' character is equal to words[i].
// Given an array of words, return the length of the shortest reference string s possible of any valid encoding of words.
//	Example 1:
//		Input: words = ["time", "me", "bell"]
//		Output: 10
//		Explanation: A valid encoding would be s = "time#bell#" and indices = [0, 2, 5].
//			words[0] = "time", the substring of s starting from indices[0] = 0 to the next '#' is underlined in "time#bell#"
//			words[1] = "me", the substring of s starting from indices[1] = 2 to the next '#' is underlined in "time#bell#"
//			words[2] = "bell", the substring of s starting from indices[2] = 5 to the next '#' is underlined in "time#bell#"
//	Example 2:
//		Input: words = ["t"]
//		Output: 2
//		Explanation: A valid encoding would be s = "t#" and indices = [0].
//	Constraints:
//		1 <= words.length <= 2000
//		1 <= words[i].length <= 7
//		words[i] consists of only lowercase letters.

func minimumLengthEncoding(words []string) int {
	usedMap, uniqueWord := make(map[string]bool), make([]string, 0)
	for _, word := range words {
		if !usedMap[word] {
			usedMap[word] = true
			uniqueWord = append(uniqueWord, word)
		}
	}
	l := len(uniqueWord)
	for i := 0; i < l; i++ {
		word1 := uniqueWord[i]
		for j := 0; j < l; j++ {
			if i == j {
				continue
			}
			word2 := uniqueWord[j]
			if strings.HasSuffix(word1, word2) {
				usedMap[word2] = false
			}
		}
	}
	result := 0
	for word, used := range usedMap {
		if used {
			result = result + len(word) + 1
		}
	}
	return result
}

func testProblem820() {
	fmt.Println(minimumLengthEncoding([]string{"time", "me", "bell"}))
	fmt.Println(minimumLengthEncoding([]string{"t"}))
}
