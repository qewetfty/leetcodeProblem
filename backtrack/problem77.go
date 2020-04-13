package backtrack

import "fmt"

// Given two integers n and k, return all possible combinations of k numbers out of 1 ... n.
//	Example:
//		Input: n = 4, k = 2
//		Output:
//		[
//		  [2,4],
//		  [3,4],
//		  [2,3],
//		  [1,2],
//		  [1,3],
//		  [1,4],
//		]

func combine(n int, k int) [][]int {
	res := make([][]int, 0)
	if n == 0 || k == 0 {
		return res
	}
	data := make([]int, 0)
	backtrack(n, k, 1, data, &res)
	return res
}

func backtrack(n int, k int, start int, data []int, res *[][]int) {
	if len(data) == k {
		d := make([]int, 0)
		d = append(d, data[0:]...)
		*res = append(*res, d)
		return
	}
	for i := start; i < n+1; i++ {
		data = append(data, i)
		backtrack(n, k, i+1, data, res)
		data = data[0 : len(data)-1]
	}
}

func testProblem77() {
	fmt.Println(combine(4, 2))
}
