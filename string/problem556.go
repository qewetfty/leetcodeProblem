package string

import (
	"fmt"
	"math"
	"strconv"
)

// Given a positive integer n, find the smallest integer which has exactly the
// same digits existing in the integer n and is greater in value than n. If no
// such positive integer exists, return -1.
// Note that the returned integer should fit in 32-bit integer, if there is a
// valid answer but it does not fit in 32-bit integer, return -1.
//	Example 1:
//		Input: n = 12
//		Output: 21
//	Example 2:
//		Input: n = 21
//		Output: -1
//	Constraints:
//		1 <= n <= 231 - 1

func nextGreaterElement(n int) int {
	s := []byte(strconv.Itoa(n))
	l := len(s)
	i := l - 2
	for i >= 0 && s[i+1] <= s[i] {
		i--
	}
	if i < 0 {
		return -1
	}
	j := l - 1
	for j >= 0 && s[j] <= s[i] {
		j--
	}
	s[i], s[j] = s[j], s[i]
	start, end := i+1, l-1
	for start < end {
		s[start], s[end] = s[end], s[start]
		start++
		end--
	}
	res, err := strconv.Atoi(string(s))
	if err != nil || res > math.MaxInt32 {
		return -1
	}
	return res
}

func testProblem556() {
	fmt.Println(nextGreaterElement(12))
	fmt.Println(nextGreaterElement(21))
	fmt.Println(nextGreaterElement(math.MaxInt32))
	fmt.Println(nextGreaterElement(98765432))
	fmt.Println(nextGreaterElement(1234567))
	fmt.Println(nextGreaterElement(2147483647))
}
