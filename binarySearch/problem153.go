package binarySearch

import "fmt"

// Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.
// (i.e.,  [0,1,2,4,5,6,7] might become  [4,5,6,7,0,1,2]).
// Find the minimum element.
// You may assume no duplicate exists in the array.
//	Example 1:
//		Input: [3,4,5,1,2]
//		Output: 1
//	Example 2:
//		Input: [4,5,6,7,0,1,2]
//		Output: 0

func findMin(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	min := nums[0]
	lo, hi := 0, len(nums)-1
	for lo <= hi {
		mid := (lo + hi) / 2
		if mid-1 >= 0 && nums[mid] < nums[mid-1] {
			min = nums[mid]
			break
		} else if nums[mid] >= nums[0] {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return min
}

func testProblem153() {
	fmt.Println(findMin([]int{3, 4, 5, 1, 2}))
	fmt.Println(findMin([]int{3, 4, 5, 7, 8, 1, 2}))
	fmt.Println(findMin([]int{4, 5, 6, 7, 0, 1, 2}))
	fmt.Println(findMin([]int{7, 8, 9, 0, 1, 2, 3, 4, 5, 6}))
	fmt.Println(findMin([]int{0, 1, 2, 3, 4, 5, 6}))
}
