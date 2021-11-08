package backtrack

import "fmt"

// Given an integer n, return the number of structurally unique BST's (binary
// search trees) which has exactly n nodes of unique values from 1 to n.
//	Example 1:
//		Input: n = 3
//		Output: 5
//	Example 2:
//		Input: n = 1
//		Output: 1
//	Constraints:
//		1 <= n <= 19

var resList96 []int

func numTrees(n int) int {
	resList96 = make([]int, 20)
	resList96[0], resList96[1], resList96[2] = 1, 1, 2
	return backtrack96(n)
}

func backtrack96(i int) int {
	if resList96[i] > 0 {
		return resList96[i]
	}
	num := 0
	cal := i - 1
	for j := 0; j <= cal; j++ {
		num = num + backtrack96(j)*backtrack96(cal-j)
	}
	resList96[i] = num
	return num
}

func testProblem96() {
	fmt.Println(numTrees(3))
	fmt.Println(numTrees(4))
	fmt.Println(numTrees(5))
	fmt.Println(numTrees(1))
	fmt.Println(numTrees(2))
}
