package array

import (
	"fmt"
	"github.com/leetcodeProblem/utils"
)

// Given an array nums, partition it into two (contiguous) subarrays left and right so that:
//	Every element in left is less than or equal to every element in right.
//	left and right are non-empty.
//	left has the smallest possible size.
// Return the length of left after such a partitioning.  It is guaranteed that such a partitioning exists.
//	Example 1:
//		Input: nums = [5,0,3,8,6]
//		Output: 3
//		Explanation: left = [5,0,3], right = [8,6]
//	Example 2:
//		Input: nums = [1,1,1,0,6,12]
//		Output: 4
//		Explanation: left = [1,1,1,0], right = [6,12]
//	Note:
//		2 <= nums.length <= 30000
//		0 <= nums[i] <= 106
//		It is guaranteed there is at least one way to partition nums as described.

func partitionDisjoint(nums []int) int {
	left, leftMax, curMax := 0, nums[0], nums[0]
	l, i := len(nums), 1
	for i < l {
		for i < l && leftMax <= nums[i] {
			curMax = utils.Max(curMax, nums[i])
			i++
		}
		if i == l {
			break
		}
		left, leftMax = i, curMax
		i++
	}
	return left + 1
}

func testProblem915() {
	fmt.Println(partitionDisjoint([]int{5, 0, 3, 8, 6}))
	fmt.Println(partitionDisjoint([]int{1, 1, 1, 0, 6, 12}))
}
