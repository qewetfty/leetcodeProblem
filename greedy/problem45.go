package greedy

import "fmt"

// Given an array of non-negative integers, you are initially positioned at the first index of the array.
// Each element in the array represents your maximum jump length at that position.
// Your goal is to reach the last index in the minimum number of jumps.
//	Example:
//		Input: [2,3,1,1,4]
//		Output: 2
//	Explanation: The minimum number of jumps to reach the last index is 2.
//    	Jump 1 step from index 0 to 1, then 3 steps to the last index.
//	Note:
//		You can assume that you can always reach the last index.

// dp method
func jump(nums []int) int {
	length := len(nums)
	if length <= 1 {
		return 0
	}
	dp := make([]int, length)
	dp[0] = 0
	for i, num := range nums {
		for j := i + 1; j < i+1+num; j++ {
			if dp[j] == 0 {
				dp[j] = dp[i] + 1
			}
			if j == length-1 {
				return dp[length-1]
			}
		}
	}
	return -1
}

// greedy method
func jump1(nums []int) int {
	var (
		length = len(nums)
		end    = 0
		maxPos = 0
		res    = 0
	)
	for i := 0; i < length-1; i++ {
		maxPos = max(maxPos, i+nums[i])
		if i == end {
			end = maxPos
			res++
		}
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func testProblem45() {
	fmt.Println(jump1([]int{2, 3, 1, 1, 4}))
}
