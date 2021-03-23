package math

import (
	"fmt"
	"strconv"
)

// Starting with a positive integer N, we reorder the digits in any order
// (including the original order) such that the leading digit is not zero.
// Return true if and only if we can do this in a way such that the resulting number is a power of 2.
//	Example 1:
//		Input: 1
//		Output: true
//	Example 2:
//		Input: 10
//		Output: false
//	Example 3:
//		Input: 16
//		Output: true
//	Example 4:
//		Input: 24
//		Output: false
//	Example 5:
//		Input: 46
//		Output: true
//	Note:
//		1 <= N <= 10^9

// 暴力法，直接把范围内的2的幂次全存起来，判断长度内的是否符合最终结果
func reorderedPowerOf2(N int) bool {
	numberMap := getAll2Numbers()
	curStr := strconv.Itoa(N)
	length := len(curStr)
	numbers := numberMap[length]
	for _, byteMap := range numbers {
		for i := 0; i < length; i++ {
			byteMap[curStr[i]]--
		}
		find := true
		for _, count := range byteMap {
			if count != 0 {
				find = false
				break
			}
		}
		if find {
			return true
		}
	}
	return false
}

func getAll2Numbers() map[int][]map[byte]int {
	resultMap := make(map[int][]map[byte]int)
	start := 1
	for start <= 1000000000 {
		curStr := strconv.Itoa(start)
		length := len(curStr)
		charMap := make(map[byte]int)
		for i := 0; i < length; i++ {
			charMap[curStr[i]]++
		}
		resultMap[length] = append(resultMap[length], charMap)
		start = start << 1
	}
	return resultMap
}

func testProblem869() {
	fmt.Println(reorderedPowerOf2(1))
	fmt.Println(reorderedPowerOf2(10))
	fmt.Println(reorderedPowerOf2(16))
	fmt.Println(reorderedPowerOf2(24))
	fmt.Println(reorderedPowerOf2(46))
}
