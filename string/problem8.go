package string

import (
	"fmt"
	"math"
	"strings"
)

// Implement atoi which converts a string to an integer.
// The function first discards as many whitespace characters as necessary
// until the first non-whitespace character is found.
// Then, starting from this character, takes an optional initial plus or minus
// sign followed by as many numerical digits as possible, and interprets them as a numerical value.

// The string can contain additional characters after those that form the integral number,
// which are ignored and have no effect on the behavior of this function.

// If the first sequence of non-whitespace characters in str is not a valid integral number,
// or if no such sequence exists because either str is empty or
// it contains only whitespace characters, no conversion is performed.

//If no valid conversion could be performed, a zero value is returned.

func myAtoi(str string) int {
	str = strings.Trim(str, " ")
	if len(str) == 0 {
		return 0
	}
	start := str[0]
	if !(start == '-' || start == '+' || ('0' <= start && start <= '9')) {
		return 0
	}
	numString := ""
	if start == '-' || start == '+' {
		numString += string(start)
		str = str[1:]
	}
	for _, ch := range []byte(str) {
		if !('0' <= ch && ch <= '9') {
			break
		}
		numString += string(ch)
	}
	return string2int(numString)
}

func string2int(s string) int {
	s0 := s
	if s0[0] == '-' || s[0] == '+' {
		s0 = s0[1:]
	}
	for _, ch := range []byte(s0) {
		if ch != '0' {
			break
		}
		s0 = s0[1:]
	}
	sLen := len(s0)
	if sLen == 0 {
		return 0
	}
	if !(0 < sLen && sLen < 19) {
		if s[0] == '-' {
			return math.MinInt32
		} else {
			return math.MaxInt32
		}
	}
	n := int64(0)
	for _, ch := range []byte(s0) {
		ch -= '0'
		n = n*10 + int64(ch)
	}
	if s[0] == '-' {
		n = -n
	}
	if n < math.MinInt32 {
		return math.MinInt32
	}
	if n > math.MaxInt32 {
		return math.MaxInt32
	}
	return int(n)
}

func testProblem8() {
	input1 := "   -420"
	fmt.Println(myAtoi(input1))
	input2 := "3541500 with words"
	fmt.Println(myAtoi(input2))
	input3 := "words and 8124"
	fmt.Println(myAtoi(input3))
	input4 := "-9999999999999"
	fmt.Println(myAtoi(input4))
	input5 := "+124"
	fmt.Println(myAtoi(input5))
	input6 := " 0000000000012345678"
	fmt.Println(myAtoi(input6))
	input7 := "9223372036854775808"
	fmt.Println(myAtoi(input7))
}
