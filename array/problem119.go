package array

import "fmt"

// Given a non-negative index k where k ≤ 33, return the kth index row of the Pascal's triangle.
// Note that the row index starts from 0.
// In Pascal's triangle, each number is the sum of the two numbers directly above it.
//	Example:
//		Input: 3
//		Output: [1,3,3,1]
//	Follow up:
//		Could you optimize your algorithm to use only O(k) extra space?

func getRow(rowIndex int) []int {
	res := make([]int, rowIndex+1)
	for i := 0; i < rowIndex+1; i++ {
		res[i] = 1
	}
	for i := 2; i < rowIndex+1; i++ {
		for j := 1; j < i; j++ {
			index := i - j
			res[index] = res[index] + res[index-1]
		}
	}
	return res
}

func testProblem119() {
	fmt.Println(getRow(0))
	fmt.Println(getRow(1))
	fmt.Println(getRow(2))
	fmt.Println(getRow(3))
	fmt.Println(getRow(4))
	fmt.Println(getRow(5))
	fmt.Println(getRow(6))
	fmt.Println(getRow(7))
}
