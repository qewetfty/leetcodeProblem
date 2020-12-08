package array

import "fmt"

// Given a positive integer n, generate an n x n matrix filled with elements from 1 to n2 in spiral order.
//	Example 1:
//		Input: n = 3
//		Output: [[1,2,3],[8,9,4],[7,6,5]]
//	Example 2:
//		Input: n = 1
//		Output: [[1]]
//	Constraints:
//		1 <= n <= 20

func generateMatrix(n int) [][]int {
	l, r, t, b := 0, n-1, 0, n-1
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}
	num, target := 1, n*n
	for num <= target {
		for i := l; i <= r; i++ {
			res[t][i] = num
			num++
		}
		t++
		for i := t; i <= b; i++ {
			res[i][r] = num
			num++
		}
		r--
		for i := r; i >= l; i-- {
			res[b][i] = num
			num++
		}
		b--
		for i := b; i >= t; i-- {
			res[i][l] = num
			num++
		}
		l++
	}
	return res
}

func testProblem59() {
	fmt.Println(generateMatrix(1))
	fmt.Println(generateMatrix(2))
	fmt.Println(generateMatrix(3))
	fmt.Println(generateMatrix(4))
}
