package backtrack

import "fmt"

// Return all non-negative integers of length N such that the absolute difference between every two consecutive digits is K.
// Note that every number in the answer must not have leading zeros except for the number 0 itself. For example,
// 01 has one leading zero and is invalid, but 0 is valid.
// You may return the answer in any order.
//	Example 1:
//		Input: N = 3, K = 7
//		Output: [181,292,707,818,929]
//		Explanation: Note that 070 is not a valid number, because it has leading zeroes.
//	Example 2:
//		Input: N = 2, K = 1
//		Output: [10,12,21,23,32,34,43,45,54,56,65,67,76,78,87,89,98]
//	Note:
//		1 <= N <= 9
//		0 <= K <= 9

var visited map[int]bool

func numsSameConsecDiff(N int, K int) []int {
	visited = make(map[int]bool)
	res := make([]int, 0)
	if N == 1 {
		res = append(res, 0)
	}
	for i := 1; i <= 9; i++ {
		backtrack967(N, K, i, &res)
	}
	return res
}

func backtrack967(N, K, curNum int, res *[]int) {
	if N <= 1 {
		if visited[curNum] {
			return
		}
		visited[curNum] = true
		*res = append(*res, curNum)
		return
	}
	curInt := curNum % 10
	if curInt+K < 10 {
		backtrack967(N-1, K, curNum*10+curInt+K, res)
	}
	if curInt-K >= 0 {
		backtrack967(N-1, K, curNum*10+curInt-K, res)
	}
}

func pow(x, n int) int {
	if n == 0 {
		return 1
	}
	subProblem := pow(x, n>>1)
	res := 1
	if n%2 == 1 {
		res = subProblem * subProblem * x
	} else {
		res = subProblem * subProblem
	}
	return res
}

func testProblem967() {
	fmt.Println(numsSameConsecDiff(3, 0))
	fmt.Println(numsSameConsecDiff(1, 0))
	fmt.Println(numsSameConsecDiff(3, 7))
	fmt.Println(numsSameConsecDiff(2, 1))
	fmt.Println(numsSameConsecDiff(1, 1))
	fmt.Println(numsSameConsecDiff(5, 1))
	fmt.Println(numsSameConsecDiff(3, 9))
	fmt.Println(numsSameConsecDiff(3, 8))
}
