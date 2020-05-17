package binarySearch

import "fmt"

// Given a positive integer num, write a function which returns True if num is a perfect square else False.
//	Follow up: Do not use any built-in library function such as sqrt.
//	Example 1:
//		Input: num = 16
//		Output: true
//	Example 2:
//		Input: num = 14
//		Output: false
//	Constraints:
//		1 <= num <= 2^31 - 1

func isPerfectSquare(num int) bool {
	lo, hi := 0, num
	for lo <= hi {
		mid := (lo + hi) / 2
		if mid*mid == num {
			return true
		} else if mid*mid < num {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return false
}

func testProblem367() {
	fmt.Println(isPerfectSquare(16))
	fmt.Println(isPerfectSquare(14))
	fmt.Println(isPerfectSquare(1))
	fmt.Println(isPerfectSquare(2))
	fmt.Println(isPerfectSquare(4))
}
