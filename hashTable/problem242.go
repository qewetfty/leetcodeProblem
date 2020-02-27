package hashTable

import "fmt"

// Given two strings s and t , write a function to determine if t is an anagram of s.
//	Example 1:
//		Input: s = "anagram", t = "nagaram"
//		Output: true
//	Example 2:
//		Input: s = "rat", t = "car"
//		Output: false
//	Note:
//		You may assume the string contains only lowercase alphabets.
//	Follow up:
//		What if the inputs contain unicode characters? How would you adapt your solution to such case?

func isAnagram(s string, t string) bool {
	strMap := make(map[int32]int, 0)
	for _, char := range s {
		if _, exist := strMap[char]; exist {
			strMap[char]++
		} else {
			strMap[char] = 1
		}
	}
	for _, char := range t {
		if _, exist := strMap[char]; !exist {
			return false
		}
		strMap[char]--
	}
	for _, v := range strMap {
		if v != 0 {
			return false
		}
	}
	return true
}

func testProblem242() {
	s := "anagram"
	t := "nagaram"
	fmt.Println(isAnagram(s, t))
	s = "rat"
	t = "car"
	fmt.Println(isAnagram(s, t))
}
