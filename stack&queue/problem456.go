package stack_queue

import (
	"fmt"
	"github.com/leetcodeProblem/utils"
)

// Given an array of n integers nums, a 132 pattern is a subsequence of three
// integers nums[i], nums[j] and nums[k] such that i < j < k and nums[i] <
// nums[k] < nums[j].
// Return true if there is a 132 pattern in nums, otherwise, return false.
// Follow up: The O(n^2) is trivial, could you come up with the O(n logn) or the O(n) solution?
//	Example 1:
//		Input: nums = [1,2,3,4]
//		Output: false
//		Explanation: There is no 132 pattern in the sequence.
//	Example 2:
//		Input: nums = [3,1,4,2]
//		Output: true
//		Explanation: There is a 132 pattern in the sequence: [1, 4, 2].
//	Example 3:
//		Input: nums = [-1,3,2,0]
//		Output: true
//		Explanation: There are three 132 patterns in the sequence: [-1, 3, 2], [-1, 3, 0] and [-1, 2, 0].
//	Constraints:
//		n == nums.length
//		1 <= n <= 104
//		-109 <= nums[i] <= 109

func find132pattern(nums []int) bool {
	l := len(nums)
	if l < 3 {
		return false
	}
	min, stack := make([]int, l), make([]int, 0)
	min[0] = nums[0]
	for i := 1; i < l; i++ {
		min[i] = utils.Min(min[i-1], nums[i])
	}
	for j := l - 1; j >= 0; j-- {
		if nums[j] > min[j] {
			for len(stack) > 0 && stack[len(stack)-1] <= min[j] {
				stack = stack[:len(stack)-1]
			}
			if len(stack) > 0 && stack[len(stack)-1] < nums[j] {
				return true
			}
			stack = append(stack, nums[j])
		}
	}
	return false
}

func testProblem456() {
	fmt.Println(find132pattern([]int{1, 2, 3, 4}))
	fmt.Println(find132pattern([]int{3, 1, 4, 2}))
}
