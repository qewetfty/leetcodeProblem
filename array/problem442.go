package array

import (
	"fmt"
	"math"
)

// Given an array of integers, 1 ≤ a[i] ≤ n (n = size of array), some elements appear twice and others appear once.
// Find all the elements that appear twice in this array.
// Could you do it without extra space and in O(n) runtime?
//	Example:
//		Input:
//			[4,3,2,7,8,2,3,1]
//		Output:
//			[2,3]

func findDuplicates(nums []int) []int {
	res := make([]int, 0)
	for _, num := range nums {
		v := int(math.Abs(float64(num)))
		if nums[v-1] < 0 {
			res = append(res, v)
		} else {
			nums[v-1] = -nums[v-1]
		}
	}
	return res
}

func testProblem442() {
	fmt.Println(findDuplicates([]int{4, 3, 2, 7, 8, 2, 3, 1}))
}
