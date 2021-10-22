package hashTable

import (
	"fmt"
	"sort"
)

// Given a string s, sort it in decreasing order based on the frequency of the
// characters. The frequency of a character is the number of times it appears in
// the string.
// Return the sorted string. If there are multiple answers, return any of them.
//	Example 1:
//		Input: s = "tree"
//		Output: "eert"
//		Explanation: 'e' appears twice while 'r' and 't' both appear once.
//					 So 'e' must appear before both 'r' and 't'. Therefore "eetr" is also a valid answer.
//	Example 2:
//		Input: s = "cccaaa"
//		Output: "aaaccc"
//		Explanation: Both 'c' and 'a' appear three times, so both "cccaaa" and "aaaccc" are valid answers.
//					 Note that "cacaca" is incorrect, as the same characters must be together.
//	Example 3:
//		Input: s = "Aabb"
//		Output: "bbAa"
//		Explanation: "bbaA" is also a valid answer, but "Aabb" is incorrect.
//					 Note that 'A' and 'a' are treated as two different characters.
//	Constraints:
//		1 <= s.length <= 5 * 105
//		s consists of uppercase and lowercase English letters and digits.

func frequencySort(s string) string {
	charMap := make(map[byte]int)
	l := len(s)
	for i := 0; i < l; i++ {
		charMap[s[i]]++
	}
	charList := make([]fre, 0)
	for char, count := range charMap {
		charList = append(charList, fre{
			char:  char,
			count: count,
		})
	}
	sort.Slice(charList, func(i, j int) bool {
		return charList[i].count > charList[j].count
	})
	resByte := make([]byte, 0)
	for _, f := range charList {
		for i := 0; i < f.count; i++ {
			resByte = append(resByte, f.char)
		}
	}
	return string(resByte)
}

type fre struct {
	char  byte
	count int
}

func testProblem451() {
	fmt.Println(frequencySort("tree"))
	fmt.Println(frequencySort("cccaaa"))
	fmt.Println(frequencySort("Aabb"))
}
