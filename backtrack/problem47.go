package backtrack

import (
	"fmt"
	"sort"
)

// Given a collection of numbers that might contain duplicates, return all possible unique permutations.
//	Example:
//		Input: [1,1,2]
//		Output:
//		[
//		  [1,1,2],
//		  [1,2,1],
//		  [2,1,1]
//		]

func permuteUnique(nums []int) [][]int {
	res := make([][]int, 0)
	if len(nums) == 0 {
		return res
	}
	sort.Ints(nums)
	data := make([]int, 0)
	useList := make([]bool, len(nums))
	backtrack1(nums, useList, data, &res)
	return res
}

func backtrack1(nums []int, useList []bool, data []int, res *[][]int) {
	if len(nums) == len(data) {
		newData := make([]int, 0)
		newData = append(newData, data[0:]...)
		*res = append(*res, newData)
		return
	}
	for i := 0; i < len(nums); i++ {
		if useList[i] {
			continue
		}
		if i > 0 && nums[i] == nums[i-1] && !useList[i-1] {
			continue
		}
		data = append(data, nums[i])
		useList[i] = true
		backtrack1(nums, useList, data, res)
		data = data[:len(data)-1]
		useList[i] = false
	}
}

func testProblem47() {
	fmt.Println(permuteUnique([]int{1, 1, 2}))
}
