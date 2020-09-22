package array

import "fmt"

// Given an integer array of size n, find all elements that appear more than ⌊ n/3 ⌋ times.
//	Note: The algorithm should run in linear time and in O(1) space.
//	Example 1:
//		Input: [3,2,3]
//		Output: [3]
//	Example 2:
//		Input: [1,1,1,3,3,2,2,2]
//		Output: [1,2]

// 摩尔投票法
func majorityElement(nums []int) []int {
	l := len(nums)
	if l <= 1 {
		return nums
	}
	res := make([]int, 0)
	// 初始化投票人
	candidate1, count1, candidate2, count2 := nums[0], 0, nums[0], 0
	// 计票抵消阶段
	for _, num := range nums {
		if candidate1 == num {
			count1++
			continue
		}
		if candidate2 == num {
			count2++
			continue
		}
		if count1 == 0 {
			candidate1 = num
			count1++
			continue
		}
		if count2 == 0 {
			candidate2 = num
			count2++
			continue
		}
		count1--
		count2--
	}
	// 统计票数阶段
	count1, count2 = 0, 0
	for _, num := range nums {
		if candidate1 == num {
			count1++
		} else if candidate2 == num {
			count2++
		}
	}
	if count1 > l/3 {
		res = append(res, candidate1)
	}
	if count2 > l/3 {
		res = append(res, candidate2)
	}
	return res
}

func testProblem229() {
	fmt.Println(majorityElement([]int{3, 2, 3}))
	fmt.Println(majorityElement([]int{1, 1, 1, 3, 3, 2, 2, 2}))
}
