package dynamic_programming

import "fmt"

// Given a non-empty string s and a dictionary wordDict containing a list of non-empty words, add spaces in s to
// construct a sentence where each word is a valid dictionary word. Return all such possible sentences.
//	Note:
//		The same word in the dictionary may be reused multiple times in the segmentation.
//		You may assume the dictionary does not contain duplicate words.
//	Example 1:
//		Input:
//			s = "catsanddog"
//			wordDict = ["cat", "cats", "and", "sand", "dog"]
//		Output:
//			[
//			  "cats and dog",
//			  "cat sand dog"
//			]
//	Example 2:
//		Input:
//			s = "pineapplepenapple"
//			wordDict = ["apple", "pen", "applepen", "pine", "pineapple"]
//		Output:
//			[
//			  "pine apple pen apple",
//			  "pineapple pen apple",
//			  "pine applepen apple"
//			]
//		Explanation: Note that you are allowed to reuse a dictionary word.
//	Example 3:
//		Input:
//			s = "catsandog"
//			wordDict = ["cats", "dog", "sand", "and", "cat"]
//		Output:
//			[]

func wordBreak2(s string, wordDict []string) []string {
	// first judge if it can be broken to deal with the time over limit method
	if !judge(s, wordDict) {
		return []string{}
	}
	wordMap := make(map[string]bool, 0)
	for _, word := range wordDict {
		wordMap[word] = true
	}
	dp := make([][]string, len(s)+1)
	dp[0] = []string{""}
	for i := 1; i <= len(s); i++ {
		dp[i] = make([]string, 0)
		for j := i - 1; j >= 0; j-- {
			if len(dp[j]) > 0 && wordMap[s[j:i]] {
				for _, res := range dp[j] {
					if res == "" {
						dp[i] = append(dp[i], s[j:i])
					} else {
						dp[i] = append(dp[i], res+" "+s[j:i])
					}
				}
			}
		}
	}
	return dp[len(s)]
}

// same function as problem 139
func judge(s string, wordDict []string) bool {
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

func testProblem140() {
	b := wordBreak2("aaaaaaa", []string{"aaaa", "aa", "a"})
	fmt.Println(b)
	a := wordBreak2("catsanddog", []string{"cat", "cats", "and", "sand", "dog"})
	fmt.Println(a)
	fmt.Println(wordBreak2("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))
	fmt.Println(wordBreak2("leetcode", []string{"leet", "code"}))
	fmt.Println(wordBreak2("pineapplepenapple", []string{"apple", "pen", "applepen", "pine", "pineapple"}))
}
