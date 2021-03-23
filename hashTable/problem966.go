package hashTable

import (
	"fmt"
	"strings"
)

// Given a wordlist, we want to implement a spellchecker that converts a query word into a correct word.
// For a given query word, the spell checker handles two categories of spelling mistakes:
//	Capitalization: If the query matches a word in the wordlist
//	(case-insensitive), then the query word is returned with the same case as the
//	case in the wordlist.
//		Example: wordlist = ["yellow"], query = "YellOw": correct = "yellow"
//		Example: wordlist = ["Yellow"], query = "yellow": correct = "Yellow"
//		Example: wordlist = ["yellow"], query = "yellow": correct = "yellow"
//	Vowel Errors: If after replacing the vowels ('a', 'e', 'i', 'o', 'u') of the
//	query word with any vowel individually, it matches a word in the wordlist
//	(case-insensitive), then the query word is returned with the same case as the
//	match in the wordlist.
//		Example: wordlist = ["YellOw"], query = "yollow": correct = "YellOw"
//		Example: wordlist = ["YellOw"], query = "yeellow": correct = "" (no match)
//		Example: wordlist = ["YellOw"], query = "yllw": correct = "" (no match)
//	In addition, the spell checker operates under the following precedence rules:
//		When the query exactly matches a word in the wordlist (case-sensitive), you should return the same word back.
//		When the query matches a word up to capitlization, you should return the first such match in the wordlist.
//		When the query matches a word up to vowel errors, you should return the first such match in the wordlist.
//		If the query has no matches in the wordlist, you should return the empty string.
// Given some queries, return a list of words answer, where answer[i] is the correct word for query = queries[i].
//	Example 1:
//		Input: wordlist = ["KiTe","kite","hare","Hare"], queries = ["kite","Kite","KiTe","Hare","HARE","Hear","hear","keti","keet","keto"]
//		Output: ["kite","KiTe","KiTe","Hare","hare","","","KiTe","","KiTe"]
//	Note:
//		1 <= wordlist.length <= 5000
//		1 <= queries.length <= 5000
//		1 <= wordlist[i].length <= 7
//		1 <= queries[i].length <= 7
//		All strings in wordlist and queries consist only of english letters.

func spellchecker(wordlist []string, queries []string) []string {
	perfectMap, capMap, vowelMap := make(map[string]bool), make(map[string]string), make(map[string]string)
	for _, word := range wordlist {
		perfectMap[word] = true

		lowerStr := strings.ToLower(word)
		if _, ok := capMap[lowerStr]; !ok {
			capMap[lowerStr] = word
		}

		withoutVowelStr := removeVowel(lowerStr)
		if _, ok := vowelMap[withoutVowelStr]; !ok {
			vowelMap[withoutVowelStr] = word
		}
	}
	result := make([]string, 0)
	for _, query := range queries {
		if perfectMap[query] {
			result = append(result, query)
			continue
		}
		queryLower := strings.ToLower(query)
		if _, ok := capMap[queryLower]; ok {
			result = append(result, capMap[queryLower])
			continue
		}
		withoutVowel := removeVowel(queryLower)
		if _, ok := vowelMap[withoutVowel]; ok {
			result = append(result, vowelMap[withoutVowel])
			continue
		}
		result = append(result, "")
	}
	return result
}

func removeVowel(s string) string {
	resultByte, l := make([]byte, 0), len(s)
	for i := 0; i < l; i++ {
		if isVowel(s[i]) {
			resultByte = append(resultByte, '*')
		} else {
			resultByte = append(resultByte, s[i])
		}
	}
	return string(resultByte)
}

func isVowel(b byte) bool {
	return b == 'a' || b == 'e' || b == 'i' || b == 'o' || b == 'u'
}

func testProblem966() {
	fmt.Println(spellchecker([]string{"KiTe", "kite", "hare", "Hare"}, []string{"kite", "Kite", "KiTe", "Hare", "HARE", "Hear", "hear", "keti", "keet", "keto"}))
}
