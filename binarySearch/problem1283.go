package binarySearch

import "fmt"

// Given an array of integers nums and an integer threshold, we will choose a
// positive integer divisor and divide all the array by it and sum the result of
// the division. Find the smallest divisor such that the result mentioned above
// is less than or equal to threshold.
//
// Each result of division is rounded to the nearest integer greater than or
// equal to that element. (For example: 7/3 = 3 and 10/2 = 5).
// It is guaranteed that there will be an answer.
//	Example 1:
//		Input: nums = [1,2,5,9], threshold = 6
//		Output: 5
//		Explanation: We can get a sum to 17 (1+2+5+9) if the divisor is 1.
//		If the divisor is 4 we can get a sum to 7 (1+1+2+3) and if the divisor is 5 the sum will be 5 (1+1+1+2).
//	Example 2:
//		Input: nums = [2,3,5,7,11], threshold = 11
//		Output: 3
//	Example 3:
//		Input: nums = [19], threshold = 5
//		Output: 4
//	Constraints:
//		1 <= nums.length <= 5 * 10^4
//		1 <= nums[i] <= 10^6
//		nums.length <= threshold <= 10^6

func smallestDivisor(nums []int, threshold int) int {
	l, h, res := 1, findMaxNumInList(nums), -1
	for l <= h {
		mid, total := (l+h)/2, 0
		for _, num := range nums {
			total = total + (num-1)/mid + 1
		}
		if total <= threshold {
			res, h = mid, mid-1
		} else {
			l = mid + 1
		}
	}
	return res
}

func findMaxNumInList(nums []int) int {
	max, l := nums[0], len(nums)
	for i := 1; i < l; i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

func testProblem1283() {
	fmt.Println(smallestDivisor([]int{1, 2, 5, 9}, 6))
	fmt.Println(smallestDivisor([]int{2, 3, 5, 7, 11}, 11))
	fmt.Println(smallestDivisor([]int{19}, 5))
}
