package binarySearch

import "fmt"

// Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.
// (i.e., [0,1,2,4,5,6,7] might become [4,5,6,7,0,1,2]).
// You are given a target value to search. If found in the array return its index, otherwise return -1.
// You may assume no duplicate exists in the array.
// Your algorithm's runtime complexity must be in the order of O(log n).
//	Example 1:
//		Input: nums = [4,5,6,7,0,1,2], target = 0
//		Output: 4
//	Example 2:
//		Input: nums = [4,5,6,7,0,1,2], target = 3
//		Output: -1

func search(nums []int, target int) int {
	lo, hi := 0, len(nums)-1
	for lo <= hi {
		mid := (lo + hi) / 2
		num := nums[mid]
		if num == target {
			return mid
			// can use xOR to replace
		} else if (target < num && nums[0] <= target) || (num < nums[0] && target < num) || (num < nums[0] && nums[0] <= target) {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return -1
}

func testProblem33() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	fmt.Println(search([]int{5, 1, 3}, 5))
	fmt.Println(search([]int{1, 3, 5}, 1))
	fmt.Println(search(nums, 0))
	fmt.Println(search(nums, 3))
}
