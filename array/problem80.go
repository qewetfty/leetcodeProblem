package array

import "fmt"

// Given a sorted array nums, remove the duplicates in-place such that duplicates
// appeared at most twice and return the new length.
// Do not allocate extra space for another array; you must do this by modifying
// the input array in-place with O(1) extra memory.
//	Clarification:
//		Confused why the returned value is an integer, but your answer is an array?
//		Note that the input array is passed in by reference, which means a modification to the input array will be known to the caller.
//	Internally you can think of this:
//		// nums is passed in by reference. (i.e., without making a copy)
//		int len = removeDuplicates(nums);
//		// any modification to nums in your function would be known by the caller.
//		// using the length returned by your function, it prints the first len elements.
//			for (int i = 0; i < len; i++) {
//			    print(nums[i]);
//			}
//	Example 1:
//		Input: nums = [1,1,1,2,2,3]
//		Output: 5, nums = [1,1,2,2,3]
//		Explanation: Your function should return length = 5, with the first five elements of nums being 1, 1, 2, 2 and 3 respectively. It doesn't matter what you leave beyond the returned length.
//	Example 2:
//		Input: nums = [0,0,1,1,1,1,2,3,3]
//		Output: 7, nums = [0,0,1,1,2,3,3]
//		Explanation: Your function should return length = 7, with the first seven elements of nums being modified to 0, 0, 1, 1, 2, 3 and 3 respectively. It doesn't matter what values are set beyond the returned length.
//	Constraints:
//		0 <= nums.length <= 3 * 104
//		-104 <= nums[i] <= 104
//		nums is sorted in ascending order.

func removeDuplicates80(nums []int) int {
	l := len(nums)
	if l <= 2 {
		return l
	}
	first, second, count := 0, 0, 1
	for first = first + 1; first < l; first++ {
		if nums[first] == nums[second] {
			if count < 2 {
				second++
				nums[second] = nums[first]
				count++
			}
		} else {
			second++
			nums[second] = nums[first]
			count = 1
		}
	}
	return second + 1
}

func testProblem80() {
	b := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	fmt.Println(removeDuplicates80(b))
	fmt.Println(b)
	a := []int{1, 1, 1, 2, 2, 3}
	fmt.Println(removeDuplicates80(a))
	fmt.Println(a)
}
