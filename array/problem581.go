package array

import (
	"fmt"
	"github.com/leetcodeProblem/utils"
	"sort"
)

// Given an integer array nums, you need to find one continuous subarray that if
// you only sort this subarray in ascending order, then the whole array will be
// sorted in ascending order.
// Return the shortest such subarray and output its length.
//	Example 1:
//		Input: nums = [2,6,4,8,10,9,15]
//		Output: 5
//		Explanation: You need to sort [6, 4, 8, 10, 9] in ascending order to make the whole array sorted in ascending order.
//	Example 2:
//		Input: nums = [1,2,3,4]
//		Output: 0
//	Example 3:
//		Input: nums = [1]
//		Output: 0
//	Constraints:
//		1 <= nums.length <= 104
//		-105 <= nums[i] <= 105

// 使用栈，O(n)复杂度
func findUnsortedSubarray(nums []int) int {
	l := len(nums)
	stack, start, end := make([]int, 0), l-1, 0
	for i := 0; i < l; i++ {
		j := len(stack)
		for j > 0 && nums[stack[j-1]] > nums[i] {
			start = utils.Min(start, stack[j-1])
			j--
		}
		stack = stack[:j]
		stack = append(stack, i)
	}
	stack = make([]int, 0)
	for i := l - 1; i >= 0; i-- {
		j := len(stack)
		for j > 0 && nums[stack[j-1]] < nums[i] {
			end = utils.Max(end, stack[j-1])
			j--
		}
		stack = stack[:j]
		stack = append(stack, i)
	}
	if end <= start {
		return 0
	}
	return end - start + 1
}

// 排序法
func findUnsortedSubarray2(nums []int) int {
	l := len(nums)
	sortedNum := make([]int, l)
	copy(sortedNum, nums)
	sort.Ints(sortedNum)
	start, end := l-1, 0
	for i := 0; i < l; i++ {
		if sortedNum[i] != nums[i] {
			start = utils.Min(start, i)
			end = utils.Max(end, i)
		}
	}
	if end <= start {
		return 0
	}
	return end - start + 1
}

// 暴力法
func findUnsortedSubarray3(nums []int) int {
	l := len(nums)
	startPoint, endPoint := -1, -1
	for i := 0; i < l; i++ {
		for j := 0; j < i; j++ {
			if nums[j] > nums[i] {
				if startPoint == -1 {
					startPoint = j
				} else {
					startPoint = utils.Min(startPoint, j)
				}
				if endPoint == -1 {
					endPoint = i
				} else {
					endPoint = utils.Max(endPoint, i)
				}
			}
		}
	}
	if startPoint == -1 {
		return 0
	}
	return endPoint - startPoint + 1
}

func testProblem581() {
	fmt.Println(findUnsortedSubarray([]int{2, 6, 4, 8, 10, 9, 15}))
	fmt.Println(findUnsortedSubarray([]int{1, 2, 3, 4}))
	fmt.Println(findUnsortedSubarray([]int{1}))
	fmt.Println(findUnsortedSubarray([]int{1, 3, 2, 4}))
}
