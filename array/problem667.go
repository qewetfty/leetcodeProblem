package array

import "fmt"

// Given two integers n and k, you need to construct a list which contains n
// different positive integers ranging from 1 to n and obeys the following
// requirement:
// Suppose this list is [a1, a2, a3, ... , an], then the list [|a1 -
// a2|, |a2 - a3|, |a3 - a4|, ... , |an-1 - an|] has exactly k distinct integers.
//	If there are multiple answers, print any of them.
//	Example 1:
//		Input: n = 3, k = 1
//		Output: [1, 2, 3]
//		Explanation: The [1, 2, 3] has three different positive integers ranging from 1 to 3, and the [1, 1] has exactly 1 distinct integer: 1.
//	Example 2:
//		Input: n = 3, k = 2
//		Output: [1, 3, 2]
//		Explanation: The [1, 3, 2] has three different positive integers ranging from 1 to 3, and the [2, 1] has exactly 2 distinct integers: 1 and 2.
//	Note:
//		The n and k are in the range 1 <= k < n <= 10^4.

func constructArray(n int, k int) []int {
	usedList := make([]bool, n+1)
	result := make([]int, 0)
	usedList[1], result = true, append(result, 1)
	positive := true
	for k > 0 && len(result) < n {
		lastNum := result[len(result)-1]
		var nextNum int
		if positive {
			nextNum = lastNum + k
		} else {
			nextNum = lastNum - k
		}
		result, usedList[nextNum] = append(result, nextNum), true
		k--
		positive = !positive
	}
	for i := 1; i <= n; i++ {
		if !usedList[i] {
			result = append(result, i)
		}
	}
	return result
}

func testProblem667() {
	fmt.Println(constructArray(3, 1))
	fmt.Println(constructArray(3, 2))
}
