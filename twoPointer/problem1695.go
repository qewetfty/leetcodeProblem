package twoPointer

import (
	"fmt"
	"github.com/leetcodeProblem/utils"
)

// You are given an array of positive integers nums and want to erase a subarray
// containing unique elements. The score you get by erasing the subarray is equal
// to the sum of its elements.
// Return the maximum score you can get by erasing exactly one subarray.
// An array b is called to be a subarray of a if it forms a contiguous
// subsequence of a, that is, if it is equal to a[l],a[l+1],...,a[r] for some
// (l,r).
//	Example 1:
//		Input: nums = [4,2,4,5,6]
//		Output: 17
//		Explanation: The optimal subarray here is [2,4,5,6].
//	Example 2:
//		Input: nums = [5,2,1,2,5,2,1,2,5]
//		Output: 8
//		Explanation: The optimal subarray here is [5,2,1] or [1,2,5].
//	Constraints:
//		1 <= nums.length <= 105
//		1 <= nums[i] <= 104

func maximumUniqueSubarray(nums []int) int {
	numAppear := make([]bool, 10000)
	l := len(nums)
	lo, hi := 0, 1
	numAppear[nums[0]] = true
	curSum := nums[0]
	result := nums[0]
	for lo < l && hi < l {
		curNum := nums[hi]
		if numAppear[curNum] {
			for nums[lo] != curNum {
				curSum -= nums[lo]
				numAppear[nums[lo]] = false
				lo++
			}
			lo++
			result = utils.Max(result, curSum)
		} else {
			curSum += curNum
			result = utils.Max(result, curSum)
			numAppear[curNum] = true
		}
		hi++
	}
	return result
}

func testProblem1695() {
	fmt.Println(maximumUniqueSubarray([]int{4, 2, 4, 5, 6}))
	fmt.Println(maximumUniqueSubarray([]int{5, 2, 1, 2, 5, 2, 1, 2, 5}))
	fmt.Println(maximumUniqueSubarray([]int{187, 470, 25, 436, 538, 809, 441, 167, 477, 110, 275, 133, 666, 345, 411, 459, 490, 266, 987, 965, 429, 166, 809, 340, 467, 318, 125, 165, 809, 610, 31, 585, 970, 306, 42, 189, 169, 743, 78, 810, 70, 382, 367, 490, 787, 670, 476, 278, 775, 673, 299, 19, 893, 817, 971, 458, 409, 886, 434}))
}
