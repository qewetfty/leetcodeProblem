package math

import (
	"fmt"
	"math"
)

// Given a non-negative integer c, decide whether there're two integers a and b such that a2 + b2 = c.
//	Example 1:
//		Input: c = 5
//		Output: true
//		Explanation: 1 * 1 + 2 * 2 = 5
//	Example 2:
//		Input: c = 3
//		Output: false
//	Example 3:
//		Input: c = 4
//		Output: true
//	Example 4:
//		Input: c = 2
//		Output: true
//	Example 5:
//		Input: c = 1
//		Output: true
//	Constraints:
//		0 <= c <= 231 - 1

func judgeSquareSum(c int) bool {
	left, right := 0, int(math.Sqrt(float64(c)))
	for left <= right {
		sum := left*left + right*right
		if sum == c {
			return true
		} else if sum < c {
			left++
		} else {
			right--
		}
	}
	return false
}

func testProblem633() {
	fmt.Println(judgeSquareSum(1))
	fmt.Println(judgeSquareSum(2))
	fmt.Println(judgeSquareSum(3))
	fmt.Println(judgeSquareSum(4))
	fmt.Println(judgeSquareSum(5))
}
