package array

// Given an array of integers,
// return indices of the two numbers such that they add up to a specific target.

// You may assume that each input would have exactly one solution,
// and you may not use the same element twice.

func twoSum(nums []int, target int) []int {
	set := make(map[int]int)
	for index, n := range nums {
		set[n] = index
	}
	for index, n := range nums {
		other := target - n
		if value, exist := set[other]; exist && value != index {
			return []int{index, value}
		}
	}
	return []int{-1, -1}
}
