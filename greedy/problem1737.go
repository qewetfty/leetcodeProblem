package greedy

import (
	"fmt"
	"github.com/leetcodeProblem/utils"
	"math"
)

// You are given two strings a and b that consist of lowercase letters. In one
// operation, you can change any character in a or b to any lowercase letter.
// Your goal is to satisfy one of the following three conditions:
//	Every letter in a is strictly less than every letter in b in the alphabet.
//	Every letter in b is strictly less than every letter in a in the alphabet.
//	Both a and b consist of only one distinct letter.
// Return the minimum number of operations needed to achieve your goal.
//	Example 1:
//		Input: a = "aba", b = "caa"
//		Output: 2
//		Explanation: Consider the best way to make each condition true:
//			1) Change b to "ccc" in 2 operations, then every letter in a is less than every letter in b.
//			2) Change a to "bbb" and b to "aaa" in 3 operations, then every letter in b is less than every letter in a.
//			3) Change a to "aaa" and b to "aaa" in 2 operations, then a and b consist of one distinct letter.
//			The best way was done in 2 operations (either condition 1 or condition 3).
//	Example 2:
//		Input: a = "dabadd", b = "cda"
//		Output: 3
//		Explanation: The best way is to make condition 1 true by changing b to "eee".
//	Constraints:
//		1 <= a.length, b.length <= 105
//		a and b consist only of lowercase letters.

// 遍历每种情况，
func minCharacters(a string, b string) int {
	allMap, aMap, bMap := make([]int, 26), make([]int, 26), make([]int, 26)
	la, lb := len(a), len(b)
	for i := 0; i < la; i++ {
		cur := int(a[i] - 'a')
		aMap[cur]++
		allMap[cur]++
	}
	for i := 0; i < lb; i++ {
		cur := int(b[i] - 'a')
		bMap[cur]++
		allMap[cur]++
	}
	// 前缀和处理
	for i := 1; i < 26; i++ {
		aMap[i] = aMap[i-1] + aMap[i]
	}
	for i := 1; i < 26; i++ {
		bMap[i] = bMap[i-1] + bMap[i]
	}
	// 计算全部变化的
	maxFreq := 0
	for _, num := range allMap {
		maxFreq = utils.Max(num, maxFreq)
	}
	result := la + lb - maxFreq
	result = utils.Min(result, utils.Min(findChangeStep(aMap, bMap), findChangeStep(bMap, aMap)))
	return result
}

func findChangeStep(aMap, bMap []int) int {
	result := math.MaxInt32
	// 以某一个字母为基准，让所有a的字符都小于等于这个字母，b的字符都大于这个字母
	for i := 0; i < 25; i++ {
		cur := aMap[25] - aMap[i] + bMap[i]
		result = utils.Min(cur, result)
	}
	return result
}

func testProblem1737() {
	fmt.Println(minCharacters("a", "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz"))
	fmt.Println(minCharacters("aba", "caa"))
	fmt.Println(minCharacters("dabadd", "cda"))
}
