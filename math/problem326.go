package math

// Given an integer n, return true if it is a power of three. Otherwise, return false.
// An integer n is a power of three, if there exists an integer x such that n == 3x.
//	Example 1:
//		Input: n = 27
//		Output: true
//	Example 2:
//		Input: n = 0
//		Output: false
//	Example 3:
//		Input: n = 9
//		Output: true
//	Example 4:
//		Input: n = 45
//		Output: false
//	Constraints:
//		-231 <= n <= 231 - 1

// 直接列出所有结果
var resultMap = map[int]bool{1: true, 3: true, 9: true, 27: true, 81: true, 243: true, 729: true, 2187: true, 6561: true,
	19683: true, 59049: true, 177147: true, 531441: true, 1594323: true, 4782969: true, 14348907: true, 43046721: true, 129140163: true,
	387420489: true, 1162261467: true}

func isPowerOfThree(n int) bool {
	return resultMap[n]
}

// 3486784401为3^20，看能不能整除
func isPowerOfThree2(n int) bool {
	return n > 0 && 3486784401%n == 0
}

// 一个个除看结果
func isPowerOfThree3(n int) bool {
	for n >= 3 {
		if n%3 == 0 {
			n /= 3
		} else {
			return false
		}
	}
	return n == 1
}
