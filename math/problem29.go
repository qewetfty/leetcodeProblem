package math

import (
	"fmt"
	"github.com/leetcodeProblem/utils"
	"math"
)

// Given two integers dividend and divisor, divide two integers without using
// multiplication, division, and mod operator.
// Return the quotient after dividing dividend by divisor.
// The integer division should truncate toward zero, which means losing its
// fractional part. For example, truncate(8.345) = 8 and truncate(-2.7335) = -2.
//	Note:
// 		Assume we are dealing with an environment that could only store integers
// 		within the 32-bit signed integer range: [−231, 231 − 1]. For this problem,
// 		assume that your function returns 231 − 1 when the division result overflows.
//	Example 1:
//		Input: dividend = 10, divisor = 3
//		Output: 3
//		Explanation: 10/3 = truncate(3.33333..) = 3.
//	Example 2:
//		Input: dividend = 7, divisor = -3
//		Output: -2
//		Explanation: 7/-3 = truncate(-2.33333..) = -2.
//	Example 3:
//		Input: dividend = 0, divisor = 1
//		Output: 0
//	Example 4:
//		Input: dividend = 1, divisor = 1
//		Output: 1
//	Constraints:
//		-2^31 <= dividend, divisor <= 2^31 - 1
//		divisor != 0

func divide(dividend int, divisor int) int {
	var positive bool
	if (dividend >= 0 && divisor >= 0) || (dividend < 0 && divisor < 0) {
		positive = true
	} else {
		positive = false
	}
	did, divi := utils.Abs(dividend), utils.Abs(divisor)
	result := div(did, divi)
	if !positive {
		result = -result
	}
	result = utils.Min(math.MaxInt32, result)
	return result
}

func div(a, b int) int {
	if a < b {
		return 0
	}
	count, tb := 1, b
	for tb+tb <= a {
		count = count + count
		tb = tb + tb
	}
	return count + div(a-tb, b)
}

func testProblem29() {
	fmt.Println(divide(10, 3))
	fmt.Println(divide(7, -3))
	fmt.Println(divide(0, 1))
	fmt.Println(divide(1, 1))
	fmt.Println(divide(-2147483648, -1))
}
