package string

import (
	"fmt"
	"regexp"
)

// Given a word, you need to judge whether the usage of capitals in it is right or not.
// We define the usage of capitals in a word to be right when one of the following cases holds:
//	All letters in this word are capitals, like "USA".
//	All letters in this word are not capitals, like "leetcode".
//	Only the first letter in this word is capital, like "Google".
//	Otherwise, we define that this word doesn't use capitals in a right way.
//	Example 1:
//		Input: "USA"
//		Output: True
//	Example 2:
//		Input: "FlaG"
//		Output: False
//Note: The input will be a non-empty word consisting of uppercase and lowercase latin letters.

// Regex method
func detectCapitalUse(word string) bool {
	regex := "^[A-Z]+$|^[a-z]+$|^[A-Z][a-z]+$"
	res, _ := regexp.MatchString(regex, word)
	return res
}

// word one by one method
func detectCapitalUse2(word string) bool {
	if isCapital(word) {
		if isUpper(word[1:]) || isLower(word[1:]) {
			return true
		}
		return false
	}
	return isLower(word)
}

func isCapital(word string) bool {
	char := word[0]
	return char >= 'A' && char <= 'Z'
}

func isUpper(word string) bool {
	for _, char := range word {
		if char < 'A' || char > 'Z' {
			return false
		}
	}
	return true
}

func isLower(word string) bool {
	for _, char := range word {
		if char < 'a' || char > 'z' {
			return false
		}
	}
	return true
}

func testProblem520() {
	fmt.Println(detectCapitalUse2("FlaG"))
	fmt.Println(detectCapitalUse2("USA"))
	fmt.Println(detectCapitalUse2("leetcode"))
	fmt.Println(detectCapitalUse2("Google"))
}
