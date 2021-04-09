package hashTable

import (
	"fmt"
	"github.com/leetcodeProblem/utils"
)

// In an alien language, surprisingly they also use english lowercase letters,
// but possibly in a different order. The order of the alphabet is some
// permutation of lowercase letters.
// Given a sequence of words written in the alien language, and the order of the
// alphabet, return true if and only if the given words are sorted
// lexicographicaly in this alien language.
//	Example 1:
//		Input: words = ["hello","leetcode"], order = "hlabcdefgijkmnopqrstuvwxyz"
//		Output: true
//		Explanation: As 'h' comes before 'l' in this language, then the sequence is sorted.
//	Example 2:
//		Input: words = ["word","world","row"], order = "worldabcefghijkmnpqstuvxyz"
//		Output: false
//		Explanation: As 'd' comes after 'l' in this language, then words[0] > words[1], hence the sequence is unsorted.
//	Example 3:
//		Input: words = ["apple","app"], order = "abcdefghijklmnopqrstuvwxyz"
//		Output: false
//		Explanation: The first three characters "app" match, and the second string is
//					shorter (in size.) According to lexicographical rules "apple" > "app", because
//					'l' > '∅', where '∅' is defined as the blank character which is less than any
//					other character (More info).
//	Constraints:
//		1 <= words.length <= 100
//		1 <= words[i].length <= 20
//		order.length == 26
//		All characters in words[i] and order are English lowercase letters.

func isAlienSorted(words []string, order string) bool {
	charMap := make(map[byte]int)
	for i := 0; i < 26; i++ {
		charMap[order[i]] = i
	}
	l := len(words)
	for i := 0; i < l-1; i++ {
		if !compare953(words[i], words[i+1], charMap) {
			return false
		}
	}
	return true
}

func compare953(word1, word2 string, charMap map[byte]int) bool {
	l1, l2 := len(word1), len(word2)
	l := utils.Min(l1, l2)
	for i := 0; i < l; i++ {
		if charMap[word1[i]] > charMap[word2[i]] {
			return false
		} else if charMap[word1[i]] < charMap[word2[i]] {
			return true
		}
	}
	return l1 <= l
}

func testProblem953() {
	fmt.Println(isAlienSorted([]string{"hello", "leetcode"}, "hlabcdefgijkmnopqrstuvwxyz"))
	fmt.Println(isAlienSorted([]string{"word", "world", "row"}, "worldabcefghijkmnpqstuvxyz"))
	fmt.Println(isAlienSorted([]string{"apple", "app"}, "abcdefghijklmnopqrstuvwxyz"))
}
