package greedy

import "fmt"

// You are given an integer array nums and an integer x. In one operation, you
// can either remove the leftmost or the rightmost element from the array nums
// and subtract its value from x. Note that this modifies the array for future
// operations.
// Return the minimum number of operations to reduce x to exactly 0 if it's possible, otherwise, return -1.
//	Example 1:
//		Input: nums = [1,1,4,2,3], x = 5
//		Output: 2
//		Explanation: The optimal solution is to remove the last two elements to reduce x to zero.
//	Example 2:
//		Input: nums = [5,6,7,8,9], x = 4
//		Output: -1
//	Example 3:
//		Input: nums = [3,2,20,1,1,3], x = 10
//		Output: 5
//		Explanation: The optimal solution is to remove the last three elements and the first two elements (5 operations in total) to reduce x to zero.
//	Constraints:
//		1 <= nums.length <= 105
//		1 <= nums[i] <= 104
//		1 <= x <= 109

func minOperations(nums []int, x int) int {
	l := len(nums)
	sum := 0
	for i := 0; i < l; i++ {
		sum += nums[i]
	}
	target := sum - x
	preMap := make(map[int]int)
	preMap[0] = 0
	preSum := 0
	length := -1
	for i := 1; i <= l; i++ {
		preSum += nums[i-1]
		if _, ok := preMap[preSum]; !ok {
			preMap[preSum] = i
		}
		if _, ok := preMap[preSum-target]; ok && i-preMap[preSum-target] > length {
			length = i - preMap[preSum-target]
		}
	}
	if length != -1 {
		length = l - length
	}
	return length
}

// 前缀和O(n^2)复杂度，超时
func minOperations2(nums []int, x int) int {
	l := len(nums)
	sum, preSum := 0, make([]int, l)
	for i := 0; i < l; i++ {
		sum += nums[i]
		preSum[i] = sum
	}
	target := sum - x
	if target == 0 {
		return l
	}
	length := -1
	for i := -1; i < l; i++ {
		var preSumStart int
		if i == -1 {
			preSumStart = 0
		} else {
			preSumStart = preSum[i]
		}
		for j := i + 1; j < l; j++ {
			preSumEnd := preSum[j]
			if preSumEnd-preSumStart == target && j-i > length {
				length = j - i
			}
		}
	}
	if length != -1 {
		length = l - length
	}
	return length
}

func testProblem1658() {
	fmt.Println(minOperations([]int{1}, 1))
	fmt.Println(minOperations([]int{1, 1, 4, 2, 3}, 5))
	fmt.Println(minOperations([]int{5, 6, 7, 8, 9}, 4))
	fmt.Println(minOperations([]int{3, 2, 20, 1, 1, 3}, 10))
}
