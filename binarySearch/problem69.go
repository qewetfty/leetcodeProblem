package binarySearch

import "fmt"

// Implement int sqrt(int x).
// Compute and return the square root of x, where x is guaranteed to be a non-negative integer.
// Since the return type is an integer, the decimal digits are truncated and only the integer part of the result is returned.
//	Example 1:
//		Input: 4
//		Output: 2
//	Example 2:
//		Input: 8
//		Output: 2
//	Explanation: The square root of 8 is 2.82842..., and since
//             	the decimal part is truncated, 2 is returned.

func mySqrt(x int) int {
	if x <= 1 {
		return x
	}
	lo, hi := 0, x
	for lo <= hi {
		mid := (lo + hi) / 2
		if mid*mid == x {
			return mid
		} else if mid*mid > x {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return hi
}

func testProblem69() {
	fmt.Println(mySqrt(4))
	fmt.Println(mySqrt(2))
	fmt.Println(mySqrt(8))
	fmt.Println(mySqrt(13))
	fmt.Println(mySqrt(11111233))
	fmt.Println(mySqrt(2501))
}
