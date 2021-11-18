package array

import "fmt"

// Given an array nums of n integers where nums[i] is in the range [1, n], return
// an array of all the integers in the range [1, n] that do not appear in nums.
//	Example 1:
//		Input: nums = [4,3,2,7,8,2,3,1]
//		Output: [5,6]
//	Example 2:
//		Input: nums = [1,1]
//		Output: [2]
//	Constraints:
//		n == nums.length
//		1 <= n <= 105
//		1 <= nums[i] <= n
//	Follow up: Could you do it without extra space and in O(n) runtime? You may assume the returned list does not count as extra space.

func findDisappearedNumbers(nums []int) []int {
	l := len(nums)
	for i := 1; i <= l; i++ {
		curNum := nums[i-1]
		for curNum != i && nums[curNum-1] != curNum {
			nums[i-1], nums[curNum-1] = nums[curNum-1], nums[i-1]
			curNum = nums[i-1]
		}
	}
	result := make([]int, 0)
	for i := 1; i <= l; i++ {
		if nums[i-1] != i {
			result = append(result, i)
		}
	}
	return result
}

func testProblem448() {
	fmt.Println(findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1}))
	fmt.Println(findDisappearedNumbers([]int{1, 1}))
}
