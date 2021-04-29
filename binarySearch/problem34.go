package binarySearch

import "fmt"

// Given an array of integers nums sorted in ascending order, find the starting and ending position of a given target value.
// If target is not found in the array, return [-1, -1].
// Follow up: Could you write an algorithm with O(log n) runtime complexity?
//	Example 1:
//		Input: nums = [5,7,7,8,8,10], target = 8
//		Output: [3,4]
//	Example 2:
//		Input: nums = [5,7,7,8,8,10], target = 6
//		Output: [-1,-1]
//	Example 3:
//		Input: nums = [], target = 0
//		Output: [-1,-1]
//	Constraints:
//		0 <= nums.length <= 105
//		-109 <= nums[i] <= 109
//		nums is a non-decreasing array.
//		-109 <= target <= 109

func searchRange(nums []int, target int) []int {
	findLowPos := func(nums []int, target int) int {
		lo, hi := 0, len(nums)-1
		mid := (lo + hi) >> 1
		find := false
		for lo <= hi {
			mid = (lo + hi) >> 1
			if nums[mid] == target {
				find = true
				if mid == 0 || nums[mid-1] != nums[mid] {
					break
				}
				hi = mid - 1
			} else if nums[mid] < target {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}
		if !find {
			return -1
		}
		return mid
	}

	findHighPos := func(nums []int, target int) int {
		l := len(nums)
		lo, hi := 0, len(nums)-1
		mid := (lo + hi) >> 1
		find := false
		for lo <= hi {
			mid = (lo + hi) >> 1
			if nums[mid] == target {
				find = true
				if mid == l-1 || nums[mid] != nums[mid+1] {
					break
				}
				lo = mid + 1
			} else if nums[mid] < target {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}
		if !find {
			return -1
		}
		return mid
	}

	return []int{findLowPos(nums, target), findHighPos(nums, target)}
}

func testProblem34() {
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 6))
	fmt.Println(searchRange([]int{}, 0))
	fmt.Println(searchRange([]int{1}, 1))
	fmt.Println(searchRange([]int{1, 3}, 1))
	fmt.Println(searchRange([]int{2, 2}, 2))
}
