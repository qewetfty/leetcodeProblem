package divideAndConquer

import "fmt"

// Given a triangle, find the minimum path sum from top to bottom. Each step you may move to adjacent numbers on the row below.
// For example, given the following triangle
//		[
//		     [2],
//		    [3,4],
//		   [6,5,7],
//		  [4,1,8,3]
//		]
//	The minimum path sum from top to bottom is 11 (i.e., 2 + 3 + 5 + 1 = 11).
//	Note:
//		Bonus point if you are able to do this using only O(n) extra space,
//		where n is the total number of rows in the triangle.

func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}
	triLen := len(triangle)
	res := triangle[triLen-1]
	for i := triLen - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			res[j] = min(res[j], res[j+1]) + triangle[i][j]
		}
	}
	return res[0]
}

func min(x, y int) int {
	if x <= y {
		return x
	}
	return y
}

func testProblem120() {
	fmt.Println(minimumTotal([][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}))
}
