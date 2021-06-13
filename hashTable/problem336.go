package hashTable

import (
	"fmt"
	"github.com/leetcodeProblem/utils"
)

// Given a list of unique words, return all the pairs of the distinct indices (i,
// j) in the given list, so that the concatenation of the two words words[i] +
// words[j] is a palindrome.
//	Example 1:
//		Input: words = ["abcd","dcba","lls","s","sssll"]
//		Output: [[0,1],[1,0],[3,2],[2,4]]
//		Explanation: The palindromes are ["dcbaabcd","abcddcba","slls","llssssll"]
//	Example 2:
//		Input: words = ["bat","tab","cat"]
//		Output: [[0,1],[1,0]]
//		Explanation: The palindromes are ["battab","tabbat"]
//	Example 3:
//		Input: words = ["a",""]
//		Output: [[0,1],[1,0]]
//	Constraints:
//		1 <= words.length <= 5000
//		0 <= words[i].length <= 300
//		words[i] consists of lower-case English letters.

func palindromePairs(words []string) [][]int {
	reverseWordMap := make(map[string]int)
	for i, word := range words {
		reverseWordMap[utils.Reverse(word)] = i
	}
	result := make([][]int, 0)
	for i, word := range words {
		l := len(word)
		if l == 0 {
			continue
		}
		for j := 0; j <= l; j++ {
			if isPalindrome(word, j, l-1) {
				leftId := findWord(word, 0, j-1, reverseWordMap)
				if leftId != -1 && leftId != i {
					result = append(result, []int{i, leftId})
				}
			}
			if j != 0 && isPalindrome(word, 0, j-1) {
				rightId := findWord(word, j, l-1, reverseWordMap)
				if rightId != -1 && rightId != i {
					result = append(result, []int{rightId, i})
				}
			}
		}
	}
	return result
}

func isPalindrome(s string, left, right int) bool {
	for i := 0; i < (right-left+1)>>1; i++ {
		if s[left+i] != s[right-i] {
			return false
		}
	}
	return true
}

func findWord(s string, left, right int, reverseWordMap map[string]int) int {
	if v, ok := reverseWordMap[s[left:right+1]]; ok {
		return v
	}
	return -1
}

func testProblem336() {
	fmt.Println(palindromePairs([]string{"abcd", "dcba", "lls", "s", "sssll"}))
	fmt.Println(palindromePairs([]string{"bat", "tab", "cat"}))
	fmt.Println(palindromePairs([]string{"a", ""}))
	fmt.Println(palindromePairs([]string{"a", "abc", "aba", ""}))
}
