package binarySearch

//A positive integer is magical if it is divisible by either a or b.
//
// Given the three integers n, a, and b, return the nᵗʰ magical number. Since
//the answer may be very large, return it modulo 10⁹ + 7.
//
//
// Example 1:
// Input: n = 1, a = 2, b = 3
//Output: 2
// Example 2:
// Input: n = 4, a = 2, b = 3
//Output: 6
// Example 3:
// Input: n = 5, a = 2, b = 4
//Output: 10
// Example 4:
// Input: n = 3, a = 6, b = 4
//Output: 8
//
//
// Constraints:
//
//
// 1 <= n <= 10⁹
// 2 <= a, b <= 4 * 10⁴
//
// Related Topics Math Binary Search 👍 400 👎 87

//leetcode submit region begin(Prohibit modification and deletion)
func nthMagicalNumber(n int, a int, b int) int {
	mod := int64(1_000_000_007)
	L := a * b / gcd(a, b)
	lo, hi := int64(0), int64(1e15)
	for lo < hi {
		mid := (lo + hi) >> 1
		if mid/int64(a)+mid/int64(b)-mid/int64(L) < int64(n) {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return int(lo % mod)
}

func gcd(x, y int) int {
	if x == 0 {
		return y
	}
	return gcd(y%x, x)
}
