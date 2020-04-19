package backtrack

import "fmt"

// Given a set of distinct integers, nums, return all possible subsets (the power set).
//	Note: The solution set must not contain duplicate subsets.
//	Example:
//		Input: nums = [1,2,3]
//		Output:
//		[
//		  [3],
//		  [1],
//		  [2],
//		  [1,2,3],
//		  [1,3],
//		  [2,3],
//		  [1,2],
//		  []
//		]

func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	data := make([]int, 0)
	backtrack(nums, 0, data, &res)
	return res
}

func backtrack(nums []int, lastIndex int, data []int, res *[][]int) {
	newData := make([]int, 0)
	newData = append(newData, data...)
	*res = append(*res, newData)
	for i := lastIndex; i < len(nums); i++ {
		data = append(data, nums[i])
		backtrack(nums, i+1, data, res)
		data = data[:len(data)-1]
	}
}

func testProblem78() {
	fmt.Println(subsets([]int{1, 2, 3}))
}
