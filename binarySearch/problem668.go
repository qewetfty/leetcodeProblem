package binarySearch

import "fmt"

// Nearly everyone has used the Multiplication Table. The multiplication table of
// size m x n is an integer matrix mat where mat[i][j] == i * j (1-indexed).
// Given three integers m, n, and k, return the kth smallest element in the m x n multiplication table.
//	Example 1:
//		Input: m = 3, n = 3, k = 5
//		Output: 3
//		Explanation: The 5th smallest number is 3.
//	Example 2:
//		Input: m = 2, n = 3, k = 6
//		Output: 6
//		Explanation: The 6th smallest number is 6.
//	Constraints:
//		1 <= m, n <= 3 * 104
//		1 <= k <= m * n

func findKthNumber668(m int, n int, k int) int {
	lo, hi := 1, m*n
	result := 1
	for lo <= hi {
		mid := (lo + hi) >> 1
		if binarySearch668(m, n, k, mid) {
			result = mid
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return result
}

func binarySearch668(m, n, k, target int) bool {
	j, count := n, 0
	for i := 1; i <= m; i++ {
		for i*j > target {
			j--
		}
		count += j
		if count >= k {
			return true
		}
	}
	return false
}

func testProblem668() {
	fmt.Println(findKthNumber668(3, 3, 5))
	fmt.Println(findKthNumber668(2, 3, 6))
}
