package array

import (
	"fmt"
	"sort"
)

// Given an array nums of n integers and an integer target, are there elements a, b, c,
// and d in nums such that a + b + c + d = target? Find all unique quadruplets in the array which gives
// the sum of target.
//
// Note:
//
// The solution set must not contain duplicate quadruplets.

func nSum(nums []int, target int, n int, isSort bool) [][]int {
	result := make([][]int, 0)
	length := len(nums)
	if isSort {
		sort.Ints(nums)
	}
	if n == 2 {
		left := 0
		right := length - 1
		for left < right {
			sum := nums[left] + nums[right]
			if sum < target {
				left++
			} else if sum > target {
				right--
			} else {
				find := []int{nums[left], nums[right]}
				result = append(result, find)
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			}
		}
		return result
	}
	for i := 0; i < length; i++ {
		if n*nums[i] > target {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		lastFind := nSum(nums[i+1:], target-nums[i], n-1, false)
		for j := 0; j < len(lastFind); j++ {
			lastFind[j] = append(lastFind[j], nums[i])
			result = append(result, lastFind[j])
		}
	}
	return result
}

func fourSum(nums []int, target int) [][]int {
	return nSum(nums, target, 4, true)
}

func testProblem18() {
	a := []int{1, 0, -1, 0, -2, 2}
	fmt.Println(fourSum(a, 0))
}
