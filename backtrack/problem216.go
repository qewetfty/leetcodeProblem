package backtrack

import "fmt"

// Find all possible combinations of k numbers that add up to a number n, given that only numbers from 1 to 9 can be
// used and each combination should be a unique set of numbers.
//	Note:
//		All numbers will be positive integers.
//		The solution set must not contain duplicate combinations.
//	Example 1:
//		Input: k = 3, n = 7
//		Output: [[1,2,4]]
//	Example 2:
//		Input: k = 3, n = 9
//		Output: [[1,2,6], [1,3,5], [2,3,4]]

func combinationSum3(k int, n int) [][]int {
	res, curList := make([][]int, 0), make([]int, 0)
	backtrack216(1, k, n, 0, &curList, &res)
	return res
}

func backtrack216(start, k, n, curSum int, curList *[]int, res *[][]int) {
	if len(*curList) == k {
		if curSum != n {
			return
		}
		newList := make([]int, 0)
		for i := 0; i < len(*curList); i++ {
			newList = append(newList, (*curList)[i])
		}
		*res = append(*res, newList)
		return
	}
	for i := start; i <= 9; i++ {
		*curList, curSum = append(*curList, i), curSum+i
		backtrack216(i+1, k, n, curSum, curList, res)
		*curList, curSum = (*curList)[:len(*curList)-1], curSum-i
	}
}

func testProblem216() {
	fmt.Println(combinationSum3(3, 7))
	fmt.Println(combinationSum3(3, 9))
	fmt.Println(combinationSum3(0, 0))
	fmt.Println(combinationSum3(1, 1))
}
