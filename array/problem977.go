package array

import "fmt"

// Given an array of integers A sorted in non-decreasing order,
// return an array of the squares of each number, also in sorted non-decreasing order.

func testProblem977() {
	nums1 := []int{-4, -1, 0, 3, 10}
	fmt.Println(sortedSquares(nums1))
	nums2 := []int{-7, -3, 2, 3, 11}
	fmt.Println(sortedSquares(nums2))
}

func sortedSquares(A []int) []int {
	length := len(A)
	low := 0
	hi := length - 1
	pos := hi
	res := make([]int, length)
	for low <= hi {
		number1 := A[low]
		if number1 < 0 {
			number1 = -number1
		}
		number2 := A[hi]
		if number2 < 0 {
			number2 = -number2
		}
		if number1 < number2 {
			res[pos] = number2 * number2
			pos--
			hi--
		} else {
			res[pos] = number1 * number1
			pos--
			low++
		}
	}
	return res
}
