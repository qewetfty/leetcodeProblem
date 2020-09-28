package twoPointer

import "fmt"

// Your are given an array of positive integers nums.
// Count and print the number of (contiguous) subarrays where the product of all the elements in the subarray is less than k.
//	Example 1:
//		Input: nums = [10, 5, 2, 6], k = 100
//		Output: 8
//		Explanation: The 8 subarrays that have product less than 100 are: [10], [5], [2], [6], [10, 5], [5, 2], [2, 6], [5, 2, 6].
//		Note that [10, 5, 2] is not included as the product of 100 is not strictly less than k.
//	Note:
//		0 < nums.length <= 50000.
//		0 < nums[i] < 1000.
//		0 <= k < 10^6.

func numSubarrayProductLessThanK(nums []int, k int) int {
	if k <= 1 {
		return 0
	}
	l, res, left, product := len(nums), 0, 0, 1
	for right := 0; right < l; right++ {
		product = product * nums[right]
		for product >= k {
			product = product / nums[left]
			left++
		}
		res = res + right - left + 1
	}
	return res
}

// time exceed method
func numSubarrayProductLessThanK2(nums []int, k int) int {
	l, res := len(nums), 0
	for i := 0; i < l; i++ {
		product := nums[i]
		if product >= k {
			continue
		} else {
			res++
		}
		for j := i + 1; j < l; j++ {
			product = product * nums[j]
			if product < k {
				res++
			} else {
				break
			}
		}
	}
	return res
}

func testProblem713() {
	fmt.Println(numSubarrayProductLessThanK([]int{10, 5, 2, 6}, 100))
}
