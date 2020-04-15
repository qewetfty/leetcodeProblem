package backtrack

import "fmt"

// Given a collection of distinct integers, return all possible permutations.
// 	Example:
//		Input: [1,2,3]
//		Output:
//		[
//		  [1,2,3],
//		  [1,3,2],
//		  [2,1,3],
//		  [2,3,1],
//		  [3,1,2],
//		  [3,2,1]
//		]

func permute(nums []int) [][]int {
	res := make([][]int, 0)
	if len(nums) == 0 {
		return res
	}
	data := make([]int, 0)
	selectList := make(map[int]int, len(nums))
	backtrack2(nums, selectList, data, &res)
	return res
}

func backtrack2(nums []int, selectList map[int]int, data []int, res *[][]int) {
	if len(data) == len(nums) {
		newData := make([]int, 0)
		newData = append(newData, data[0:]...)
		*res = append(*res, newData)
		return
	}
	for i := 0; i < len(nums); i++ {
		if _, contains := selectList[nums[i]]; contains {
			continue
		}
		data = append(data, nums[i])
		selectList[nums[i]] = i
		backtrack2(nums, selectList, data, res)
		data = data[:len(data)-1]
		delete(selectList, nums[i])
	}
}

func testProblem46() {
	fmt.Println(permute([]int{1, 2, 3}))
}
