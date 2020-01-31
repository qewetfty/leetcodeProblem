package array

import "fmt"

// Given an array nums, write a function to move all 0's to the end of it while maintaining the
// relative order of the non-zero elements.

// Note:
//
// 1.You must do this in-place without making a copy of the array.
// 2.Minimize the total number of operations.

func moveZeroes(nums []int) {
	lastZeroPosition := 0
	for i, num := range nums {
		if num != 0 {
			nums[lastZeroPosition] = nums[i]
			lastZeroPosition++
		}
	}
	for lastZeroPosition < len(nums) {
		nums[lastZeroPosition] = 0
		lastZeroPosition++
	}
}

func testProblem283() {
	a := []int{0, 1, 0, 3, 12}
	moveZeroes(a)
	fmt.Println(a)
}
