package array

//Given an integer array nums, return an array answer such that answer[i] is
//equal to the product of all the elements of nums except nums[i].
//
// The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit
//integer.
//
// You must write an algorithm that runs in O(n) time and without using the
//division operation.
//
//
// Example 1:
// Input: nums = [1,2,3,4]
//Output: [24,12,8,6]
// Example 2:
// Input: nums = [-1,1,0,-3,3]
//Output: [0,0,9,0,0]
//
//
// Constraints:
//
//
// 2 <= nums.length <= 10⁵
// -30 <= nums[i] <= 30
// The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit
//integer.
//
//
//
// Follow up: Can you solve the problem in O(1) extra space complexity? (The
//output array does not count as extra space for space complexity analysis.)
// Related Topics Array Prefix Sum 👍 9666 👎 634

func productExceptSelf(nums []int) []int {
	n := len(nums)
	l, r, result := make([]int, n), make([]int, n), make([]int, n)

	l[0] = 1
	for i := 1; i < n; i++ {
		l[i] = l[i-1] * nums[i-1]
	}

	r[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		r[i] = r[i+1] * nums[i+1]
	}

	for i := 0; i < n; i++ {
		result[i] = l[i] * r[i]
	}
	return result
}
