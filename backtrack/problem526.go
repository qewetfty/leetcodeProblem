package backtrack

import "fmt"

// Suppose you have n integers from 1 to n. We define a beautiful arrangement as
// an array that is constructed by these n numbers successfully if one of the
// following is true for the ith position (1 <= i <= n) in this array:
//	The number at the ith position is divisible by i.
//	i is divisible by the number at the ith position.
// Given an integer n, return the number of the beautiful arrangements that you can construct.
//	Example 1:
//		Input: n = 2
//		Output: 2
//		Explanation:
//			The first beautiful arrangement is [1, 2]:
//				Number at the 1st position (i=1) is 1, and 1 is divisible by i (i=1).
//				Number at the 2nd position (i=2) is 2, and 2 is divisible by i (i=2).
//			The second beautiful arrangement is [2, 1]:
//				Number at the 1st position (i=1) is 2, and 2 is divisible by i (i=1).
//				Number at the 2nd position (i=2) is 1, and i (i=2) is divisible by 1.
//	Example 2:
//		Input: n = 1
//		Output: 1
//	Constraints:
//		1 <= n <= 15

func countArrangement(n int) int {
	numberList, curList := make([]bool, n), make([]int, 0)
	result := 0
	backtrack526(n, numberList, curList, &result)
	return result
}

func backtrack526(n int, numberList []bool, curList []int, result *int) {
	if len(curList) == n {
		*result++
		return
	}
	for i, visited := range numberList {
		if visited {
			continue
		}
		// 判断是否满足两个条件
		number, listPosition := i+1, len(curList)+1
		if number%listPosition == 0 || listPosition%number == 0 {
			numberList[i] = true
			curList = append(curList, number)
			backtrack526(n, numberList, curList, result)
			numberList[i] = false
			curList = curList[:len(curList)-1]
		}
	}
}

func testProblem526() {
	fmt.Println(countArrangement(2))
	fmt.Println(countArrangement(1))
	fmt.Println(countArrangement(15))
	fmt.Println(countArrangement(3))
	fmt.Println(countArrangement(4))
	fmt.Println(countArrangement(5))
	fmt.Println(countArrangement(6))
	fmt.Println(countArrangement(7))
}
