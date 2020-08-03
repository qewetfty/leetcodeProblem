package string

import (
	"fmt"
	"strings"
)

// Given a string, determine if it is a palindrome, considering only alphanumeric characters and ignoring cases.
//	Note: For the purpose of this problem, we define empty string as valid palindrome.
//	Example 1:
//		Input: "A man, a plan, a canal: Panama"
//		Output: true
//	Example 2:
//		Input: "race a car"
//		Output: false
//	Constraints:
//		s consists only of printable ASCII characters.

func isPalindrome(s string) bool {
	newStr := dealString(s)
	if len(newStr) == 0 {
		return true
	}
	i, j := 0, len(newStr)-1
	for i < j {
		if newStr[i] != newStr[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func dealString(s string) string {
	newByte := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		currByte := s[i]
		if (currByte >= 'a' && currByte <= 'z') || (currByte >= 'A' && currByte <= 'Z') || (currByte >= '0' && currByte <= '9') {
			newByte = append(newByte, currByte)
		}
	}
	return strings.ToLower(string(newByte))
}

func testProblem125() {
	fmt.Println(isPalindrome("0P"))
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome("race a car"))
}
