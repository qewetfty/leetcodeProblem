package dynamic_programming

import "fmt"

// Given a non-empty string s and a dictionary wordDict containing a list of non-empty words, determine if s can be
// segmented into a space-separated sequence of one or more dictionary words.
//	Note:
//		The same word in the dictionary may be reused multiple times in the segmentation.
//		You may assume the dictionary does not contain duplicate words.
//	Example 1:
//		Input: s = "leetcode", wordDict = ["leet", "code"]
//		Output: true
//		Explanation: Return true because "leetcode" can be segmented as "leet code".
//	Example 2:
//		Input: s = "applepenapple", wordDict = ["apple", "pen"]
//		Output: true
//		Explanation: Return true because "applepenapple" can be segmented as "apple pen apple".
//		             Note that you are allowed to reuse a dictionary word.
//	Example 3:
//		Input: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
//		Output: false

func wordBreak(s string, wordDict []string) bool {
	wordMap := make(map[string]bool, 0)
	for _, word := range wordDict {
		wordMap[word] = true
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := i - 1; j >= 0; j-- {
			if dp[j] && wordMap[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}

func testProblem139() {
	fmt.Println(wordBreak("catsanddog", []string{"cat", "cats", "and", "sand", "dog"}))
	fmt.Println(wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))
	fmt.Println(wordBreak("leetcode", []string{"leet", "code"}))
	fmt.Println(wordBreak("pineapplepenapple", []string{"apple", "pen", "applepen", "pine", "pineapple"}))
}
