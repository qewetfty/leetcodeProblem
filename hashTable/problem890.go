package hashTable

import "fmt"

// Given a list of strings words and a string pattern, return a list of words[i]
// that match pattern. You may return the answer in any order.
// A word matches the pattern if there exists a permutation of letters p so that
// after replacing every letter x in the pattern with p(x), we get the desired
// word.
// Recall that a permutation of letters is a bijection from letters to letters:
// every letter maps to another letter, and no two letters map to the same
// letter.
//	Example 1:
//		Input: words = ["abc","deq","mee","aqq","dkd","ccc"], pattern = "abb"
//		Output: ["mee","aqq"]
//		Explanation: "mee" matches the pattern because there is a permutation {a -> m, b -> e, ...}.
//					"ccc" does not match the pattern because {a -> c, b -> c, ...} is not a permutation, since a and b map to the same letter.
//	Example 2:
//		Input: words = ["a","b","c"], pattern = "a"
//		Output: ["a","b","c"]
//	Constraints:
//		1 <= pattern.length <= 20
//		1 <= words.length <= 50
//		words[i].length == pattern.length
//		pattern and words[i] are lowercase English letters.

func findAndReplacePattern(words []string, pattern string) []string {
	patternMap, l := make(map[byte]int), len(pattern)
	patternList, patternIndex := make([]int, 0), 0
	for i := 0; i < l; i++ {
		char := pattern[i]
		if patternNum, ok := patternMap[char]; ok {
			patternList = append(patternList, patternNum)
		} else {
			patternMap[char] = patternIndex
			patternList = append(patternList, patternIndex)
			patternIndex++
		}
	}
	result := make([]string, 0)
	for _, word := range words {
		wordMap := make(map[byte]int)
		wordIndex, find := 0, true
		for i := 0; i < l; i++ {
			char := word[i]
			if index, ok := wordMap[char]; ok {
				if index != patternList[i] {
					find = false
					break
				}
			} else {
				wordMap[char] = wordIndex
				wordIndex++
				if wordMap[char] != patternList[i] {
					find = false
					break
				}
			}
		}
		if find {
			result = append(result, word)
		}
	}
	return result
}

func testProblem890() {
	fmt.Println(findAndReplacePattern([]string{"abc", "deq", "mee", "aqq", "dkd", "ccc"}, "abb"))
	fmt.Println(findAndReplacePattern([]string{"a", "b", "c"}, "a"))
}
