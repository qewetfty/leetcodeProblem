package dynamic_programming

import (
	"fmt"
	"sort"
)

// Given an array of unique integers, arr, where each integer arr[i] is strictly greater than 1.
// We make a binary tree using these integers, and each number may be used for
// any number of times. Each non-leaf node's value should be equal to the product
// of the values of its children.
// Return the number of binary trees we can make. The answer may be too large so return the answer modulo 109 + 7.
//	Example 1:
//		Input: arr = [2,4]
//		Output: 3
//		Explanation: We can make these trees: [2], [4], [4, 2, 2]
//	Example 2:
//		Input: arr = [2,4,5,10]
//		Output: 7
//		Explanation: We can make these trees: [2], [4], [5], [10], [4, 2, 2], [10, 2, 5], [10, 5, 2].
//	Constraints:
//		1 <= arr.length <= 1000
//		2 <= arr[i] <= 109

var m = 1000000007

func numFactoredBinaryTrees(arr []int) int {
	sort.Ints(arr)
	l := len(arr)
	dp := make([]int, l)
	for i := 0; i < l; i++ {
		dp[i] = 1
	}
	numberMap := make(map[int]int)
	for i, num := range arr {
		numberMap[num] = i
	}

	for i := 0; i < l; i++ {
		for j := 0; j < i; j++ {
			if arr[i]%arr[j] == 0 {
				other := arr[i] / arr[j]
				if _, ok := numberMap[other]; ok {
					dp[i] = (dp[i] + dp[j]*dp[numberMap[other]]) % m
				}
			}
		}
	}
	res := 0
	for _, num := range dp {
		res = (res + num) % m
	}
	return res
}

func testProblem823() {
	fmt.Println(numFactoredBinaryTrees([]int{2, 4}))
	fmt.Println(numFactoredBinaryTrees([]int{2, 4, 5, 10}))
}
