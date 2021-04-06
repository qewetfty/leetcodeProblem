package greedy

import (
	"fmt"
	"github.com/leetcodeProblem/utils"
	"sort"
)

// You are given two positive integer arrays nums1 and nums2, both of length n.
// The absolute sum difference of arrays nums1 and nums2 is defined as the sum of
// |nums1[i] - nums2[i]| for each 0 <= i < n (0-indexed).
// You can replace at most one element of nums1 with any other element in nums1
// to minimize the absolute sum difference.
// Return the minimum absolute sum difference after replacing at most one element
// in the array nums1. Since the answer may be large, return it modulo 109 + 7.
//	|x| is defined as:
//		x if x >= 0, or
//		-x if x < 0.
//	Example 1:
//		Input: nums1 = [1,7,5], nums2 = [2,3,5]
//		Output: 3
//		Explanation: There are two possible optimal solutions:
//					- Replace the second element with the first: [1,7,5] => [1,1,5], or
//					- Replace the second element with the third: [1,7,5] => [1,5,5].
//					Both will yield an absolute sum difference of |1-2| + (|1-3| or |5-3|) + |5-5| = 3.
//	Example 2:
//		Input: nums1 = [2,4,6,8,10], nums2 = [2,4,6,8,10]
//		Output: 0
//		Explanation: nums1 is equal to nums2 so no replacement is needed. This will result in an
//					absolute sum difference of 0.
//	Example 3:
//		Input: nums1 = [1,10,4,4,2,7], nums2 = [9,3,5,1,7,4]
//		Output: 20
//		Explanation: Replace the first element with the second: [1,10,4,4,2,7] => [10,10,4,4,2,7].
//					This yields an absolute sum difference of |10-9| + |10-3| + |4-5| + |4-1| + |2-7| + |7-4| = 20
//	Constraints:
//		n == nums1.length
//		n == nums2.length
//		1 <= n <= 105
//		1 <= nums1[i], nums2[i] <= 105

func minAbsoluteSumDiff(nums1 []int, nums2 []int) int {
	l := len(nums1)
	nums1Sort := make([]int, l)
	copy(nums1Sort, nums1)
	sort.Ints(nums1Sort)
	absNum, maxDiff := 0, 0
	// 二分查找的方法
	binary := func(nums []int, t int) int {
		lo, hi := 0, len(nums)-1
		l, r := nums[lo], nums[hi]
		for lo <= hi {
			mid := (lo + hi) >> 1
			curNum := nums[mid]
			if curNum == t {
				return 0
			} else if curNum < t {
				lo, l = mid+1, nums[mid]
			} else {
				hi, r = mid-1, nums[mid]
			}
		}
		return utils.Min(utils.Abs(t-l), utils.Abs(t-r))
	}
	for i := 0; i < l; i++ {
		curAbs := utils.Abs(nums1[i] - nums2[i])
		absNum += curAbs
		maxDiff = utils.Max(maxDiff, utils.Abs(curAbs-binary(nums1Sort, nums2[i])))
	}
	return (absNum - maxDiff) % 1000000007
}

// 贪心解法，虽然能够通过用例，但是结果是错的，不可使用。
func minAbsoluteSumDiff2(nums1 []int, nums2 []int) int {
	l := len(nums1)
	absNum, maxAbs, maxAbsIndex := 0, 0, 0
	for i := 0; i < l; i++ {
		curAbs := utils.Abs(nums1[i] - nums2[i])
		absNum += curAbs
		if curAbs > maxAbs {
			maxAbs, maxAbsIndex = curAbs, i
		}
	}
	compareNum, afterChangeAbs := nums2[maxAbsIndex], maxAbs
	for _, num := range nums1 {
		curAbs := utils.Abs(num - compareNum)
		if curAbs < afterChangeAbs {
			afterChangeAbs = curAbs
		}
	}
	absNum = absNum - maxAbs + afterChangeAbs
	return absNum % 1000000007
}

func testProblem1818() {
	fmt.Println(minAbsoluteSumDiff([]int{1, 7, 5}, []int{2, 3, 5}))
	fmt.Println(minAbsoluteSumDiff([]int{2, 4, 6, 8, 10}, []int{2, 4, 6, 8, 10}))
	fmt.Println(minAbsoluteSumDiff([]int{1, 10, 4, 4, 2, 7}, []int{9, 3, 5, 1, 7, 4}))
}
