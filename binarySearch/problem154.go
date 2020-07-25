package binarySearch

import "fmt"

// Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.
// (i.e.,  [0,1,2,4,5,6,7] might become  [4,5,6,7,0,1,2]).
// Find the minimum element.
// The array may contain duplicates.
//	Example 1:
//		Input: [1,3,5]
//		Output: 1
//	Example 2:
//		Input: [2,2,2,0,1]
//		Output: 0
//	Note:
//		This is a follow up problem to Find Minimum in Rotated Sorted Array.
//		Would allow duplicates affect the run-time complexity? How and why?

func findMin154(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}
	lo, hi := 0, l-1
	for lo < hi {
		mid := (lo + hi) / 2
		if nums[mid] < nums[hi] {
			hi = mid
		} else if nums[mid] > nums[hi] {
			lo = mid + 1
		} else {
			hi = hi - 1
		}
	}
	return nums[lo]
}

func testProblem154() {
	fmt.Println(findMin154([]int{10, 1, 10, 10, 10}))
	fmt.Println(findMin154([]int{2, 2, 2, 0, 1}))
	fmt.Println(findMin154([]int{1, 3, 5}))
}
