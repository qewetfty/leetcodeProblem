package array

import (
	"fmt"
	"strconv"
)

// You are given a sorted unique integer array nums.
// Return the smallest sorted list of ranges that cover all the numbers in the
// array exactly. That is, each element of nums is covered by exactly one of the
// ranges, and there is no integer x such that x is in one of the ranges but not
// in nums.
// Each range [a,b] in the list should be output as:
//	"a->b" if a != b
//	"a" if a == b
//	Example 1:
//		Input: nums = [0,1,2,4,5,7]
//		Output: ["0->2","4->5","7"]
//		Explanation: The ranges are:
//		[0,2] --> "0->2"
//		[4,5] --> "4->5"
//		[7,7] --> "7"
//	Example 2:
//		Input: nums = [0,2,3,4,6,8,9]
//		Output: ["0","2->4","6","8->9"]
//		Explanation: The ranges are:
//		[0,0] --> "0"
//		[2,4] --> "2->4"
//		[6,6] --> "6"
//		[8,9] --> "8->9"
//	Example 3:
//		Input: nums = []
//		Output: []
//	Example 4:
//		Input: nums = [-1]
//		Output: ["-1"]
//	Example 5:
//		Input: nums = [0]
//		Output: ["0"]
//	Constraints:
//		0 <= nums.length <= 20
//		-231 <= nums[i] <= 231 - 1
//		All the values of nums are unique.

func summaryRanges(nums []int) []string {
	l, res := len(nums), make([]string, 0)
	if l == 0 {
		return res
	}
	lastNum, lastBorder := nums[0], nums[0]
	for i := 1; i < l; i++ {
		// 如果差超过了1，则需要向结果中写入值
		if nums[i]-lastNum > 1 {
			if lastBorder == lastNum {
				res = append(res, strconv.Itoa(lastBorder))
			} else {
				res = append(res, strconv.Itoa(lastBorder)+"->"+strconv.Itoa(lastNum))
			}
			lastNum, lastBorder = nums[i], nums[i]
		} else {
			lastNum = nums[i]
		}
	}
	if lastBorder == lastNum {
		res = append(res, strconv.Itoa(lastBorder))
	} else {
		res = append(res, strconv.Itoa(lastBorder)+"->"+strconv.Itoa(lastNum))
	}
	return res
}

func testProblem228() {
	fmt.Println(summaryRanges([]int{0, 1, 2, 4, 5, 7}))
	fmt.Println(summaryRanges([]int{0, 2, 3, 4, 6, 8, 9}))
	fmt.Println(summaryRanges([]int{}))
	fmt.Println(summaryRanges([]int{-1}))
	fmt.Println(summaryRanges([]int{0}))
	fmt.Println(summaryRanges([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}))
}
