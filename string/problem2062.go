package string

import "fmt"

// A substring is a contiguous (non-empty) sequence of characters within a string.
// A vowel substring is a substring that only consists of vowels ('a', 'e', 'i',
// 'o', and 'u') and has all five vowels present in it.
// Given a string word, return the number of vowel substrings in word.
//	Example 1:
//		Input: word = "aeiouu"
//		Output: 2
//		Explanation: The vowel substrings of word are as follows (underlined):
//			- "aeiouu"
//			- "aeiouu"
//	Example 2:
//		Input: word = "unicornarihan"
//		Output: 0
//		Explanation: Not all 5 vowels are present, so there are no vowel substrings.
//	Example 3:
//		Input: word = "cuaieuouac"
//		Output: 7
//		Explanation: The vowel substrings of word are as follows (underlined):
//			- "cuaieuouac"
//			- "cuaieuouac"
//			- "cuaieuouac"
//			- "cuaieuouac"
//			- "cuaieuouac"
//			- "cuaieuouac"
//			- "cuaieuouac"
//	Example 4:
//		Input: word = "bbaeixoubb"
//		Output: 0
//		Explanation: The only substrings that contain all five vowels also contain consonants, so there are no vowel substrings.
//	Constraints:
//		1 <= word.length <= 100
//		word consists of lowercase English letters only.

func countVowelSubstrings(word string) int {
	l := len(word)
	if l < 5 {
		return 0
	}
	var charMap map[byte]int
	result := 0
	for i := 0; i < l; i++ {
		charMap = map[byte]int{'a': 1, 'e': 1, 'i': 1, 'o': 1, 'u': 1}
		curChar := word[i]
		if _, exist := charMap[curChar]; !exist {
			continue
		}
		charMap[curChar]--
		j := i + 1
		for j < l && (word[j] == 'a' || word[j] == 'e' || word[j] == 'i' || word[j] == 'o' || word[j] == 'u') {
			if charMap[word[j]] == 1 {
				charMap[word[j]]--
			}
			find := true
			for _, value := range charMap {
				if value > 0 {
					find = false
					break
				}
			}
			if find {
				result++
			}
			j++
		}
	}
	return result
}

func testProblem2062() {
	fmt.Println(countVowelSubstrings("aeiouu"))
	fmt.Println(countVowelSubstrings("unicornarihan"))
	fmt.Println(countVowelSubstrings("cuaieuouac"))
	fmt.Println(countVowelSubstrings("bbaeixoubb"))
}
