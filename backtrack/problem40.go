package backtrack

import (
	"fmt"
	"sort"
)

// Given a collection of candidate numbers (candidates) and a target number (target), find all unique combinations in
// candidates where the candidate numbers sums to target.
// Each number in candidates may only be used once in the combination.
//	Note:
//		All numbers (including target) will be positive integers.
//		The solution set must not contain duplicate combinations.
//	Example 1:
//		Input: candidates = [10,1,2,7,6,1,5], target = 8,
//		A solution set is:
//		[
//		  [1, 7],
//		  [1, 2, 5],
//		  [2, 6],
//		  [1, 1, 6]
//		]
//	Example 2:
//		Input: candidates = [2,5,2,1,2], target = 5,
//		A solution set is:
//		[
//		  [1,2,2],
//		  [5]
//		]

var curList []int
var result [][]int
var l int

func combinationSum2(candidates []int, target int) [][]int {
	curList, result, l = make([]int, 0), make([][]int, 0), len(candidates)
	if l == 0 {
		return result
	}
	// 排序并去除比target大的数
	sort.Ints(candidates)
	i := l - 1
	for i >= 0 {
		if candidates[i] <= target {
			break
		}
		i--
	}
	candidates = candidates[:i+1]
	l = len(candidates)
	backtrack40(candidates, 0, 0, target)
	return result
}

func backtrack40(candidates []int, curIndex, curSum, target int) {
	if curSum > target {
		return
	}
	if curSum == target {
		result = append(result, append([]int{}, curList...))
		return
	}
	for i := curIndex; i < l; i++ {
		if i > curIndex && candidates[i] == candidates[i-1] {
			continue
		}
		curList, curSum = append(curList, candidates[i]), curSum+candidates[i]
		backtrack40(candidates, i+1, curSum, target)
		curList, curSum = curList[:len(curList)-1], curSum-candidates[i]
	}
}

func testProblem40() {
	fmt.Println(combinationSum2([]int{3, 1, 3, 5, 1, 1}, 8))
	fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
	fmt.Println(combinationSum2([]int{2, 5, 2, 1, 2}, 5))
}
