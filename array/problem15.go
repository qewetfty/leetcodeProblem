package array

import (
	"fmt"
	"sort"
)

// Given an array nums of n integers, are there elements a, b, c in nums such that a + b + c = 0?
// Find all unique triplets in the array which gives the sum of zero.
//
// Note:
//
// The solution set must not contain duplicate triplets.

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}
	sort.Ints(nums)
	result := make([][]int, 0)
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		num := nums[i]
		left := i + 1
		right := len(nums) - 1
		for left < right {
			sum := nums[left] + nums[right] + num
			if sum < 0 {
				left++
			} else if sum > 0 {
				right--
			} else {
				find := []int{num, nums[left], nums[right]}
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
	}
	return result
}

func testProblem15() {
	a := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(a))
}
