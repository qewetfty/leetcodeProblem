package string

import (
	"fmt"
	"strings"
)

// Given a string s consists of upper/lower-case alphabets and empty space characters ' ', return the length of last
// word (last word means the last appearing word if we loop from left to right) in the string.
// If the last word does not exist, return 0.
//	Note: A word is defined as a maximal substring consisting of non-space characters only.
//	Example:
//		Input: "Hello World"
//		Output: 5

func lengthOfLastWord(s string) int {
	s = strings.TrimRight(s, " ")
	wordList := strings.Split(s, " ")
	if len(wordList) == 0 {
		return 0
	}
	lastWord := wordList[len(wordList)-1]
	return len(strings.TrimSpace(lastWord))
}

func testProblem58() {
	fmt.Println(lengthOfLastWord("Hello World"))
}
