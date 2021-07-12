package hashTable

import "fmt"

// Given two strings s and t, determine if they are isomorphic.
// Two strings s and t are isomorphic if the characters in s can be replaced to get t.
// All occurrences of a character must be replaced with another character while
// preserving the order of characters. No two characters may map to the same
// character, but a character may map to itself.
//	Example 1:
//		Input: s = "egg", t = "add"
//		Output: true
//	Example 2:
//		Input: s = "foo", t = "bar"
//		Output: false
//	Example 3:
//		Input: s = "paper", t = "title"
//		Output: true
//	Constraints:
//		1 <= s.length <= 5 * 104
//		t.length == s.length
//		s and t consist of any valid ascii character.

func isIsomorphic(s string, t string) bool {
	l := len(s)
	sMap, tMap := make(map[byte]int), make(map[byte]int)
	sIndex, tIndex := 1, 1
	for i := 0; i < l; i++ {
		sByte, tByte := s[i], t[i]
		sByteIndex, tByteIndex := sMap[sByte], tMap[tByte]
		if sByteIndex != tByteIndex {
			return false
		}
		if sByteIndex == 0 {
			sMap[sByte] = sIndex
			sIndex++
		}
		if tByteIndex == 0 {
			tMap[tByte] = tIndex
			tIndex++
		}
	}
	return true
}

func testProblem205() {
	fmt.Println(isIsomorphic("egg", "foo"))
	fmt.Println(isIsomorphic("foo", "bar"))
	fmt.Println(isIsomorphic("paper", "title"))
}
