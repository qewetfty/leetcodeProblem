package string

import (
	"fmt"
	"github.com/leetcodeProblem/utils"
	"strings"
)

// A valid number can be split up into these components (in order):
//	A decimal number or an integer.
//	(Optional) An 'e' or 'E', followed by an integer.
//	A decimal number can be split up into these components (in order):
//		(Optional) A sign character (either '+' or '-').
//		One of the following formats:
//		At least one digit, followed by a dot '.'.
//		At least one digit, followed by a dot '.', followed by at least one digit.
//		A dot '.', followed by at least one digit.
//	An integer can be split up into these components (in order):
//		(Optional) A sign character (either '+' or '-').
//		At least one digit.
// For example, all the following are valid numbers: ["2", "0089", "-0.1",
// "+3.14", "4.", "-.9", "2e10", "-90E3", "3e+7", "+6e-1", "53.5e93",
// "-123.456e789"], while the following are not valid numbers: ["abc", "1a",
// "1e", "e3", "99e2.5", "--6", "-+3", "95a54e53"].
// Given a string s, return true if s is a valid number.
//	Example 1:
//		Input: s = "0"
//		Output: true
//	Example 2:
//		Input: s = "e"
//		Output: false
//	Example 3:
//		Input: s = "."
//		Output: false
//	Example 4:
//		Input: s = ".1"
//		Output: true
//	Constraints:
//		1 <= s.length <= 20
//		s consists of only English letters (both uppercase and lowercase), digits (0-9), plus '+', minus '-', or dot '.'.

func isNumber(s string) bool {
	e := utils.Max(strings.Index(s, "e"), strings.Index(s, "E"))
	if e == -1 {
		return isSignedFloat(s)
	}
	return isSignedFloat(s[:e]) && isSignedInteger(s[e+1:])
}

func isNum(b byte) bool {
	return b >= '0' && b <= '9'
}

func isInteger(s string) bool {
	l := len(s)
	if l == 0 {
		return false
	}
	for i := 0; i < l; i++ {
		if !isNum(s[i]) {
			return false
		}
	}
	return true
}

func isSignedInteger(s string) bool {
	l := len(s)
	if l == 0 {
		return false
	}
	if s[0] == '+' || s[0] == '-' {
		return isInteger(s[1:])
	}
	return isInteger(s)
}

func isFloat(s string) bool {
	l := len(s)
	point := strings.Index(s, ".")
	if point == -1 {
		return isInteger(s)
	} else if point == 0 {
		return isInteger(s[1:])
	} else if point == l-1 {
		return isInteger(s[:l-1])
	}
	return isInteger(s[:point]) && isInteger(s[point+1:])
}

func isSignedFloat(s string) bool {
	l := len(s)
	if l == 0 {
		return false
	}
	if s[0] == '+' || s[0] == '-' {
		return isFloat(s[1:])
	}
	return isFloat(s)
}

func testProblem65() {
	fmt.Println(isNumber("1234"))
	fmt.Println(isNumber("09413"))
	fmt.Println(isNumber("+0"))
	fmt.Println(isNumber("0009"))
	fmt.Println(isNumber("++"))
}
