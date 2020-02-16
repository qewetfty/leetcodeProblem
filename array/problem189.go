package array

import "fmt"

// Given an array, rotate the array to the right by k steps, where k is non-negative.
//	Example 1:
//		Input: [1,2,3,4,5,6,7] and k = 3
//		Output: [5,6,7,1,2,3,4]
//	Explanation:
//		rotate 1 steps to the right: [7,1,2,3,4,5,6]
//		rotate 2 steps to the right: [6,7,1,2,3,4,5]
//		rotate 3 steps to the right: [5,6,7,1,2,3,4]
//
//	Example 2:
//		Input: [-1,-100,3,99] and k = 2
//		Output: [3,99,-1,-100]
//	Explanation:
//		rotate 1 steps to the right: [99,-1,-100,3]
//		rotate 2 steps to the right: [3,99,-1,-100]
// Note:
//	Try to come up as many solutions as you can, there are at least 3 different ways to solve this problem.
//	Could you do it in-place with O(1) extra space?

// force method
func rotate(nums []int, k int) {
	for i := 0; i < k; i++ {
		prev := nums[len(nums)-1]
		for j := 0; j < len(nums); j++ {
			temp := nums[j]
			nums[j] = prev
			prev = temp
		}
	}
}

// extra space method
func rotate2(nums []int, k int) {
	result := make([]int, len(nums))
	for i := range nums {
		result[(i+k)%len(nums)] = nums[i]
	}
	for i := range nums {
		nums[i] = result[i]
	}
}

// reverse method
func rotate3(nums []int, k int) {
	k = k % len(nums)
	reverse(nums, 0, len(nums)-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, len(nums)-1)
}

func reverse(nums []int, start, end int) {
	for start < end {
		temp := nums[start]
		nums[start] = nums[end]
		nums[end] = temp
		start++
		end--
	}
}

func testProblem189() {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	rotate3(a, 3)
	fmt.Println(a)
	b := []int{-1, -100, 3, 99}
	rotate3(b, 2)
	fmt.Println(b)
	c := []int{-1}
	rotate3(c, 2)
	fmt.Println(c)
}
