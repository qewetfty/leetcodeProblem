package array

import (
	"fmt"
	"sort"
)

// Given an array of integers nums and an integer k, return the number of unique k-diff pairs in the array.
// A k-diff pair is an integer pair (nums[i], nums[j]), where the following are true:
//	0 <= i, j < nums.length
//	i != j
//	a <= b
//	b - a == k
//	Example 1:
//		Input: nums = [3,1,4,1,5], k = 2
//		Output: 2
//		Explanation: There are two 2-diff pairs in the array, (1, 3) and (3, 5).
//			Although we have two 1s in the input, we should only return the number of unique pairs.
//	Example 2:
//		Input: nums = [1,2,3,4,5], k = 1
//		Output: 4
//		Explanation: There are four 1-diff pairs in the array, (1, 2), (2, 3), (3, 4) and (4, 5).
//	Example 3:
//		Input: nums = [1,3,1,5,4], k = 0
//		Output: 1
//		Explanation: There is one 0-diff pair in the array, (1, 1).
//	Example 4:
//		Input: nums = [1,2,4,4,3,3,0,9,2,3], k = 3
//		Output: 2
//	Example 5:
//		Input: nums = [-1,-2,-3], k = 1
//		Output: 2
//	Constraints:
//		1 <= nums.length <= 104
//		-107 <= nums[i] <= 107
//		0 <= k <= 107

func findPairs(nums []int, k int) int {
	l := len(nums)
	if l <= 1 {
		return 0
	}
	sort.Ints(nums)
	low, high, lowNum, highNum, res := 0, 1, nums[0], nums[1], 0
	for high < l {
		if nums[high]-nums[low] <= k {
			if nums[high]-nums[low] == k {
				res++
			}
			for high < l && nums[high] == highNum {
				high++
			}
			if high == l {
				break
			} else {
				highNum = nums[high]
			}
		} else if nums[high]-nums[low] > k {
			for low < l && nums[low] == lowNum {
				low++
			}
			if low == high {
				high++
			}
			if low == l {
				break
			} else {
				lowNum = nums[low]
			}
		}
	}
	return res
}

func findPairs2(nums []int, k int) int {
	numMap := make(map[int]int)
	for _, num := range nums {
		numMap[num]++
	}
	res := 0
	if k == 0 {
		for _, v := range numMap {
			if v > 1 {
				res++
			}
		}
	} else {
		for num := range numMap {
			if _, ok := numMap[num+k]; ok {
				res++
			}
		}
	}
	return res
}

func testProblem532() {
	fmt.Println(findPairs([]int{1, 3, 1, 5, 4}, 0))
	fmt.Println(findPairs([]int{3, 1, 4, 1, 5}, 2))
	fmt.Println(findPairs([]int{1, 2, 3, 4, 5}, 1))
	fmt.Println(findPairs([]int{1, 2, 4, 4, 3, 3, 0, 9, 2, 3}, 3))
	fmt.Println(findPairs([]int{-1, -2, -3}, 1))
}
