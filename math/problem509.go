package math

import "fmt"

// The Fibonacci numbers, commonly denoted F(n) form a sequence, called the
// Fibonacci sequence, such that each number is the sum of the two preceding
// ones, starting from 0 and 1. That is,
//	F(0) = 0, F(1) = 1
//	F(n) = F(n - 1) + F(n - 2), for n > 1.
// Given n, calculate F(n).
//	Example 1:
//		Input: n = 2
//		Output: 1
//		Explanation: F(2) = F(1) + F(0) = 1 + 0 = 1.
//	Example 2:
//		Input: n = 3
//		Output: 2
//		Explanation: F(3) = F(2) + F(1) = 1 + 1 = 2.
//	Example 3:
//		Input: n = 4
//		Output: 3
//		Explanation: F(4) = F(3) + F(2) = 2 + 1 = 3.
//	Constraints:
//		0 <= n <= 30

func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n <= 2 {
		return 1
	}
	a, b := 1, 1
	for i := 3; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

func testProblem509() {
	fmt.Println(fib(1))
	fmt.Println(fib(2))
	fmt.Println(fib(3))
	fmt.Println(fib(4))
	fmt.Println(fib(5))
	fmt.Println(fib(6))
	fmt.Println(fib(7))
	fmt.Println(fib(8))
	fmt.Println(fib(9))
	fmt.Println(fib(10))
}
