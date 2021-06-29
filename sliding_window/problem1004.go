package sliding_window

import (
	"fmt"
	"github.com/leetcodeProblem/utils"
)

// Given a binary array nums and an integer k, return the maximum number of
// consecutive 1's in the array if you can flip at most k 0's.
//	Example 1:
//		Input: nums = [1,1,1,0,0,0,1,1,1,1,0], k = 2
//		Output: 6
//		Explanation: [1,1,1,0,0,1,1,1,1,1,1]
//			Bolded numbers were flipped from 0 to 1. The longest subarray is underlined.
//	Example 2:
//		Input: nums = [0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1], k = 3
//		Output: 10
//		Explanation: [0,0,1,1,1,1,1,1,1,1,1,1,0,0,0,1,1,1,1]
//			Bolded numbers were flipped from 0 to 1. The longest subarray is underlined.
//	Constraints:
//		1 <= nums.length <= 105
//		nums[i] is either 0 or 1.
//		0 <= k <= nums.length

func longestOnes(nums []int, k int) int {
	left, right := 0, 0
	zeroCount := 0
	l := len(nums)
	result := 0
	for right < l {
		curNum := nums[right]
		if curNum == 0 {
			zeroCount++
		}
		for zeroCount > k && left <= right {
			num := nums[left]
			if num == 0 {
				zeroCount--
			}
			left++
		}
		result = utils.Max(result, right-left+1)
		right++
	}
	return result
}

func testProblem1004() {
	fmt.Println(longestOnes([]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2))
	fmt.Println(longestOnes([]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3))
}
