package dynamic_programming

import (
	"fmt"
	"github.com/leetcodeProblem/utils"
	"sort"
)

// Given a list of words, each word consists of English lowercase letters.
// Let's say word1 is a predecessor of word2 if and only if we can add exactly
// one letter anywhere in word1 to make it equal to word2. For example, "abc" is
// a predecessor of "abac".
// A word chain is a sequence of words [word_1, word_2, ..., word_k] with k >= 1,
// where word_1 is a predecessor of word_2, word_2 is a predecessor of word_3,
// and so on.
// Return the longest possible length of a word chain with words chosen from the given list of words.
//	Example 1:
//		Input: words = ["a","b","ba","bca","bda","bdca"]
//		Output: 4
//		Explanation: One of the longest word chain is "a","ba","bda","bdca".
//	Example 2:
//		Input: words = ["xbc","pcxbcf","xb","cxbc","pcxbc"]
//		Output: 5
//	Constraints:
//		1 <= words.length <= 1000
//		1 <= words[i].length <= 16
//		words[i] only consists of English lowercase letters.

func longestStrChain(words []string) int {
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) > len(words[j])
	})
	dp := make(map[string]int)
	for _, word := range words {
		dp[word] = 1
	}
	result := 1
	for _, word := range words {
		l := len(word)
		for i := 0; i < l; i++ {
			newWord := word[:i] + word[i+1:]
			if _, ok := dp[newWord]; ok {
				dp[newWord] = utils.Max(dp[newWord], dp[word]+1)
				result = utils.Max(dp[newWord], result)
			}
		}
	}
	return result
}

func testProblem1048() {
	fmt.Println(longestStrChain([]string{"abcd", "dbqca"}))
	fmt.Println(longestStrChain([]string{"a", "b", "ba", "bca", "bda", "bdca"}))
	fmt.Println(longestStrChain([]string{"xbc", "pcxbcf", "xb", "cxbc", "pcxbc"}))
}
