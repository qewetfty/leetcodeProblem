package greedy

import (
	"fmt"
	"github.com/leetcodeProblem/utils"
	"sort"
)

// Given an array A of integers, for each integer A[i] we need to choose either x
// = -K or x = K, and add x to A[i] (only once).
// After this process, we have some array B.
// Return the smallest possible difference between the maximum value of B and the minimum value of B.
//	Example 1:
//		Input: A = [1], K = 0
//		Output: 0
//		Explanation: B = [1]
//	Example 2:
//		Input: A = [0,10], K = 2
//		Output: 6
//		Explanation: B = [2,8]
//	Example 3:
//		Input: A = [1,3,6], K = 3
//		Output: 3
//		Explanation: B = [4,6,3]
//	Note:
//		1 <= A.length <= 10000
//		0 <= A[i] <= 10000
//		0 <= K <= 10000

func smallestRangeII(A []int, K int) int {
	l := len(A)
	sort.Ints(A)
	start, end := A[0], A[l-1]
	result := end - start
	for i := 0; i < l-1; i++ {
		a, b := A[i], A[i+1]
		high := utils.Max(end-K, a+K)
		low := utils.Min(start+K, b-K)
		result = utils.Min(result, high-low)
	}
	return result
}

func testProblem910() {
	fmt.Println(smallestRangeII([]int{1}, 0))
	fmt.Println(smallestRangeII([]int{0, 10}, 2))
	fmt.Println(smallestRangeII([]int{1, 3, 6}, 3))
}
