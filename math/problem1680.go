package math

import (
	"fmt"
	"math"
	"math/bits"
)

// Given an integer n, return the decimal value of the binary string formed by
// concatenating the binary representations of 1 to n in order, modulo 109 + 7.
//	Example 1:
//		Input: n = 1
//		Output: 1
//		Explanation: "1" in binary corresponds to the decimal value 1.
//	Example 2:
//		Input: n = 3
//		Output: 27
//		Explanation: In binary, 1, 2, and 3 corresponds to "1", "10", and "11".
//		After concatenating them, we have "11011", which corresponds to the decimal value 27.
//	Example 3:
//		Input: n = 12
//		Output: 505379714
//		Explanation: The concatenation results in "1101110010111011110001001101010111100".
//		The decimal value of that is 118505380540.
//		After modulo 109 + 7, the result is 505379714.
//	Constraints:
//		1 <= n <= 105

var d = int64(1000000007)

// 使用位运算
func concatenatedBinary(n int) int {
	ans := 0
	for i := 1; i <= n; i++ {
		ans = (ans<<bits.Len(uint(i)) | i) % 1000000007
	}
	return ans
}

// 使用log计算
func concatenatedBinary2(n int) int {
	var result int64
	result = 0
	for i := 1; i <= n; i++ {
		mul := int64(math.Log2(float64(i)))
		result = result<<(mul+1) + int64(i)
		result = result % d
	}
	return int(result)
}

func testProblem1680() {
	fmt.Println(concatenatedBinary(3))
	fmt.Println(concatenatedBinary(1))
	fmt.Println(concatenatedBinary(2))
	fmt.Println(concatenatedBinary(4))
	fmt.Println(concatenatedBinary(5))
	fmt.Println(concatenatedBinary(12))
	fmt.Println(concatenatedBinary(27))
	fmt.Println(concatenatedBinary(10000))
}
