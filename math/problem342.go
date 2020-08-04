package math

import "fmt"

// Given an integer (signed 32 bits), write a function to check whether it is a power of 4.
//	Example 1:
//		Input: 16
//		Output: true
//	Example 2:
//		Input: 5
//		Output: false
//	Follow up: Could you solve it without loops/recursion?

// (4 ^ n) mod 3 = (2 ^ 2k) mod 3 =  ((3 + 1) ^ k) mod 3 = 1 (n = 2k)
func isPowerOfFour(num int) bool {
	return num > 0 && (num&(num-1)) == 0 && (num%3 == 1)
}

func testProblem342() {
	fmt.Println(isPowerOfFour(4))
	fmt.Println(isPowerOfFour(16))
	fmt.Println(isPowerOfFour(15))
	fmt.Println(isPowerOfFour(243123546))
	fmt.Println(isPowerOfFour(256))
	fmt.Println(isPowerOfFour(2198734))
}
