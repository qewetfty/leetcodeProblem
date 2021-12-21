package math

//Given an integer n, return true if it is a power of two. Otherwise, return
//false.
//
// An integer n is a power of two, if there exists an integer x such that n == 2
//ˣ.
//
//
// Example 1:
//
//
//Input: n = 1
//Output: true
//Explanation: 2⁰ = 1
//
//
// Example 2:
//
//
//Input: n = 16
//Output: true
//Explanation: 2⁴ = 16
//
//
// Example 3:
//
//
//Input: n = 3
//Output: false
//
//
//
// Constraints:
//
//
// -2³¹ <= n <= 2³¹ - 1
//
//
//
//Follow up: Could you solve it without loops/recursion? Related Topics Math
//Bit Manipulation Recursion 👍 2214 👎 259

func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}
	onesCount := 0
	for n != 0 {
		if n%2 == 1 {
			onesCount++
		}
		n /= 2
	}
	return onesCount <= 1
}
