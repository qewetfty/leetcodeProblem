package stack_queue

import (
	"container/list"
	"fmt"
)

// Given an array nums, there is a sliding window of size k which is moving from the very left of the
// array to the very right. You can only see the k numbers in the window.
// Each time the sliding window moves right by one position. Return the max sliding window.

//	Note:
//		You may assume k is always valid, 1 ≤ k ≤ input array's size for non-empty array.
//	Follow up:
//		Could you solve it in linear time?

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums)*k == 0 {
		return []int{}
	}
	if k == 1 {
		return nums
	}
	deque := new(list.List)
	output := make([]int, 0)
	maxIndex := 0
	for i := 0; i < k; i++ {
		cleanDeque(deque, nums, i, k)
		deque.PushBack(i)
		if nums[i] > nums[maxIndex] {
			maxIndex = i
		}
	}
	output = append(output, nums[maxIndex])
	for i := k; i < len(nums); i++ {
		cleanDeque(deque, nums, i, k)
		deque.PushBack(i)
		output = append(output, nums[deque.Front().Value.(int)])
	}
	return output
}

func cleanDeque(list *list.List, nums []int, i, k int) {
	if list.Len() > 0 && list.Front().Value.(int) == i-k {
		list.Remove(list.Front())
	}
	for list.Len() > 0 && nums[i] > nums[list.Back().Value.(int)] {
		list.Remove(list.Back())
	}
}

func testProblem239() {
	a := []int{1, 3, -1, -3, 5, 3, 6, 7}
	fmt.Println(maxSlidingWindow(a, 3))
}
