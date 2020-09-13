package backtrack

import "fmt"

// Given a set of candidate numbers (candidates) (without duplicates) and a target number (target), find all unique
// combinations in candidates where the candidate numbers sums to target.
// The same repeated number may be chosen from candidates unlimited number of times.
//	Note:
//		All numbers (including target) will be positive integers.
//		The solution set must not contain duplicate combinations.
//	Example 1:
//		Input: candidates = [2,3,6,7], target = 7,
//		A solution set is:
//		[
//		  [7],
//		  [2,2,3]
//		]
//	Example 2:
//		Input: candidates = [2,3,5], target = 8,
//		A solution set is:
//		[
//		  [2,2,2,2],
//		  [2,3,3],
//		  [3,5]
//		]
//	Constraints:
//		1 <= candidates.length <= 30
//		1 <= candidates[i] <= 200
//		Each element of candidate is unique.
//		1 <= target <= 500

func combinationSum(candidates []int, target int) [][]int {
	res, curList := make([][]int, 0), make([]int, 0)
	backtrack39(candidates, target, 0, &curList, &res)
	return res
}

func backtrack39(candidates []int, target, curSum int, curList *[]int, res *[][]int) {
	if curSum > target {
		return
	}
	if curSum == target {
		newList := make([]int, 0)
		for i := 0; i < len(*curList); i++ {
			newList = append(newList, (*curList)[i])
		}
		*res = append(*res, newList)
		return
	}
	for i := 0; i < len(candidates); i++ {
		if len(*curList) != 0 && candidates[i] < (*curList)[len(*curList)-1] {
			continue
		}
		*curList, curSum = append(*curList, candidates[i]), curSum+candidates[i]
		backtrack39(candidates, target, curSum, curList, res)
		*curList, curSum = (*curList)[:(len(*curList))-1], curSum-candidates[i]
	}
}

func testProblem39() {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
	fmt.Println(combinationSum([]int{2, 3, 5}, 8))
}
