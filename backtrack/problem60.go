package backtrack

import "fmt"

// The set [1, 2, 3, ..., n] contains a total of n! unique permutations.
// By listing and labeling all of the permutations in order, we get the following sequence for n = 3:
//	"123"
//	"132"
//	"213"
//	"231"
//	"312"
//	"321"
// Given n and k, return the kth permutation sequence.
//	Example 1:
//		Input: n = 3, k = 3
//		Output: "213"
//	Example 2:
//		Input: n = 4, k = 9
//		Output: "2314"
//	Example 3:
//		Input: n = 3, k = 1
//		Output: "123"
//	Constraints:
//		1 <= n <= 9
//		1 <= k <= n!

var curNum int
var used []bool

func getPermutation(n int, k int) string {
	used = make([]bool, n+1)
	curNum = 1
	result := make([]byte, 0)
	res, _ := backtrack60(n, k, result)
	return res
}

func backtrack60(n, k int, result []byte) (string, bool) {
	if len(result) == n {
		if curNum == k {
			return string(result), true
		}
		curNum++
		return "", false
	}
	for i := 1; i <= n; i++ {
		if !used[i] {
			used[i] = true
			result = append(result, byte(i)+'0')
			res, find := backtrack60(n, k, result)
			if find {
				return res, true
			}
			result = result[:len(result)-1]
			used[i] = false
		}
	}
	return "", false
}

func testProblem60() {
	fmt.Println(getPermutation(3, 3))
	fmt.Println(getPermutation(4, 9))
	fmt.Println(getPermutation(3, 1))
}
