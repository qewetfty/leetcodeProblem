package bitmap

import (
	"fmt"
	"github.com/leetcodeProblem/utils"
)

// Given a string array words, return the maximum value of length(word[i]) *
// length(word[j]) where the two words do not share common letters. If no such
// two words exist, return 0.
//	Example 1:
//		Input: words = ["abcw","baz","foo","bar","xtfn","abcdef"]
//		Output: 16
//		Explanation: The two words can be "abcw", "xtfn".
//	Example 2:
//		Input: words = ["a","ab","abc","d","cd","bcd","abcd"]
//		Output: 4
//		Explanation: The two words can be "ab", "cd".
//	Example 3:
//		Input: words = ["a","aa","aaa","aaaa"]
//		Output: 0
//		Explanation: No such pair of words.
//	Constraints:
//		2 <= words.length <= 1000
//		1 <= words[i].length <= 1000
//		words[i] consists only of lowercase English letters.

func maxProduct(words []string) int {
	wordCountList := make([]wordCount, 0)
	for _, word := range words {
		l := len(word)
		charBit := 0
		for i := 0; i < l; i++ {
			char := word[i]
			charBit = charBit | (1 << (char - 'a'))
		}
		wordCountList = append(wordCountList, wordCount{length: len(word), charBit: charBit})
	}
	result := 0
	l := len(wordCountList)
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			w1, w2 := wordCountList[i], wordCountList[j]
			if w1.charBit&w2.charBit == 0 {
				result = utils.Max(w1.length*w2.length, result)
			}
		}
	}
	return result
}

type wordCount struct {
	length  int
	charBit int
}

func testProblem318() {
	fmt.Println(maxProduct([]string{"abcw", "baz", "foo", "bar", "xtfn", "abcdef"}))
	fmt.Println(maxProduct([]string{"a", "ab", "abc", "d", "cd", "bcd", "abcd"}))
	fmt.Println(maxProduct([]string{"a", "aa", "aaa", "aaaa"}))
}
