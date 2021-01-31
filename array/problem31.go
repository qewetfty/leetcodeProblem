package array

// Implement next permutation, which rearranges numbers into the lexicographically next greater permutation of numbers.
// If such an arrangement is not possible, it must rearrange it as the lowest possible order (i.e., sorted in ascending order).
// The replacement must be in place and use only constant extra memory.
//	Example 1:
//		Input: nums = [1,2,3]
//		Output: [1,3,2]
//	Example 2:
//		Input: nums = [3,2,1]
//		Output: [1,2,3]
//	Example 3:
//		Input: nums = [1,1,5]
//		Output: [1,5,1]
//	Example 4:
//		Input: nums = [1]
//		Output: [1]
//	Constraints:
//		1 <= nums.length <= 100
//		0 <= nums[i] <= 100

func nextPermutation(nums []int) {
	l := len(nums)
	i := l - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	if i >= 0 {
		j := l - 1
		for j >= 0 && nums[i] >= nums[j] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	reverse31(nums, i+1, l-1)
}

func reverse31(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}
