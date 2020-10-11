package greedy

import (
	"container/list"
	"fmt"
)

// Given a string s, remove duplicate letters so that every letter appears once
// and only once. You must make sure your result is the smallest in
// lexicographical order among all possible results.
// Note: This question is the same as 1081: https://leetcode.com/problems/smallest-subsequence-of-distinct-characters/
//	Example 1:
//		Input: s = "bcabc"
//		Output: "abc"
//	Example 2:
//		Input: s = "cbacdcbc"
//		Output: "acdb"
//	Constraints:
//		1 <= s.length <= 104
//		s consists of lowercase English letters.

func removeDuplicateLetters(s string) string {
	deque, letterMap, seenSet := new(list.List), make(map[byte]int), make(map[byte]bool)
	for i := range s {
		letterMap[s[i]] = i
	}
	for i := range s {
		char := s[i]
		if !seenSet[char] {
			for deque.Len() > 0 && char < deque.Back().Value.(byte) && letterMap[deque.Back().Value.(byte)] > i {
				delete(seenSet, deque.Back().Value.(byte))
				deque.Remove(deque.Back())
			}
			deque.PushBack(char)
			seenSet[char] = true
		}
	}
	res := make([]byte, 0)
	for deque.Len() > 0 {
		res = append(res, deque.Remove(deque.Front()).(byte))
	}
	return string(res)
}

func testProblem316() {
	fmt.Println(removeDuplicateLetters("bcabc"))
	fmt.Println(removeDuplicateLetters("cbacdcbc"))
}
