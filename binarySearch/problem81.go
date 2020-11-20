package binarySearch

import "fmt"

// Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.
// (i.e., [0,0,1,2,2,5,6] might become [2,5,6,0,0,1,2]).
// You are given a target value to search. If found in the array return true, otherwise return false.
//	Example 1:
//		Input: nums = [2,5,6,0,0,1,2], target = 0
//		Output: true
//	Example 2:
//		Input: nums = [2,5,6,0,0,1,2], target = 3
//		Output: false
//	Follow up:
//		This is a follow up problem to Search in Rotated Sorted Array, where nums may contain duplicates.
//		Would this affect the run-time complexity? How and why?

func search81(nums []int, target int) bool {
	l := len(nums)
	if l == 0 {
		return false
	}
	lo, hi := 0, l-1
	for lo <= hi {
		mid := (lo + hi) / 2
		if nums[mid] == target {
			return true
		}
		// 如果mid和开始相等，不知道哪边有序，直接lo++
		if nums[lo] == nums[mid] {
			lo++
			continue
		}
		// 如果mid比开始大，说明左半边有序
		if nums[lo] < nums[mid] {
			if nums[lo] <= target && target < nums[mid] {
				hi = mid - 1
			} else {
				lo = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[hi] {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}
	}
	return false
}

// 循环写法
func searchWithCycle(nums []int, target int) bool {
	for _, num := range nums {
		if num == target {
			return true
		}
	}
	return false
}

func testProblem81() {
	fmt.Println(search81([]int{1, 3, 1, 1, 1}, 3))
	fmt.Println(search81([]int{3, 1}, 1))
	fmt.Println(search81([]int{2, 5, 6, 0, 0, 1, 2}, 0))
	fmt.Println(search81([]int{2, 5, 6, 0, 0, 1, 2}, 3))
}
