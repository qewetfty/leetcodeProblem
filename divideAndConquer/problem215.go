package divideAndConquer

import (
	"fmt"
	"sort"
)

// Find the kth largest element in an unsorted array. Note that it is the kth
// largest element in the sorted order, not the kth distinct element.
//	Example 1:
//		Input: [3,2,1,5,6,4] and k = 2
//		Output: 5
//	Example 2:
//		Input: [3,2,3,1,2,4,5,5,6] and k = 4
//		Output: 4
//	Note:
//		You may assume k is always valid, 1 ≤ k ≤ array's length.

// 基于快排的思路
func findKthLargest(nums []int, k int) int {
	tar := len(nums) - k
	return partition(nums, 0, len(nums)-1, tar)
}

func partition(nums []int, start, end, target int) int {
	pivot := nums[start]
	i, j := start, end
	for i < j {
		for i < j && nums[j] > pivot {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
		for i < j && nums[i] <= pivot {
			i++
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	if i == target {
		return nums[i]
	} else if i < target {
		return partition(nums, i+1, end, target)
	} else {
		return partition(nums, start, i-1, target)
	}
}

// 直接排序，时间复杂度O(n*log(n))
func findKthLargest2(nums []int, k int) int {
	sort.Ints(nums)
	return nums[len(nums)-k]
}

func testProblem215() {
	fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
	fmt.Println(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4))
}
