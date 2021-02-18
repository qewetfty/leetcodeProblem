package math

import "fmt"

// A sequence of numbers is called arithmetic if it consists of at least three
// elements and if the difference between any two consecutive elements is the
// same.
// For example, these are arithmetic sequences:
//	1, 3, 5, 7, 9
//	7, 7, 7, 7
//	3, -1, -5, -9
// The following sequence is not arithmetic.
//	1, 1, 2, 5, 7
// A zero-indexed array A consisting of N numbers is given. A slice of that array
// is any pair of integers (P, Q) such that 0 <= P < Q < N.
// A slice (P, Q) of the array A is called arithmetic if the sequence:
// A[P], A[P + 1], ..., A[Q - 1], A[Q] is arithmetic. In particular, this means that P + 1 < Q.
// The function should return the number of arithmetic slices in the array A.
//	Example:
//		A = [1, 2, 3, 4]
//		return: 3, for 3 arithmetic slices in A: [1, 2, 3], [2, 3, 4] and [1, 2, 3, 4] itself.

func numberOfArithmeticSlices(A []int) int {
	l := len(A)
	if l < 3 {
		return 0
	}
	diff, diffCount := A[1]-A[0], 2
	result := 0
	for i := 2; i < l; i++ {
		curDiff := A[i] - A[i-1]
		if curDiff == diff {
			diffCount++
		} else {
			if diffCount >= 3 {
				result = result + (diffCount-2)*(diffCount-1)/2
			}
			diffCount, diff = 2, curDiff
		}
	}
	if diffCount >= 3 {
		result = result + (diffCount-2)*(diffCount-1)/2
	}
	return result
}

func testProblem413() {
	fmt.Println(numberOfArithmeticSlices([]int{1, 2, 3, 4}))
	fmt.Println(numberOfArithmeticSlices([]int{1, 2, 3, 4, 5}))
	fmt.Println(numberOfArithmeticSlices([]int{1}))
	fmt.Println(numberOfArithmeticSlices([]int{1, 2}))
}
